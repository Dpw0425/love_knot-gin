package web

import (
	"github.com/google/wire"
	v1 "love_knot/internal/app/api/handler/web/v1"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(v1.Common), "*"),
	wire.Struct(new(v1.User), "*"),

	wire.Struct(new(V1), "*"),
)
