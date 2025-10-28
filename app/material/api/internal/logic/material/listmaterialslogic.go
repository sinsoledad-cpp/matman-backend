// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"context"

	"matman-backend/app/material/api/internal/svc"
	"matman-backend/app/material/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMaterialsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取物料列表 (分页/查询)
func NewListMaterialsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMaterialsLogic {
	return &ListMaterialsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMaterialsLogic) ListMaterials(req *types.ListMaterialsRequest) (resp *types.ListMaterialsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
