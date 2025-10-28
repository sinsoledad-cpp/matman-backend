// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
