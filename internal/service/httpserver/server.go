package httpserver

import (
	"log/slog"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/sportgroup-hq/api/internal/config"
)

func (s Server) Start() error {
	r := gin.Default()

	addOpenAPIDocsRouter(r)

	cfg := config.Get()

	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(200, `<html><head></head><body><a href="/oauth2callback">login</a></body></html>`)
	})

	slog.Info("Starting HTTP server on " + cfg.HTTP.Address + "...")

	return r.Run()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
