package v1

import (
	ctx "love_knot/internal/pkg/context"
	"love_knot/schema/genproto/web/v1/common"
)

type Common struct {
}

func (c *Common) SendEmailCode(ctx *ctx.Context) error {
	params := &common.web{}

}
