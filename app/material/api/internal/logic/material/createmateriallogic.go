// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"context"
	"errors"
	"matman-backend/app/material/api/internal/logic/errcode"
	"matman-backend/app/material/domain/entity"
	"matman-backend/app/material/domain/repository"
	"matman-backend/common/utils/response"

	"matman-backend/app/material/api/internal/svc"
	"matman-backend/app/material/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建新物料
func NewCreateMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMaterialLogic {
	return &CreateMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMaterialLogic) CreateMaterial(req *types.CreateMaterialRequest) (resp *types.MaterialInfo, err error) {
	// 1. 检查物料编码是否已存在
	// (模仿 RegisterLogic)
	_, err = l.svcCtx.MaterialRepo.FindByCode(l.ctx, req.Code)
	if err == nil {
		// 找到了，说明已存在
		return nil, errcode.ErrMaterialCodeExists
	}
	// 必须确保错误是 ErrNotFound 才能继续
	if !errors.Is(err, repository.ErrNotFound) {
		// 其他数据库错误
		l.Logger.Errorf("FindByCode error: %v", err)
		return nil, errcode.ErrInternalError
	}

	// 2. 创建领域实体 (Domain Entity)
	// (我们已在 domain/entity 中移除了 DTO 依赖)
	mat, err := entity.NewMaterial(
		req.Code,
		req.Name,
		req.MaterialType,
		req.Spec,
		req.Unit,
		req.SupplierName,
		req.Price,
		req.StockQuantity,
	)
	if err != nil {
		// 实体创建失败 (例如 price < 0)，返回 400 业务错误
		l.Logger.Infof("NewMaterial validation error: %v", err)
		return nil, response.NewBizError(response.RequestError, err.Error())
	}

	// 3. 持久化到仓储
	if err := l.svcCtx.MaterialRepo.Create(l.ctx, mat); err != nil {
		l.Logger.Errorf("MaterialRepo.Create error: %v", err)
		return nil, errcode.ErrInternalError
	}

	// 4. 转换实体为 DTO (视图) 并返回
	// (mat 实体现在已经有了 ID 和 CreatedAt)
	return l.svcCtx.Converter.ToMaterialInfoResponse(mat), nil
}
