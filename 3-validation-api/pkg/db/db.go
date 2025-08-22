// Package db
package db

import (
	"go-adv/http/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(conf *configs.Config) *DB {
	db, err := gorm.Open(postgres.Open(conf.DB.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DB{db}
}
