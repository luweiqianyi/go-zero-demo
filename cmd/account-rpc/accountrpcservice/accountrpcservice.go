// Code generated by goctl. DO NOT EDIT.
// Source: account.proto

package accountrpcservice

import (
	"context"

	"go-zero-demo/cmd/account-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GenerateTokenReq  = pb.GenerateTokenReq
	GenerateTokenResp = pb.GenerateTokenResp
	TokenValidateReq  = pb.TokenValidateReq
	TokenValidateResp = pb.TokenValidateResp

	AccountRpcService interface {
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		ValidateToken(ctx context.Context, in *TokenValidateReq, opts ...grpc.CallOption) (*TokenValidateResp, error)
	}

	defaultAccountRpcService struct {
		cli zrpc.Client
	}
)

func NewAccountRpcService(cli zrpc.Client) AccountRpcService {
	return &defaultAccountRpcService{
		cli: cli,
	}
}

func (m *defaultAccountRpcService) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewAccountRpcServiceClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultAccountRpcService) ValidateToken(ctx context.Context, in *TokenValidateReq, opts ...grpc.CallOption) (*TokenValidateResp, error) {
	client := pb.NewAccountRpcServiceClient(m.cli.Conn())
	return client.ValidateToken(ctx, in, opts...)
}
