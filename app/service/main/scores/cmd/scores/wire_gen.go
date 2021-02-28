// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mittacy/himusic/service/scores/internal/biz"
	"github.com/mittacy/himusic/service/scores/internal/conf"
	"github.com/mittacy/himusic/service/scores/internal/data"
	"github.com/mittacy/himusic/service/scores/internal/server"
	"github.com/mittacy/himusic/service/scores/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, error) {
	httpServer := server.NewHTTPServer(confServer)
	grpcServer := server.NewGRPCServer(confServer)
	dataData, err := data.NewData(confData, logger)
	if err != nil {
		return nil, err
	}
	scoresRepo := data.NewUserRepo(dataData)
	scoresUseCase := biz.NewScoresUseRepo(scoresRepo)
	scoresService := service.NewScoresService(scoresUseCase)
	app := newApp(logger, httpServer, grpcServer, scoresService)
	return app, nil
}