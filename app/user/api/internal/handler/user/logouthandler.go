// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.1

package user

import (
	"net/http"

	"matman-backend/app/user/api/internal/logic/user"
	"matman-backend/app/user/api/internal/svc"

	"matman-backend/common/utils/response"
)

// 用户登出
func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
		if err != nil {
			response.LogicError(r.Context(), w, err)
		} else {
			response.Ok(r.Context(), w, nil)
		}
	}
}
