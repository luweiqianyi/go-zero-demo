package logic

import (
	"context"
	"fmt"
	"go-zero-demo/cmd/account-rpc/entity"
	"go-zero-demo/cmd/account-rpc/internal/svc"
	"go-zero-demo/cmd/account-rpc/pb"
	"go-zero-demo/cmd/account-rpc/store"
	"go-zero-demo/pkg/token"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	TokenExpireSecond = 1
	TokenExpireHour   = TokenExpireSecond * 3600
	TokenExpire15Day  = TokenExpireHour * 24 * 15
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	tokenData := entity.TokenData{
		AccountName: in.AccountName,
	}
	accessToken, err := token.GenerateToken(l.svcCtx.Config.TokenSecretKey, tokenData, time.Second*TokenExpire15Day)
	if err != nil {
		return &pb.GenerateTokenResp{
			Success: false,
			Token:   "",
		}, fmt.Errorf("token generate failed, err: %v", err)
	}

	err = store.SaveTokenByAccountName(in.AccountName, accessToken, TokenExpire15Day)
	if err != nil {
		return &pb.GenerateTokenResp{
			Success: false,
			Token:   "",
		}, fmt.Errorf("token generate failed, err: %v", err)
	}

	return &pb.GenerateTokenResp{
		Success: true,
		Token:   accessToken,
	}, nil
}
