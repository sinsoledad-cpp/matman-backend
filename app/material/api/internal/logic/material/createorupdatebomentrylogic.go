// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
