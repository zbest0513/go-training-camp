package handle

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"notify-server/internal/biz"
)
import api "notify/api/server/service"

type UserService struct {
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{
		uc: uc,
	}
}

func (us *UserService) CreateUser(c *gin.Context) {
	var dto api.CreateUserReqDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.WithValue(context.Background(), "ginContext", c)
	//DTO => DO
	user := biz.User{
		Name:   dto.Name,
		Mobile: dto.Mobile,
		Email:  dto.Email,
	}
	result, err := us.uc.CreateUser(ctx, user)

	if err != nil {
		log.Println(fmt.Sprintf("错误:%+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//DO => DTO
	c.JSON(http.StatusOK, gin.H{"data": api.UserDto{
		UUID:   result.Uuid,
		Name:   result.Name,
		Email:  result.Email,
		Status: result.Status,
		Mobile: result.Mobile,
	}})
	return
}

func (us *UserService) UpdateUserStatus(c *gin.Context) {
	var dto api.UpdateUserStatusReqDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.WithValue(context.Background(), "ginContext", c)
	//DTO => DO
	user := biz.User{
		Status: (int8)(dto.Status),
		Uuid:   dto.UUID,
	}
	err := us.uc.UpdateUserStatus(ctx, user)
	if err != nil {
		log.Println(fmt.Sprintf("错误:%+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//DO => DTO
	c.JSON(http.StatusOK, gin.H{"data": "success"})
	return
}
