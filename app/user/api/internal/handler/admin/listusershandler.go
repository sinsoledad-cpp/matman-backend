// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"matman-backend/app/user/api/internal/logic/admin"
	"matman-backend/app/user/api/internal/svc"
	"matman-backend/app/user/api/internal/types"

	"matman-backend/common/utils/response"
)

// 获取所有用户列表 (管理员)
func ListUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListUsersRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ClientError(r.Context(), w, response.RequestError, err.Error())
			return
		}

		l := admin.NewListUsersLogic(r.Context(), svcCtx)
		resp, err := l.ListUsers(&req)
		if err != nil {
			response.LogicError(r.Context(), w, err)
		} else {
			response.Ok(r.Context(), w, resp)
		}
	}
}
