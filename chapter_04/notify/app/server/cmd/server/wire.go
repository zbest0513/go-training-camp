//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"notify-server/internal/biz"
	"notify-server/internal/data"
	"notify-server/internal/pkg"
	"notify-server/internal/service/handle"
	"notify-server/internal/service/router"
	"notify/pkg/config"
)

func initApp(fc config.FileConfig) (*pkg.App, error) {
	panic(wire.Build(pkg.ProviderSetPkg, data.ProviderSetData, biz.ProviderSetBiz, handle.
		ProviderSetHandle, router.ProviderSetRouter, data.SqlClientProviderSet))
}
