package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"notify-server/internal/biz"
	"notify/api/server/service/dto"
	cErr "notify/api/server/service/error"
	"notify/pkg/utils"
)

type UserService struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{
		uc: uc,
	}
}
func (us *UserService) user2DTO(user *biz.User) dto.UserDto {
	return dto.UserDto{
		UUID:       user.Uuid,
		Name:       user.Name,
		Email:      user.Email,
		Status:     user.Status,
		Mobile:     user.Mobile,
		CreateTime: utils.Time2DateStr(user.CreatedAt),
		UpdateTime: utils.Time2DateStr(user.UpdatedAt),
	}
}

func (us *UserService) CreateUser(c *gin.Context) {
	var req dto.CreateUserReqDto
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
	user := biz.User{
		Name:   req.Name,
		Mobile: req.Mobile,
		Email:  req.Email,
		Uuid:   uuid.New().String(),
	}
	result, err := us.uc.CreateUser(ctx, user)

	if err != nil {
		log.Println(fmt.Sprintf("错误:%+v", err))
		bError := cErr.NewBError(http.StatusInternalServerError, cErr.
			WrapCommonError(cErr.SystemError, "创建用户失败"))
		utils.WriteErr(c, bError)
		return
	}

	//DO => DTO
	utils.WriteData(c, us.user2DTO(result))
	return
}

func (us *UserService) AllUsers(c *gin.Context) {
	ctx := context.WithValue(context.Background(), "ginContext", c)
	result, err := us.uc.QueryAll(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("错误:%+v", err))
		bError := cErr.NewBError(http.StatusInternalServerError, cErr.
			WrapCommonError(cErr.SystemError, "查询用户列表失败"))
		utils.WriteErr(c, bError)
		return
	}
	list := make([]dto.UserDto, len(result))
	for i, user := range result {
		list[i] = us.user2DTO(user)
	}
	//DO => DTO
	utils.WriteData(c, list)
}

func (us *UserService) UpdateUserStatus(c *gin.Context) {
	var req dto.UpdateUserStatusReqDto
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

	err := us.uc.UpdateUserStatus(ctx, req.UUID, req.Status)
	if err != nil {
		log.Println(fmt.Sprintf("错误:%+v", err))
		bError := cErr.NewBError(http.StatusInternalServerError, cErr.
			WrapCommonError(cErr.SystemError, "修改用户状态失败"))
		utils.WriteErr(c, bError)
		return
	}
	//DO => DTO
	utils.WriteData(c, "success")
	return
}

func (us *UserService) AddTags(c *gin.Context) {
	var req dto.UserAddTagsReqDto
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
	err := us.uc.AddTags(ctx, req.UserUuid, req.TagUuids)
	if err != nil {
		log.Println(fmt.Sprintf("用户添加标签失败:%v", err))
		bError := cErr.NewBError(http.StatusInternalServerError, cErr.
			WrapCommonError(cErr.SystemError, "用户添加标签失败"))
		utils.WriteErr(c, bError)
		return
	}
	utils.WriteData(c, "success")
}
