package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gomall/service/cart/api/internal/logic"
	"gomall/service/cart/api/internal/svc"
	"gomall/service/cart/api/internal/types"
)

func EmptyCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyCartRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewEmptyCartLogic(r.Context(), svcCtx)
		resp, err := l.EmptyCart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
