package service

import (
	"github.com/google/wire"
	api "notify/api/server/service"
)

// ProviderSetHandle is service providers.
var ProviderSetHandle = wire.NewSet(
	NewUserService,
	wire.Bind(new(api.UserApi), new(*UserService)),
	NewTagService,
	wire.Bind(new(api.TagApi), new(*TagService)),
	NewTemplateService,
	wire.Bind(new(api.TemplateApi), new(*TemplateService)),
)
