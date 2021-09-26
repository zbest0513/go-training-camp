//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"notify-server/internal/biz"
	"notify-server/internal/data"
	"notify-server/internal/service/handle"
	"notify-server/internal/service/router"
	"notify/pkg/config"
)

func initApp(fc config.FileConfig) (*router.Router, error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, handle.ProviderSet, router.RouterSet, data.SqlClientProviderSet))
}
