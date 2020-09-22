package appctx

import (
	"github.com/labstack/echo"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/datasource"
)

type APPContext struct {
	echo.Context
	Redis datasource.RedisDriver
}

func NewAPPContext(ctx echo.Context, driver datasource.RedisDriver) *APPContext {
	return &APPContext{
		Context: ctx,
		Redis:   driver,
	}
}
