package lib

import "github.com/google/wire"

var LibSet = wire.NewSet(NewJwtService, NewUserService)
