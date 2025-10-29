// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"context"
	"errors"
	"fmt"
	"matman-backend/app/material/api/internal/logic/errcode"
	"matman-backend/app/material/domain/entity"
	"matman-backend/app/material/domain/repository"
	"matman-backend/common/utils/response"

	"matman-backend/app/material/api/internal/svc"
	"matman-backend/app/material/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateBomEntryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建或更新BOM条目 (唯一键: parent/child)
func NewCreateOrUpdateBomEntryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateBomEntryLogic {
	return &CreateOrUpdateBomEntryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateBomEntryLogic) CreateOrUpdateBomEntry(req *types.CreateOrUpdateBomEntryRequest) (resp *types.BomEntryInfo, err error) {
	// 1. 创建领域实体 (包含业务规则校验，如 Parent != Child)
	status := 1 // 默认为 1
	if req.Status == 0 {
		status = 0
	}

	bomEntity, err := entity.NewBomEntry(req.ParentMaterialCode, req.ChildMaterialCode, req.Quantity, status)
	if err != nil {
		// (例如 Parent == Child, Quantity <= 0)
		l.Logger.Infof("NewBomEntry validation error: %v", err)
		return nil, response.NewBizError(response.RequestError, err.Error())
	}

	// 2. (安全检查) 确保父物料和子物料都存在于
	_, err = l.svcCtx.MaterialRepo.FindByCode(l.ctx, req.ParentMaterialCode)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			msg := fmt.Sprintf("父物料 %s 不存在", req.ParentMaterialCode)
			return nil, response.NewBizError(response.RequestError, msg)
		}
		return nil, errcode.ErrInternalError
	}

	_, err = l.svcCtx.MaterialRepo.FindByCode(l.ctx, req.ChildMaterialCode)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			msg := fmt.Sprintf("子物料 %s 不存在", req.ChildMaterialCode)
			return nil, response.NewBizError(response.RequestError, msg)
		}
		return nil, errcode.ErrInternalError
	}

	// 3. 调用仓储 (impl 中已实现 "CreateOrUpdate" 逻辑)
	savedEntity, err := l.svcCtx.BomRepo.CreateOrUpdate(l.ctx, bomEntity)
	if err != nil {
		l.Logger.Errorf("BomRepo.CreateOrUpdate error: %v", err)
		return nil, errcode.ErrInternalError
	}

	// 4. 转换并返回
	return l.svcCtx.Converter.ToBomEntryInfoResponse(savedEntity), nil
}
