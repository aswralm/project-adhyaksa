package handler

import (
	"project-adhyaksa/services/event/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.RouterGroup, usecase *usecase.UseCase) {
	eventHandler := NewEventHandler(usecase.EventUseCase)
	documentationHandler := NewDocumentationHandler(usecase.DocumentationUseCase)

	routesEvent := r.Group("/event")
	{
		routesEvent.POST("", eventHandler.RegisterEvent)
		routesEvent.GET("", eventHandler.GetListEventPaginated)

	}

	routesDocumentation := r.Group("/documentation")
	{
		routesDocumentation.POST("", documentationHandler.RegisterDocumentation)
	}

}
