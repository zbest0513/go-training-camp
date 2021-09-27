package handle

import (
	"github.com/google/wire"
)

// ProviderSetHandle is service providers.
var ProviderSetHandle = wire.NewSet(NewUserService)
