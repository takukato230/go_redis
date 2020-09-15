package appctx

import (
	"github.com/labstack/echo"
	"go_redis/infrastructure/datasorce"
)

type APPContext struct {
	echo.Echo
	redis datasorce.RedisDriver
}

