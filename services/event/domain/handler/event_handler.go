package handler

import "github.com/gin-gonic/gin"

type EventHandler interface {
	RegisterEvent(ctx *gin.Context)
	GetListEventPaginated(ctx *gin.Context)
	GetEventByID(ctx *gin.Context)
}
