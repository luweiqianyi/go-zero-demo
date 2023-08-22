package logic

import (
	"context"
	"fmt"
	"go-zero-demo/cmd/account-rpc/accountrpcservice"
	_const "go-zero-demo/cmd/account/internal/const"

	"go-zero-demo/cmd/account/internal/svc"
	"go-zero-demo/cmd/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorizationLogic {
	return &AuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthorizationLogic) Authorization(req *types.AuthorizationReq) (resp *types.AuthorizationResp, err error) {
	resp = new(types.AuthorizationResp)

	_, err = l.svcCtx.AccountRpcClient.ValidateToken(
		context.Background(),
		&accountrpcservice.TokenValidateReq{
			Token: req.AccessToken,
		})
	if err != nil {
		resp.Result = _const.ApiFailed
		resp.Message = fmt.Sprintf("auth failed, err: %v", err)
		return
	}
	resp.Result = _const.ApiSuccess

	return
}
