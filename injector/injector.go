package injector

import (
	"context"
	"github.com/takutakukatokatojapan/go_redis/api/server"
	"github.com/takutakukatokatojapan/go_redis/domain/repository"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/config"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/datasource"
	"go.uber.org/dig"
)

var c *dig.Container

func init() {
	c = dig.New()
	c.Provide(config.NewConfig)
	c.Provide(datasource.NewRedisDriverImpl)
	c.Provide(server.NewServer)
	c.Provide(datasource.NewDBImpl)
	c.Provide(repository.NewUserRepository)
}

func Run() error {
	if err := c.Invoke(func(s server.Server) error {
		return s.RunServer(context.Background())
	}); err != nil {
		return err
	}
	return nil
}
