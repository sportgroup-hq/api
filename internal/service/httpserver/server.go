package httpserver

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sportgroup-hq/common-lib/validation"
)

func (s *Server) Start() error {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validation.Register(v)
	}

	addOpenAPIDocsRouter(r)

	api := r.Group("/api/v1")

	api.GET("/ping", pingHandler)

	authorized := api.Use(s.authMiddleware)

	authorized.GET("/users/me", s.getMeHandler)

	authorized.GET("/groups", s.getGroupsHandler)
	authorized.POST("/groups", s.createGroupsHandler)

	slog.Info("Starting HTTP server on " + s.cfg.HTTP.Address + "...")

	return r.Run(s.cfg.HTTP.Address)
}
