package logic

import (
	"context"
	_const "go-zero-demo/cmd/account/internal/const"
	"go-zero-demo/cmd/account/internal/service"

	"go-zero-demo/cmd/account/internal/svc"
	"go-zero-demo/cmd/account/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	err = service.Register(req.AccountName, req.Password)
	resp = new(types.RegisterResp)
	if err != nil {
		resp.Result = _const.ApiFailed
		resp.Message = err.Error()
	} else {
		resp.Result = _const.ApiSuccess
	}
	return
}
