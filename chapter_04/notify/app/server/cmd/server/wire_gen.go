// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"notify-server/internal/biz"
	"notify-server/internal/data"
	"notify-server/internal/pkg"
	"notify-server/internal/service"
	"notify-server/internal/service/router"
	"notify/pkg/config"
)

// Injectors from wire.go:

func initApp(fc config.FileConfig) (*pkg.App, error) {
	configConfig, err := data.NewConfig(fc)
	if err != nil {
		return nil, err
	}
	client := data.NewEntClient(configConfig)
	userRepo := data.NewUserRepo(client)
	userUseCase := biz.NewUserUseCase(userRepo)
	userService := service.NewUserService(userUseCase)
	monitorRouter := &router.DocRouter{
		UserApi: userService,
	}
	tagRepo := data.NewTagRepo(client)
	tagUseCase := biz.NewTagUseCase(tagRepo)
	tagService := service.NewTagService(tagUseCase)
	templateRepo := data.NewTemplateRepo(client)
	templateUseCase := biz.NewTemplateUseCase(templateRepo, tagRepo)
	templateService := service.NewTemplateService(templateUseCase)
	manageRouter := &router.ManageRouter{
		UserApi:     userService,
		TagApi:      tagService,
		TemplateApi: templateService,
	}
	app, err := pkg.NewApp(configConfig, monitorRouter, manageRouter)
	if err != nil {
		return nil, err
	}
	return app, nil
}
