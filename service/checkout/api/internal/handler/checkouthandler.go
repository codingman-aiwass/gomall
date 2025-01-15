package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gomall/service/checkout/api/internal/logic"
	"gomall/service/checkout/api/internal/svc"
	"gomall/service/checkout/api/internal/types"
)

func CheckoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckoutRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCheckoutLogic(r.Context(), svcCtx)
		resp, err := l.Checkout(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
