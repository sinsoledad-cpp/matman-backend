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
	// 1. (安全检查) 确保父物料存在
	_, err = l.svcCtx.MaterialRepo.FindByCode(l.ctx, req.ParentCode)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, errcode.ErrMaterialNotFound
		}
		l.Logger.Errorf("FindByCode (Parent) error: %v", err)
		return nil, errcode.ErrInternalError
	}

	// 2. 获取BOM实体列表
	bomEntities, err := l.svcCtx.BomRepo.FindByParentCode(l.ctx, req.ParentCode)
	if err != nil {
		// (Repo 已处理 NotFound，这里只可能是数据库错误)
		l.Logger.Errorf("FindByParentCode error: %v", err)
		return nil, errcode.ErrInternalError
	}

	// 3. 转换为 DTO 并填充子物料信息
	dtos := make([]types.BomEntryInfo, 0, len(bomEntities))

	for _, e := range bomEntities {
		// 3.1 基础转换
		dto := l.svcCtx.Converter.ToBomEntryInfoResponse(e)

		// 3.2 跨聚合查询，填充子物料的名称和规格
		childMat, err := l.svcCtx.MaterialRepo.FindByCode(l.ctx, e.ChildMaterialCode)
		if err == nil {
			dto.ChildMaterialName = childMat.Name
			dto.ChildMaterialSpec = childMat.Spec
		} else {
			// 如果子物料被异常删除，标记出来
			dto.ChildMaterialName = "[物料不存在]"
			l.Logger.Errorf("GetBom: Child material %s not found for bom id %d", e.ChildMaterialCode, e.ID)
		}

		dtos = append(dtos, *dto)
	}

	return &types.GetBomResponse{
		Entries: dtos,
	}, nil
}
