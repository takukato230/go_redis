package server

import (
	"context"
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/takutakukatokatojapan/go_redis/api/server/handler"
	custom_middleware "github.com/takutakukatokatojapan/go_redis/api/server/middleware"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/config"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/datasorce"
)

type Server struct {
	config config.Config
	redis  datasorce.RedisDriver
}

func NewServer(config2 config.Config, driver datasorce.RedisDriver) Server {
	return Server{
		config: config2,
		redis:  driver,
	}
}

func (s Server) RunServer(ctx context.Context) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(custom_middleware.SetAPPCtxMiddleware(s.redis))
	e.POST("/set", handler.SetHandler)
	e.POST("/get", handler.GetHandler)
	if err := e.Start(s.config.EchoURL); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return errors.New("context canceled")
	}
}
