package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/kim/go-crud/helper"
)

const (
	host	= "localhost"
	port 	= 5432
	user	= "postgres"
	password= "12345678"
	dbName = "gocrud"
)

func DatabaseConnection() * gorm.DB {
	
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}