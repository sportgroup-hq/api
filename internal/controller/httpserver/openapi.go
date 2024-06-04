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
	r.GET("/swagger-docs/openapi.yaml", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, docs.OpenAPI)
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger-docs/openapi.yaml")))
	r.GET("/swagger", func(context *gin.Context) {
		context.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})
}
