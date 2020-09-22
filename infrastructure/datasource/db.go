package datasource

import (
	"database/sql"
	"fmt"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/config"
)

type (
	DB interface {
		GetConn() (*sql.DB, error)
	}
	DBImpl struct {
		dbHost     string
		dbUser     string
		dbPort     int
		dbPassword string
		dbName     string
		dbSslMode  string
	}
)

func NewDBImpl(conf config.Config) DB {
	return &DBImpl{
		dbHost:     conf.DBHost,
		dbUser:     conf.DBUser,
		dbPort:     conf.DBPort,
		dbPassword: conf.DBPassword,
		dbName:     conf.DBName,
		dbSslMode:  conf.DBSslMode,
	}
}

func (d DBImpl) GetConn() (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.dbUser, d.dbPassword, d.dbHost, d.dbPort, d.dbName))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
