package handler

import (
	"project-adhyaksa/services/event/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.RouterGroup, usecase *usecase.UseCase) {
	eventHandler := NewEventHandler(usecase.EventUseCase)
	routes := r.Group("/event")
	{
		routes.POST("/", eventHandler.RegisterEvent)
	}

}
