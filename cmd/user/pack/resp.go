package pack

import (
	"errors"

	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *user.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// BuildCommonResp build commonResp from error
// func BuildCommonResp(err error) *user.BaseResp {}
