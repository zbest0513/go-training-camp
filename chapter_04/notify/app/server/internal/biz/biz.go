package biz

import "github.com/google/wire"

// ProviderSetBiz is biz providers.
var ProviderSetBiz = wire.NewSet(NewUserUseCase, NewTagUseCase, NewTemplateUseCase)
