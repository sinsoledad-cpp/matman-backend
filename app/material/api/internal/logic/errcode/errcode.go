package errcode

import "matman-backend/common/utils/response"

// 业务错误实例
// Logic 层将直接返回这些预先定义好的错误
var (
	// --- 物料 (20001+) ---
	ErrMaterialCodeExists   = response.NewBizError(20001, "物料编码已存在")
	ErrMaterialNotFound     = response.NewBizError(20002, "物料未找到")
	ErrMaterialInUse        = response.NewBizError(20003, "物料正在被BOM使用，无法删除")
	ErrBomEntryExists       = response.NewBizError(20004, "BOM条目已存在 (父物料和子物料组合唯一)")
	ErrBomEntryNotFound     = response.NewBizError(20005, "BOM条目未找到")
	ErrBomCannotSelfContain = response.NewBizError(20006, "BOM父物料和子物料不能相同")

	// --- 通用 ---
	ErrInternalError = response.NewBizError(50000, "服务器内部错误")
)
