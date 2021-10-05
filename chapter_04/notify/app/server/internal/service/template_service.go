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

type TemplateService struct {
	uc *biz.TemplateUseCase
}

func NewTemplateService(uc *biz.TemplateUseCase) *TemplateService {
	return &TemplateService{
		uc: uc,
	}
}

func (ts *TemplateService) template2DTO(template *biz.Template) dto.TemplateDto {
	return dto.TemplateDto{
		UUID:       template.Uuid,
		Name:       template.Name,
		Desc:       template.Desc,
		Content:    template.Content,
		Status:     template.Status,
		CreateTime: utils.Time2DateStr(template.CreatedAt),
		UpdateTime: utils.Time2DateStr(template.UpdatedAt),
	}
}

func (ts *TemplateService) CreateTemplate(c *gin.Context) {
	var req dto.CreateTemplateReqDto
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
	template := biz.Template{
		Name:    req.Name,
		Desc:    req.Desc,
		Uuid:    utils.UUID(),
		Content: req.Content,
	}
	result, err := ts.uc.CreateTemplate(ctx, template)

	if err != nil {
		log.Println(fmt.Sprintf("错误:%+v", err))
		bError := cErr.NewBError(http.StatusInternalServerError, cErr.
			WrapCommonError(cErr.SystemError, "创建模版失败"))
		utils.WriteErr(c, bError)
		return
	}

	//DO => DTO
	utils.WriteData(c, ts.template2DTO(result))
	return
}
