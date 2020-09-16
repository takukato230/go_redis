package middleware

import (
	"github.com/labstack/echo"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/appctx"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/datasorce"
)

func SetAPPCtxMiddleware(driver datasorce.RedisDriver) echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			c := appctx.NewAPPContext(context, driver)
			return handlerFunc(c)
		}
	}
}
