package httpserver

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sportgroup-hq/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func addOpenAPIDocsRouter(r *gin.Engine) {
	r.GET("/openapi.yaml", func(c *gin.Context) {
		c.String(http.StatusOK, docs.OpenAPI)
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/openapi.yaml")))
	r.GET("/swagger", func(context *gin.Context) {
		context.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})
}
