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

type ListBomsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 检索BOM列表 (支持父物料名称、状态分页查询)
func NewListBomsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBomsLogic {
	return &ListBomsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBomsLogic) ListBoms(req *types.ListBomsRequest) (resp *types.ListBomsResponse, err error) {
	// 1. 处理分页参数
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

	// 2. 调用仓储层 (Repository)
	// (传入父物料名称过滤器 和 状态过滤器)
	detailsList, total, err := l.svcCtx.BomRepo.ListDetails(l.ctx, page, pageSize, req.ParentName, req.Status)
	if err != nil {
		l.Logger.Errorf("BomRepo.ListDetails error: %v", err)
		return nil, errcode.ErrInternalError
	}

	// 3. 将领域读模型(Entity)列表 转换为 DTO 列表 (使用 Converter)
	dtoEntries := l.svcCtx.Converter.ToBomListEntryList(detailsList)

	// 4. 返回响应
	return &types.ListBomsResponse{
		Entries: dtoEntries,
		Total:   total,
	}, nil
}
