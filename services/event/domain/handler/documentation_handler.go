package handler

import "github.com/gin-gonic/gin"

type DocumentationHandler interface {
	RegisterDocumentation(ctx *gin.Context)
	GetListDocumentationPaginated(ctx *gin.Context)
}
