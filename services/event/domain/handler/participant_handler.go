package handler

import "github.com/gin-gonic/gin"

type ParticipantHandler interface {
	ConfirmEvent(ctx *gin.Context)
}
