package model

import (
	"github.com/takutakukatokatojapan/go_redis/infrastructure/datasource/mysql_entities"
)

type User struct {
	mysql_entities.User
}

func NewUser(id int, name, email string) User {
	return User{
		mysql_entities.User{
			ID:    id,
			Name:  name,
			Email: email,
		},
	}
}
