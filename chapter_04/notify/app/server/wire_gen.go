// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"notify-server/internal/biz"
	"notify-server/internal/data"
	"notify-server/internal/service/handle"
	"notify-server/internal/service/router"
	"notify/pkg/config"
)

// Injectors from wire.go:

func initApp(fc config.FileConfig) (*router.Router, error) {
	configConfig, err := data.NewConfig(fc)
	if err != nil {
		return nil, err
	}
	client := data.NewEntClient(configConfig)
	userRepo := data.NewUserRepo(client)
	userUseCase := biz.NewUserUseCase(userRepo)
	userService := handle.NewUserService(userUseCase)
	routerRouter := &router.Router{
		ManageApi: userService,
	}
	return routerRouter, nil
}