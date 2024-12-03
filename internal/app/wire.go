package app

import (
	"github.com/google/wire"
	"love_knot/internal/app/storage/cache"
	"love_knot/internal/app/storage/repo"
	"love_knot/internal/job"
)

var SQLProviderSet = wire.NewSet(
	wire.Struct(new(job.SQLProvider), "*"),
)

var CacheProviderSet = wire.NewSet(
	cache.NewEmailStorage,
	cache.NewTokenSessionStorage,
)

var RepoProviderSet = wire.NewSet(
	repo.NewUsers,
	repo.NewSource,
	repo.NewDevice,
	repo.NewFriend,
)
