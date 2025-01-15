package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/service/auth/rpc/authservice"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	AuthRpc authservice.AuthService
}

func NewAuthMiddleware(cli zrpc.Client) *AuthMiddleware {
	return &AuthMiddleware{AuthRpc: authservice.NewAuthService(cli)}
}
func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取 Authorization
		authHeader := r.Header.Get("Authorization")
		logx.Infof("authHeader: %s", authHeader)
		if authHeader == "" {
			logx.Errorf("Authorization header not found")
			return
		}

		// 提取 Token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			logx.Errorf("Invalid Authorization header")
			return
		}

		// 调用 Auth 微服务验证 Token
		resp, err := m.AuthRpc.VerifyTokenByRPC(context.Background(), &authservice.VerifyTokenReq{
			Token:     token,
			TokenType: "access",
		})
		logx.Infof("resp: %v", resp)
		if err != nil || !resp.Res {
			logx.Errorf("Token verification failed: %s", err)
			return
		}
		logx.Infof("Token verification passed")
		logx.Infof("userId: %d", resp.UserId)
		// 将验证通过的用户 ID 写入请求上下文
		ctx := context.WithValue(r.Context(), "userId", int64(resp.UserId))
		r = r.WithContext(ctx)
		next(w, r)
	}
}
