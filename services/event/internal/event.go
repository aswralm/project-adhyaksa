package event

import (
	"net/http"
	routes "project-adhyaksa/services/event/internal/handler"
	"project-adhyaksa/services/event/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewEvent(handler *gin.Engine, uc *usecase.UseCase) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	// swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	// handler.GET("/swagger/*any", swaggerHandler)

	// check health server
	handler.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, "server is online") })

	// Routers
	h := handler.Group("/api/v1")
	{
		routes.NewRoutes(h, uc)
	}
}
