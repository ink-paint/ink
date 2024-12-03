package main

import (
	"context"

	"github.com/ink-paint/ink/config"
	"github.com/ink-paint/ink/dal"
	"github.com/ink-paint/ink/handler"
	"github.com/ink-paint/ink/injection"
	"github.com/ink-paint/ink/log"
	"go.uber.org/fx"
)

func main() {
	app := InitApp()

	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}

	<-app.Done()
}

func InitApp() *fx.App {
	options := injection.GetOptions()
	options = append(options,
		fx.Provide(
			log.NewLogger,
			log.NewGormLogger,
			dal.NewGormDB,
			config.NewConfig,
			handler.NewServer,
		),
		fx.Populate(&dal.DB),
		fx.Invoke(
			func(s *handler.Server) {
				s.RegisterRouters()
			},
		),
	)
	app := fx.New(
		options...,
	)
	return app
}
