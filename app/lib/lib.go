package lib

import "github.com/google/wire"

// LibSet 服务providers
var LibSet = wire.NewSet(NewJwtService, NewUserService)
