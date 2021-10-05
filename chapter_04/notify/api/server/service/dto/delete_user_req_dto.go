package dto

import (
	"errors"
	"fmt"
	cerror "notify/api/server/service/error"
)

type DeleteTagReqDto struct {
	UUID string
}

func (dto DeleteTagReqDto) Check() error {
	if dto.UUID == "" {
		return errors.New(fmt.Sprintf("error:%v", cerror.WrapCommonError(cerror.ParamMiss, "t001")))
	}
	return nil
}
