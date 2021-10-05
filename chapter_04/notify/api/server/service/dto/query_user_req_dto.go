package dto

import (
	"errors"
	"fmt"
	cerror "notify/api/server/service/error"
)

type QueryUserByUUIDReqDto struct {
	UUID string
}

func (dto QueryUserByUUIDReqDto) Check() error {
	if dto.UUID == "" {
		return errors.New(fmt.Sprintf("error:%v", cerror.WrapCommonError(cerror.ParamMiss, "u001")))
	}
	return nil
}
