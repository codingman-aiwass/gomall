package logic

import (
	"context"
	"strings"

	"gomall/service/auth/rpc/internal/svc"
	"gomall/service/auth/rpc/types/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyPathInWhiteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyPathInWhiteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyPathInWhiteListLogic {
	return &VerifyPathInWhiteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func isInWriteList(resource string, whiteList []string) bool {
	for _, path := range whiteList {
		if strings.HasPrefix(resource, path) || resource == path {
			return true
		}
	}
	return false
}
func (l *VerifyPathInWhiteListLogic) VerifyPathInWhiteList(in *auth.VerifyPathInWhiteListReq) (*auth.VerifyPathInWhiteListResp, error) {
	ok := isInWriteList(in.Path, l.svcCtx.WhiteList)
	return &auth.VerifyPathInWhiteListResp{Res: ok}, nil
}
