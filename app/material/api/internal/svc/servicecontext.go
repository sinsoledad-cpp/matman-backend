// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package svc

import (
	"matman-backend/app/material/api/internal/config"
	"matman-backend/app/material/api/internal/logic"
	"matman-backend/app/material/data/model"
	"matman-backend/app/material/domain/repository"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	Converter    *logic.Converter              // (新增) 转换器
	MaterialRepo repository.MaterialRepository // (新增) 物料仓储接口
	BomRepo      repository.BomRepository      // (新增) BOM仓储接口
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Database.DataSource)

	materialsModel := model.NewMaterialsModel(conn)
	bomModel := model.NewBomModel(conn)

	materialRepo := repository.NewMaterialRepoImpl(materialsModel, bomModel)
	bomRepo := repository.NewBomRepoImpl(bomModel)

	// 4. 初始化 converter
	converter := logic.NewConverter()

	return &ServiceContext{
		Config:       c,
		Converter:    converter,
		MaterialRepo: materialRepo, // 注入 repo 实现
		BomRepo:      bomRepo,      // 注入 repo 实现
	}
}
