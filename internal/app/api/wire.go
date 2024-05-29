package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"love_knot/internal/app/api/router"
	"love_knot/internal/app/config"
)

type AppProvider struct {
	Config *config.Config
	Engine *gin.Engine
}

var ProviderSet = wire.NewSet(
	router.NewRouter,

	wire.Struct(new(AppProvider), "*"),
)
