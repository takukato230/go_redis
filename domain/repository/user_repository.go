package repository

import (
	"context"
	"github.com/takutakukatokatojapan/go_redis/domain/model"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/datasource"
	"github.com/volatiletech/sqlboiler/boil"
)

type (
	UserRepository interface {
		Save(ctx context.Context, user model.User) error
		UpdateEmail(ctx context.Context, user model.User) error
	}

	UserRepositoryImpl struct {
		datasource.DB
	}
)

func NewUserRepository(datasource datasource.DB) UserRepository  {
	return &UserRepositoryImpl{
		datasource,
	}
}

func (u UserRepositoryImpl) Save(ctx context.Context, user model.User) error {
	db, errConn := u.GetConn()
	defer db.Close()
	if errConn != nil {
		return errConn
	}
	err := user.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepositoryImpl) UpdateEmail(ctx context.Context, user model.User) error {
	db, errConn := u.GetConn()
	defer db.Close()
	if errConn != nil {
		return errConn
	}
	_, err := user.Update(ctx, db, boil.Whitelist("email"))
	if err != nil {
		return err
	}
	return nil
}

