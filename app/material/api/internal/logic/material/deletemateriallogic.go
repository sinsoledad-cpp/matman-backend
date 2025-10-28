// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
