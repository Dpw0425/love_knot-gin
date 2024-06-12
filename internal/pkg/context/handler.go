package ctx

import (
	"github.com/gin-gonic/gin"
	"love_knot/pkg/response"
)

func HandlerFunc(fn func(ctx *Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(New(c)); err != nil {
			response.ErrResponse(c, err)
			return
		}
	}
}
