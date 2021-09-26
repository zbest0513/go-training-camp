package service

import "github.com/gin-gonic/gin"

type UserDto struct {
	UUID   string
	Name   string
	Mobile string
	Email  string
	Status int8
}

type CreateUserReqDto struct {
	UUID   string
	Name   string
	Mobile string
	Email  string
}

type QueryUserReqDto struct {
	UUID string
}

type ManageApi interface {
	CreateUser(*gin.Context)
	//QueryUser(*gin.Engine,QueryUserReqDto) UserDto
}
