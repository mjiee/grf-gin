package api

import "github.com/google/wire"

// ApiSet 为api providers
var ApiSet = wire.NewSet(NewAuthHandler, NewUserHandler)
