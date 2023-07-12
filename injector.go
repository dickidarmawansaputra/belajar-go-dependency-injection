//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/app"
	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/controller"
	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/middleware"
	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/repository"
	"github.com/dickidarmawansaputra/belajar-go-dependency-injection/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryServiceImpl,
	// kalo ada yg butuh category service kita kirimkan categoryimpl
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryControllerImpl,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	// jika struct tidak perlu binding interface
	wire.Build(
		app.InitDB,
		validator.New,
		categorySet,
		app.Router,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
