//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"love_knot/internal/app"
	"love_knot/internal/app/api"
	"love_knot/internal/app/config"
	"love_knot/internal/app/db"
	"love_knot/internal/provider"
)

var providerSet = wire.NewSet(
	provider.NewHttpClient,
	provider.NewRequestClient,
	provider.NewMysqlClient,
)

func NewHttpInjector(conf *config.Config) *api.AppProvider {
	panic(
		wire.Build(
			// providerSet,
			api.ProviderSet,
		),
	)
}

func NewSQLInjector(conf *config.Config) *db.SQLProvider {
	panic(
		wire.Build(
			providerSet,
			app.SQLProviderSet,
		),
	)
}
