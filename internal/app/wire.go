package app

import (
	"github.com/google/wire"
	"love_knot/internal/app/db"
)

var SQLProviderSet = wire.NewSet(
	wire.Struct(new(db.SQLProvider), "*"),
)
