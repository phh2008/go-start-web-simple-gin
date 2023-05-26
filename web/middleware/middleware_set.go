package middleware

import "github.com/google/wire"

var MiddleWareSet = wire.NewSet(AuthSet)
