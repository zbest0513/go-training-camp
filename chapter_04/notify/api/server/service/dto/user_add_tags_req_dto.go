package dto

import (
	"errors"
	"fmt"
	cErr "notify/api/server/service/error"
	"strconv"
)

type UserAddTagsReqDto struct {
	UserUuid string
	TagUuids []string
}

func (dto UserAddTagsReqDto) Check() error {
	if dto.UserUuid == "" {
		return errors.New(fmt.Sprintf("error:%v", cErr.WrapCommonError(cErr.ParamMiss, "u001")))
	}
	if len(dto.TagUuids) == 0 { //空值
		return errors.New(fmt.Sprintf("error:%v", cErr.WrapCommonError(cErr.ParamMiss, "u002")))
	}
	for i, uuid := range dto.TagUuids {
		if uuid == "" {
			return errors.New(fmt.Sprintf("error:%v", cErr.
				WrapCommonError(cErr.ParamMiss, "u003"+"_"+strconv.Itoa(i))))
		}
	}
	return nil
}
