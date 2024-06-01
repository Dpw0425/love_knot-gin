package handler

import (
	"github.com/google/wire"
	"love_knot/internal/app/api/handler/web"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(web.Handler), "*"),
	wire.Struct(new(Handler), "*"),
)
