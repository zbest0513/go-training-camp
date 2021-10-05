package dto

import (
	"errors"
	"fmt"
	cerror "notify/api/server/service/error"
)

type CreateTemplateReqDto struct {
	Name    string
	Desc    string
	Content string
}

func (dto CreateTemplateReqDto) Check() error {
	if dto.Name == "" {
		return errors.New(fmt.Sprintf("error:%v", cerror.WrapCommonError(cerror.ParamMiss, "t001")))
	}
	if dto.Desc == "" {
		return errors.New(fmt.Sprintf("error:%v", cerror.WrapCommonError(cerror.ParamMiss, "t002")))
	}
	if dto.Content == "" {
		return errors.New(fmt.Sprintf("error:%v", cerror.WrapCommonError(cerror.ParamMiss, "t003")))
	}
	return nil
}
