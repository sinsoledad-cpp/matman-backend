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

type DeleteMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除物料 (根据 Code)
func NewDeleteMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMaterialLogic {
	return &DeleteMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMaterialLogic) DeleteMaterial(req *types.MaterialRequest) (resp *types.GeneralSuccessResponse, err error) {
	// 1. (安全检查) 检查物料是否正在被 BOM 用作"子物料"
	inUse, err := l.svcCtx.MaterialRepo.IsMaterialInUse(l.ctx, req.Code)
	if err != nil {
		l.Logger.Errorf("IsMaterialInUse error: %v", err)
		return nil, errcode.ErrInternalError
	}
	if inUse {
		// 返回业务错误，阻止删除
		return nil, errcode.ErrMaterialInUse
	}

	// 2. 执行删除
	err = l.svcCtx.MaterialRepo.DeleteByCode(l.ctx, req.Code)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			// 如果没找到，也算删除成功 (幂等)
			return &types.GeneralSuccessResponse{Success: true}, nil
		}
		l.Logger.Errorf("DeleteByCode error: %v", err)
		return nil, errcode.ErrInternalError
	}

	return &types.GeneralSuccessResponse{Success: true}, nil
}
