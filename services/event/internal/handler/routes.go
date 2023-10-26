package handler

import (
	"project-adhyaksa/services/event/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.RouterGroup, usecase *usecase.UseCase) {
	eventHandler := NewEventHandler(usecase.EventUseCase)
	documentationHandler := NewDocumentationHandler(usecase.DocumentationUseCase)
	participantHandler := NewParticipantHandler(usecase.ParticipantUseCase)

	routesEvent := r.Group("/event")
	{
		routesEvent.POST("", eventHandler.RegisterEvent)
		routesEvent.GET("", eventHandler.GetListEventPaginated)
		routesEvent.GET("/:id", eventHandler.GetEventByID)
		routesEvent.POST("/confirm", participantHandler.ConfirmEvent)

	}

	routesDocumentation := r.Group("/documentation")
	{
		routesDocumentation.POST("", documentationHandler.RegisterDocumentation)
		routesDocumentation.GET("", documentationHandler.GetListDocumentationPaginated)
		routesDocumentation.GET("/:id", documentationHandler.GetEventByID)
	}

}
