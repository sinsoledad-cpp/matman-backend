// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
