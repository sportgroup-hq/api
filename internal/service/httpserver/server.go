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

	r.GET("/", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html")
		ctx.String(200, `<html><head></head><body><a href="http://localhost:8081/oauth2callback">login</a></body></html>`)
	})

	authorized := r.Use(s.authMiddleware)

	authorized.GET("/me", func(ctx *gin.Context) {
		value, exists := ctx.Get("userID")
		if !exists {
			ctx.AbortWithStatus(500)
			return
		}

		ctx.String(200, "Hi, "+value.(uuid.UUID).String())
	})

	slog.Info("Starting HTTP server on " + s.cfg.HTTP.Address + "...")

	return r.Run()
}
