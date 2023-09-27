package handler

import "github.com/gin-gonic/gin"

type EventHandler interface {
	RegisterEvent(ctx *gin.Context)
}
