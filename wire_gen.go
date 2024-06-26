// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Laeeqdev/AttendanceMangements/API/Auth"
	"github.com/Laeeqdev/AttendanceMangements/API/Repository"
	"github.com/Laeeqdev/AttendanceMangements/API/RestHandler"
	"github.com/Laeeqdev/AttendanceMangements/API/Router"
	"github.com/Laeeqdev/AttendanceMangements/API/Service"
	"github.com/go-pg/pg"
)

// Injectors from wire.go:

func InitializeApp(db *pg.DB) *router.RouterImpl {
	userRepositoryImpl := repository.NewUserRepositoryImpl(db)
	userServiceImpl := service.NewUserServiceImpl(userRepositoryImpl)
	userAuthHandlerImpl := auth.NewUserAuthHandlerImpl(userServiceImpl)
	detailsRepositoryImpl := repository.NewDetailsRepositoryImpl(db)
	getDeatilsServiceImpl := service.NewGetDeatilsServiceImpl(detailsRepositoryImpl, userRepositoryImpl)
	detailsHandlerImpl := resthandler.NewDetailsHandlerImpl(getDeatilsServiceImpl, userAuthHandlerImpl)
	principalHandlerImpl := resthandler.NewPrincipalHandlerImpl(userServiceImpl, userAuthHandlerImpl)
	punchinRepositoryImpl := repository.NewPunchinRepositoryImpl(db)
	punchinServiceImpl := service.NewPunchinServiceImpl(punchinRepositoryImpl)
	punchInPunchOutHandlerImpl := resthandler.NewPunchInPunchOutHandlerImpl(punchinServiceImpl, userAuthHandlerImpl)
	routerImpl := router.NewRouterImpl(userAuthHandlerImpl, detailsHandlerImpl, principalHandlerImpl, punchInPunchOutHandlerImpl)
	return routerImpl
}
