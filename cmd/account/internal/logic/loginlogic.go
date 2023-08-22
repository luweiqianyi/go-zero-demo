package logic

import (
	"context"
	"fmt"
	"go-zero-demo/cmd/account-rpc/accountrpcservice"
	_const "go-zero-demo/cmd/account/internal/const"
	"go-zero-demo/pkg/cryptx"
	"time"

	"go-zero-demo/cmd/account/internal/svc"
	"go-zero-demo/cmd/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	TokenExpireTime = time.Hour * 24 * 15
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	encryptedPassword := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
	record, err := l.svcCtx.TbUserAccountModel.FindOne(l.ctx, req.AccountName)

	resp = new(types.LoginResp)
	if err != nil {
		resp.Result = _const.ApiFailed
		resp.Message = fmt.Sprintf("login failed, err: %v", err)
		return
	}

	if record.Password.String != encryptedPassword {
		resp.Result = _const.ApiFailed
		resp.Message = fmt.Sprintf("login failed, password wrong")
		return
	}

	generateTokenResp, err := l.svcCtx.AccountRpcClient.GenerateToken(
		context.Background(),
		&accountrpcservice.GenerateTokenReq{
			AccountName: req.AccountName,
		},
	)
	if err != nil {
		resp.Result = _const.ApiFailed
		resp.Message = fmt.Sprintf("token generate failed, err: %v", err)
		return
	}

	resp.Result = _const.ApiSuccess
	resp.Token = generateTokenResp.Token
	return
}
