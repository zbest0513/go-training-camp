package dto

import (
	"errors"
	"fmt"
	cerror "notify/api/server/service/error"
)

type CreateUserReqDto struct {
	UUID   string
	Name   string
	Mobile string
	Email  string
}

func (dto CreateUserReqDto) Check() error {
	if dto.Name == "" {
		return errors.New(fmt.Sprintf("error:%v", cerror.WrapCommonError(cerror.ParamMiss, "u002")))
	}
	if dto.Email == "" {
		return errors.New(fmt.Sprintf("error:%v", cerror.WrapCommonError(cerror.ParamMiss, "u003")))
	}
	if dto.Mobile == "" {
		return errors.New(fmt.Sprintf("error:%v", cerror.WrapCommonError(cerror.ParamMiss, "u004")))
	}
	if len(dto.Mobile) != 11 {
		return errors.New(fmt.Sprintf("error:%v", cerror.WrapCommonError(cerror.ParamMiss, "u005")))
	}
	return nil
}
