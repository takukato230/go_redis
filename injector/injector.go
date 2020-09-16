package injector

import (
	"context"
	"go.uber.org/dig"
	"github.com/takutakukatokatojapan/go_redis/api/server"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/config"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/datasorce"
)

var c *dig.Container

func init() {
	c = dig.New()
	c.Provide(config.NewConfig)
	c.Provide(datasorce.NewRedisDriverImpl)
	c.Provide(server.NewServer)
}

func Run() error {
	if err := c.Invoke(func(s server.Server) error {
		return s.RunServer(context.Background())
	}); err != nil {
		return err
	}
	return nil
}
