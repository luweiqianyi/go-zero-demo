package logic

import (
	"context"
	"database/sql"
	"fmt"
	_const "go-zero-demo/cmd/account/internal/const"
	"go-zero-demo/cmd/account/internal/svc"
	"go-zero-demo/cmd/account/internal/types"
	"go-zero-demo/cmd/account/model"

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
	record := &model.TbUserAccount{
		AccountName: req.AccountName,
		Password: sql.NullString{
			String: req.Password,
			Valid:  true,
		},
	}
	result, err := l.svcCtx.TbUserAccountModel.Insert(l.ctx, record)

	resp = new(types.RegisterResp)
	if err != nil {
		resp.Result = _const.ApiFailed
		resp.Message = fmt.Sprintf("result: %#v, err:%v", result, err)
	} else {
		resp.Result = _const.ApiSuccess
	}
	return
}
