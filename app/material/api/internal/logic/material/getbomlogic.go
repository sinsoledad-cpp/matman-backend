// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"context"

	"matman-backend/app/material/api/internal/svc"
	"matman-backend/app/material/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取一个父物料的BOM清单
func NewGetBomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBomLogic {
	return &GetBomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBomLogic) GetBom(req *types.GetBomRequest) (resp *types.GetBomResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
