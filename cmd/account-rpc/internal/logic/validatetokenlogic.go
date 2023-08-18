package logic

import (
	"context"
	"errors"
	"go-zero-demo/pkg/store"

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

func (l *ValidateTokenLogic) ValidateToken(in *pb.TokenValidateReq) (*pb.TokenValidateResp, error) {
	token, err := store.Get(in.AccountName)
	if err != nil {
		return &pb.TokenValidateResp{
			Ok: false,
		}, nil
	}

	if token != in.Token {
		return &pb.TokenValidateResp{
			Ok: false,
		}, errors.New("token error")
	}

	return &pb.TokenValidateResp{
		Ok: true,
	}, nil
}
