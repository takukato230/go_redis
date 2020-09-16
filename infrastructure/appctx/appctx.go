package appctx

import (
	"github.com/labstack/echo"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/datasorce"
)

type APPContext struct {
	echo.Context
	Redis datasorce.RedisDriver
}

func NewAPPContext(ctx echo.Context, driver datasorce.RedisDriver) *APPContext {
	return &APPContext{
		Context: ctx,
		Redis:   driver,
	}
}
