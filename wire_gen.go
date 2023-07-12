// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/app"
	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/controller"
	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/middleware"
	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/repository"
	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	categoryRepositoryImpl := repository.NewCategoryRepositoryImpl()
	db := app.InitDB()
	validate := validator.New()
	categoryServiceImpl := service.NewCategoryServiceImpl(categoryRepositoryImpl, db, validate)
	categoryControllerImpl := controller.NewCategoryControllerImpl(categoryServiceImpl)
	router := app.Router(categoryControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var categorySet = wire.NewSet(repository.NewCategoryRepositoryImpl, wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)), service.NewCategoryServiceImpl, wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)), controller.NewCategoryControllerImpl, wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)))