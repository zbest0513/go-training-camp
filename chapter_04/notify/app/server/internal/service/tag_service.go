package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"notify-server/internal/biz"
	"notify/api/server/service/dto"
	cErr "notify/api/server/service/error"
	"notify/pkg/utils"
)

type TagService struct {
	uc *biz.TagUseCase
}

func NewTagService(uc *biz.TagUseCase) *TagService {
	return &TagService{
		uc: uc,
	}
}

func (ts *TagService) tag2DTO(tag *biz.Tag) dto.TagDto {
	return dto.TagDto{
		UUID:       tag.Uuid,
		Name:       tag.Name,
		Desc:       tag.Desc,
		Status:     tag.Status,
		CreateTime: utils.Time2DateStr(tag.CreatedAt),
		UpdateTime: utils.Time2DateStr(tag.UpdatedAt),
	}
}

func (ts *TagService) CreateTag(c *gin.Context) {
	var req dto.CreateTagReqDto
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteErr(c, cErr.NewBError(http.StatusInternalServerError, cErr.ParamBindError))
		return
	}
	if err := req.Check(); err != nil {
		bError := cErr.NewBError(http.StatusInternalServerError, cErr.
			WrapCommonError(cErr.ParamError, err.Error()))
		utils.WriteErr(c, bError)
		return
	}

	ctx := context.WithValue(context.Background(), "ginContext", c)
	//DTO => DO
	tag := biz.Tag{
		Name: req.Name,
		Desc: req.Desc,
		Uuid: utils.UUID(),
	}
	result, err := ts.uc.CreateTag(ctx, tag)

	if err != nil {
		log.Println(fmt.Sprintf("错误:%+v", err))
		bError := cErr.NewBError(http.StatusInternalServerError, cErr.
			WrapCommonError(cErr.SystemError, "创建标签失败"))
		utils.WriteErr(c, bError)
		return
	}

	//DO => DTO
	utils.WriteData(c, ts.tag2DTO(result))
	return
}

func (ts *TagService) DeleteTag(c *gin.Context) {
	var req dto.DeleteTagReqDto
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteErr(c, cErr.NewBError(http.StatusInternalServerError, cErr.ParamBindError))
		return
	}
	if err := req.Check(); err != nil {
		bError := cErr.NewBError(http.StatusInternalServerError, cErr.
			WrapCommonError(cErr.ParamError, err.Error()))
		utils.WriteErr(c, bError)
		return
	}

	ctx := context.WithValue(context.Background(), "ginContext", c)

	err := ts.uc.DeleteTag(ctx, req.UUID)
	if err != nil {
		log.Println(fmt.Sprintf("错误:%+v", err))
		bError := cErr.NewBError(http.StatusInternalServerError, cErr.
			WrapCommonError(cErr.SystemError, "删除标签失败"))
		utils.WriteErr(c, bError)
		return
	}

	//DO => DTO
	utils.WriteData(c, "删除成功")
	return
}
