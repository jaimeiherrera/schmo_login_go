package api

import (
	"github.com/jaimeiherrera/schmo_login_go/pkg/db"
	"github.com/jaimeiherrera/schmo_login_go/src/adapter"
	"github.com/jaimeiherrera/schmo_login_go/src/gateway"
)

type Components struct {
	UserRepository gateway.UserRepository
}

func NewComponents() *Components {
	db := db.NewLocalDB()
	return &Components{
		UserRepository: getUserRepository(db),
	}
}

func getUserRepository(db db.Database) gateway.UserRepository {
	return adapter.NewUserRepository(db)
}
