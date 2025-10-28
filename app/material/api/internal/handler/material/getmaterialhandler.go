// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package material

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"matman-backend/app/material/api/internal/logic/material"
	"matman-backend/app/material/api/internal/svc"
	"matman-backend/app/material/api/internal/types"

	"matman-backend/common/utils/response"
)

// 获取单个物料详情 (根据 Code)
func GetMaterialHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MaterialRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ClientError(r.Context(), w, response.RequestError, err.Error())
			return
		}

		l := material.NewGetMaterialLogic(r.Context(), svcCtx)
		resp, err := l.GetMaterial(&req)
		if err != nil {
			response.LogicError(r.Context(), w, err)
		} else {
			response.Ok(r.Context(), w, resp)
		}
	}
}
