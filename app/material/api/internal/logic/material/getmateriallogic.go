// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"context"

	"matman-backend/app/material/api/internal/svc"
	"matman-backend/app/material/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取单个物料详情 (根据 Code)
func NewGetMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMaterialLogic {
	return &GetMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMaterialLogic) GetMaterial(req *types.MaterialRequest) (resp *types.MaterialInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
