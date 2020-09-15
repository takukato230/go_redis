package server

import (
	"context"
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go_redis/api/types"
	"go_redis/infrastructure/config"
	"go_redis/infrastructure/datasorce"
	"net/http"
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
	e.POST("/set", func(c echo.Context) error {
		var req types.SetRequest
		if err := c.Bind(&req); err != nil {
			return c.String(http.StatusBadRequest, "json can not bind")
		}
		if err := s.redis.Set(req.ID, req.Name); err != nil {
			return c.JSON(http.StatusInternalServerError, types.SetResponse{Code: "500"})
		}
		return c.JSON(http.StatusOK, types.SetResponse{Code: "200"})
	})
	e.POST("/get", func(c echo.Context) error {
		return nil
	})
	if err := e.Start(s.config.EchoURL); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return errors.New("context canceled")
	}
}
