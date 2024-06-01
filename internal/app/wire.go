package app

import (
	"github.com/google/wire"
	"love_knot/internal/job"
)

var SQLProviderSet = wire.NewSet(
	wire.Struct(new(job.SQLProvider), "*"),
)
