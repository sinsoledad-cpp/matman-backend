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

type DeleteBomEntryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除一个BOM条目
func NewDeleteBomEntryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBomEntryLogic {
	return &DeleteBomEntryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBomEntryLogic) DeleteBomEntry(req *types.DeleteBomEntryRequest) (resp *types.GeneralSuccessResponse, err error) {
	err = l.svcCtx.BomRepo.Delete(l.ctx, req.ParentCode, req.ChildCode)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			// 未找到，也视为成功 (幂等)
			return &types.GeneralSuccessResponse{Success: true}, nil
		}
		l.Logger.Errorf("BomRepo.Delete error: %v", err)
		return nil, errcode.ErrInternalError
	}

	return &types.GeneralSuccessResponse{Success: true}, nil
}
