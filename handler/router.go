package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ink-paint/ink/config"
)

func (s *Server) RegisterRouters() {
	router := s.Router
	if config.IsDev() {
		router.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowOrigins:     []string{},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Admin-Authorization", "Content-Type"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Content-Length"},
		}))
	}

	{
		router.GET("/ping", func(ctx *gin.Context) {
			_, _ = ctx.Writer.Write([]byte("pong"))
		})

	}
}
