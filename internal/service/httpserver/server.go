package httpserver

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sportgroup-hq/common-lib/validation"
)

func (s *Server) Start() error {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validation.Register(v)
	}

	addOpenAPIDocsRouter(r)

	api := r.Group("/api/v1")

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "api pong")
	})

	authorized := api.Use(s.authMiddleware)

	authorized.GET("/me", func(ctx *gin.Context) {
		value, exists := ctx.Get("userID")
		if !exists {
			ctx.AbortWithStatus(500)
			return
		}

		ctx.String(200, "Hi, "+value.(uuid.UUID).String())
	})

	slog.Info("Starting HTTP server on " + s.cfg.HTTP.Address + "...")

	return r.Run(s.cfg.HTTP.Address)
}
