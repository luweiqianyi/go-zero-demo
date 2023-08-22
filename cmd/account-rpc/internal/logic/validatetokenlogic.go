package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"go-zero-demo/cmd/account-rpc/entity"
	"go-zero-demo/cmd/account-rpc/store"
	"go-zero-demo/pkg/token"

	"go-zero-demo/cmd/account-rpc/internal/svc"
	"go-zero-demo/cmd/account-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ValidateToken TODO: 优化Token验证逻辑
func (l *ValidateTokenLogic) ValidateToken(in *pb.TokenValidateReq) (*pb.TokenValidateResp, error) {
	data, err := token.ParseToken(in.Token, l.svcCtx.Config.TokenSecretKey)
	if err != nil {
		return &pb.TokenValidateResp{
			Ok: false,
		}, fmt.Errorf("validate token failed, token parse error")
	}

	strData := data.(string)
	tokenData := &entity.TokenData{}
	err = json.Unmarshal([]byte(strData), tokenData)
	if err != nil {
		return nil, fmt.Errorf("validate token failed, token format error")
	}

	tokenStore, err := store.FindTokenByAccountName(tokenData.AccountName)
	if err != nil {
		return &pb.TokenValidateResp{
			Ok: false,
		}, fmt.Errorf("validate token failed, err:%v", err)
	}

	if tokenStore != in.Token {
		return &pb.TokenValidateResp{
			Ok: false,
		}, fmt.Errorf("validate token failed, err: invalid token")
	}

	return &pb.TokenValidateResp{
		Ok: true,
	}, nil
}
