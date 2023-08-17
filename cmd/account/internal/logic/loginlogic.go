package logic

import (
	"context"
	"fmt"
	_const "go-zero-demo/cmd/account/internal/const"
	"go-zero-demo/pkg/cryptx"
	"go-zero-demo/pkg/token"
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

type TokenData struct {
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
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

	tokenData := TokenData{
		AccountName: req.AccountName,
		Password:    req.Password,
	}
	accessToken, err := token.GenerateToken(l.svcCtx.Config.TokenSecretKey, tokenData, TokenExpireTime)
	if err != nil {
		resp.Result = _const.ApiFailed
		resp.Message = fmt.Sprintf("token generate failed, err: %v", err)
		return
	}

	resp.Result = _const.ApiSuccess
	resp.Token = accessToken
	return
}
