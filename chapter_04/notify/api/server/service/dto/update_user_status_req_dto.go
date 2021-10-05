package dto

import (
	"errors"
	"fmt"
	cErr "notify/api/server/service/error"
)

type UpdateUserStatusReqDto struct {
	UUID   string
	Status int
}

func (dto UpdateUserStatusReqDto) Check() error {
	if dto.UUID == "" {
		return errors.New(fmt.Sprintf("error:%v", cErr.WrapCommonError(cErr.ParamMiss, "u001")))
	}
	if dto.Status == 0 { //空值
		return errors.New(fmt.Sprintf("error:%v", cErr.WrapCommonError(cErr.ParamMiss, "u002")))
	}
	if dto.Status != 1 && dto.Status != 2 { //非法值
		return errors.New(fmt.Sprintf("error:%v", cErr.WrapCommonError(cErr.ParamMiss, "u003")))
	}
	return nil
}
