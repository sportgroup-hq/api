package httpserver

import (
	"log/slog"

	"github.com/gin-contrib/gzip"
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

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	addOpenAPIDocsRouter(r)

	api := r.Group("/api/v1")

	api.GET("/ping", pingHandler)

	authorized := api.Use(s.authMiddleware)

	// Users
	authorized.GET("/users/me", s.getMeHandler)

	// Groups
	authorized.GET("/groups", s.getGroupsHandler)
	authorized.POST("/groups", s.createGroupHandler)
	authorized.POST("/groups/join", s.joinGroupHandler)
	authorized.DELETE("/groups/:group_id", s.deleteGroupHandler)
	authorized.PUT("/groups/:group_id", s.updateGroupHandler)
	authorized.POST("/groups/:group_id/leave", s.leaveGroupHandler)

	slog.Info("Starting HTTP server on " + s.cfg.HTTP.Address + "...")

	return r.Run(s.cfg.HTTP.Address)
}
