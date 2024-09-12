package main

import (
	"net/http"
	"github.com/kim/go-crud/config"
	"github.com/kim/go-crud/controller"
	"github.com/kim/go-crud/helper"
	"github.com/kim/go-crud/model"
	"github.com/kim/go-crud/repository"
	"github.com/kim/go-crud/router"
	"github.com/kim/go-crud/service"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})
	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository, validate)
	userController := controller.NewUserController(userService)
	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
