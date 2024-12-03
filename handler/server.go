package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ink-paint/ink/config"
	"go.uber.org/dig"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	logger     *zap.Logger
	Config     *config.Config
	HTTPServer *http.Server
	Router     *gin.Engine
}

type ServerParams struct {
	dig.In
	Config *config.Config
	Logger *zap.Logger
}

func NewServer(param ServerParams, lifecycle fx.Lifecycle) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	conf := param.Config

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port),
		Handler: router,
	}

	s := &Server{
		logger:     param.Logger,
		Config:     param.Config,
		HTTPServer: httpServer,
		Router:     router,
	}

	lifecycle.Append(fx.Hook{
		OnStop:  httpServer.Shutdown,
		OnStart: s.Run,
	})
	return s
}

func (s *Server) Run(ctx context.Context) error {
	if config.IsDev() {
		gin.SetMode(gin.DebugMode)
	}
	go func() {
		if err := s.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// print err info when httpServer start failed
			s.logger.Error("unexpected error from ListenAndServe", zap.Error(err))
			fmt.Printf("http server start error:%s\n", err.Error())
			os.Exit(1)
		}
	}()
	return nil
}
