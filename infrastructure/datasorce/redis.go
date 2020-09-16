package datasorce

import (
	"github.com/gomodule/redigo/redis"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/config"
	"log"
)

type (
	RedisDriver interface {
		Set(key, value string) error
		Get(key string) (value string, err error)
	}
	RedisDriverImpl struct {
		config config.Config
	}
)

func NewRedisDriverImpl(config config.Config) RedisDriver {
	return &RedisDriverImpl{config: config}
}

func (r RedisDriverImpl) Set(key, value string) error {
	conn, errConn := r.getConnection()
	if errConn != nil {
		return errConn
	}
	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

func (r RedisDriverImpl) Get(key string) (value string, err error) {
	conn, errConn := r.getConnection()
	if errConn != nil {
		return "", errConn
	}
	c, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return c, nil
}

func (r RedisDriverImpl) getConnection() (redis.Conn, error) {
	log.Printf("redis url: %s\n", r.config.RedisURL)
	conn, err := redis.Dial("tcp", r.config.RedisURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
