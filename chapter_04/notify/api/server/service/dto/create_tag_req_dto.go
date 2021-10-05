package dto

import (
	"errors"
	"fmt"
	cerror "notify/api/server/service/error"
)

type CreateTagReqDto struct {
	Name string
	Desc string
}

func (dto CreateTagReqDto) Check() error {
	if dto.Name == "" {
		return errors.New(fmt.Sprintf("error:%v", cerror.WrapCommonError(cerror.ParamMiss, "t001")))
	}
	if dto.Desc == "" {
		return errors.New(fmt.Sprintf("error:%v", cerror.WrapCommonError(cerror.ParamMiss, "t002")))
	}
	return nil
}
