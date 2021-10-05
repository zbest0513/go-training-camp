package service

import "github.com/gin-gonic/gin"
import _ "notify/doc"

type UserApi interface {
	// CreateUser method param use dto.CreateUserReqDto
	// @Summary 创建用户
	// @Description 创建用户
	// @Accept  json
	// @Produce  json
	// @Param   Mobile	string	string	true	"手机号"
	// @Param   Email	string	string	true	"邮箱"
	// @Param   Name	string	string	true	"姓名"
	// @Success 200 {object} dto.UserDto    "ok"
	CreateUser(*gin.Context)
	// UpdateUserStatus method param use dto.UpdateUserStatusReqDto
	UpdateUserStatus(*gin.Context)
	// AllUsers 查询所有用户
	AllUsers(c *gin.Context)
	// AddTags 给用户添加标签
	AddTags(c *gin.Context)
}

type TagApi interface {
	CreateTag(c *gin.Context)
	DeleteTag(c *gin.Context)
}

type TemplateApi interface {
	CreateTemplate(c *gin.Context)
}
