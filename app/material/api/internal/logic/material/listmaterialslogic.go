// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"context"
	"matman-backend/app/material/api/internal/logic/errcode"

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
	// 1. 处理分页参数 (模仿 ListUsersLogic)
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	switch {
	case pageSize <= 0:
		pageSize = 20 // 默认值
	case pageSize > 100:
		pageSize = 100 // 最大值
	}

	// 2. 调用仓储层
	// (req.Name 用于模糊查询，在 repo impl 中处理)
	entities, total, err := l.svcCtx.MaterialRepo.List(l.ctx, page, pageSize, req.Name)
	if err != nil {
		l.Logger.Errorf("MaterialRepo.List error: %v", err)
		return nil, errcode.ErrInternalError
	}

	// 3. 将领域实体(Entity)列表转换为 DTO 列表
	dtoMaterials := l.svcCtx.Converter.ToMaterialInfoList(entities)

	// 4. 返回响应
	return &types.ListMaterialsResponse{
		Materials: dtoMaterials,
		Total:     total,
	}, nil
}
