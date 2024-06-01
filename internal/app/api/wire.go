package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"love_knot/internal/app/api/handler"
	"love_knot/internal/app/api/handler/web"
	"love_knot/internal/app/api/router"
	"love_knot/internal/config"
)

type AppProvider struct {
	Config *config.Config
	Engine *gin.Engine
}

var ProviderSet = wire.NewSet(
	router.NewRouter,

	handler.ProviderSet,
	web.ProviderSet,

	wire.Struct(new(AppProvider), "*"),
)
