// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"context"
	"errors"
	"matman-backend/app/material/api/internal/logic/errcode"
	"matman-backend/app/material/domain/repository"

	"matman-backend/app/material/api/internal/svc"
	"matman-backend/app/material/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新物料 (根据 Code)
func NewUpdateMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMaterialLogic {
	return &UpdateMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *UpdateMaterialLogic) UpdateMaterial(req *types.UpdateMaterialRequest) (resp *types.MaterialInfo, err error) {
	// 1. 根据 code 查找实体
	mat, err := l.svcCtx.MaterialRepo.FindByCode(l.ctx, req.Code)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, errcode.ErrMaterialNotFound
		}
		l.Logger.Errorf("FindByCode error: %v", err)
		return nil, errcode.ErrInternalError
	}

	// 2. "修补" (Patch) 实体字段
	// 注意: 这是 DTO 字段为 "optional" 但非指针的妥协处理
	if req.Name != "" {
		mat.Name = req.Name
	}
	if req.MaterialType != "" {
		mat.MaterialType = req.MaterialType
	}
	if req.Spec != "" {
		mat.Spec = req.Spec
	}
	if req.Unit != "" {
		mat.Unit = req.Unit
	}
	if req.SupplierName != "" {
		mat.SupplierName = req.SupplierName
	}
	// 对于数值，这种模式意味着你无法将值更新为 0
	if req.Price != 0 {
		mat.Price = req.Price
	}
	if req.StockQuantity != 0 {
		mat.StockQuantity = req.StockQuantity
	}

	// 3. 持久化更新
	if err := l.svcCtx.MaterialRepo.Update(l.ctx, mat); err != nil {
		l.Logger.Errorf("MaterialRepo.Update error: %v", err)
		return nil, errcode.ErrInternalError
	}

	// 4. 返回更新后的DTO
	return l.svcCtx.Converter.ToMaterialInfoResponse(mat), nil
}
