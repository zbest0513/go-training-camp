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

type UpdateUserStatusReqDto struct {
	UUID   string
	Status int
}

type ManageApi interface {
	// CreateUser method param use CreateUserReqDto
	CreateUser(*gin.Context)
	// UpdateUserStatus method param use UpdateUserStatusReqDto
	UpdateUserStatus(*gin.Context)
}
