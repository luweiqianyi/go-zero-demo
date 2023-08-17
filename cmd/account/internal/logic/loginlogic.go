package logic

import (
	"context"
	"fmt"
	_const "go-zero-demo/cmd/account/internal/const"
	"go-zero-demo/pkg/cryptx"
	"go-zero-demo/pkg/token"

	"go-zero-demo/cmd/account/internal/svc"
	"go-zero-demo/cmd/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	} else {
		if record.Password.String == encryptedPassword {
			resp.Result = _const.ApiSuccess
			resp.Token = token.GenerateToken()
		} else {
			resp.Result = _const.ApiFailed
			resp.Message = fmt.Sprintf("login failed, password wrong")
		}
	}
	return
}
