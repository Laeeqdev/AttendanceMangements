//go:build wireinject
// +build wireinject

package main

import (
	auth "github.com/Laeeqdev/AttendanceMangements/API/Auth"
	repository "github.com/Laeeqdev/AttendanceMangements/API/Repository"
	resthandler "github.com/Laeeqdev/AttendanceMangements/API/RestHandler"
	router "github.com/Laeeqdev/AttendanceMangements/API/Router"
	service "github.com/Laeeqdev/AttendanceMangements/API/Service"
	"github.com/go-pg/pg"
	"github.com/google/wire"
)

func InitializeApp(*pg.DB) *router.RouterImpl {

	wire.Build(auth.NewUserAuthHandlerImpl, wire.Bind(new(auth.UserAuthHandler), new(*auth.UserAuthHandlerImpl)),
		repository.NewDetailsRepositoryImpl, wire.Bind(new(repository.DetailsRepository), new(*repository.DetailsRepositoryImpl)),
		repository.NewPunchinRepositoryImpl, wire.Bind(new(repository.PunchinRepository), new(*repository.PunchinRepositoryImpl)),
		repository.NewUserRepositoryImpl, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
		resthandler.NewDetailsHandlerImpl, wire.Bind(new(resthandler.DetailsHandler), new(*resthandler.DetailsHandlerImpl)),
		resthandler.NewPrincipalHandlerImpl, wire.Bind(new(resthandler.PrincipalHandler), new(*resthandler.PrincipalHandlerImpl)),
		resthandler.NewPunchInPunchOutHandlerImpl, wire.Bind(new(resthandler.PunchInPunchOutHandler), new(*resthandler.PunchInPunchOutHandlerImpl)),
		router.NewRouterImpl,
		service.NewGetDeatilsServiceImpl, wire.Bind(new(service.GetDeatilsService), new(*service.GetDeatilsServiceImpl)),
		service.NewPunchinServiceImpl, wire.Bind(new(service.PunchinService), new(*service.PunchinServiceImpl)),
		service.NewUserServiceImpl, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))
	return &router.RouterImpl{}
}
