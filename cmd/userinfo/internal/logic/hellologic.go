package logic

import (
	"context"

	"go-zero-demo/cmd/userinfo/internal/svc"
	"go-zero-demo/cmd/userinfo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req *types.HelloRequest) (resp *types.HelloResponse, err error) {
	resp = new(types.HelloResponse)
	resp.Message = "Hello, visitor!"
	return
}
