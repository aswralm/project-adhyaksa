package handler

import (
	"net/http"
	recovery "project-adhyaksa/pkg/recover"
	"project-adhyaksa/pkg/res"
	"project-adhyaksa/services/event/domain/handler"
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/customerror"
	"project-adhyaksa/services/event/internal/handler/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type documentationHandler struct {
	documentationUseCase usecase.DocumentatitonUseCase
}

func NewDocumentationHandler(documentationUseCase usecase.DocumentatitonUseCase) handler.DocumentationHandler {
	return &documentationHandler{documentationUseCase: documentationUseCase}
}

func (h *documentationHandler) RegisterDocumentation(ctx *gin.Context) {
	defer recovery.Recover(ctx)
	var (
		documentationRequest request.RegisterDocumentationRequest
		errorMessages        = make(map[string]string)
	)

	// Validate request
	if err := ctx.ShouldBind(&documentationRequest); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", nil))
			return
		}

		for _, e := range validationErr {
			fieldJSONName := documentationRequest.GetJsonFieldName(e.Field())
			errorMessages[fieldJSONName] = documentationRequest.ErrMessages()[fieldJSONName][e.ActualTag()]
		}
	}
	if len(errorMessages) > 0 {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			res.JSON(false, "Failed to register documentation", &customerror.Err{
				Code:   customerror.ERROR_INVALID_REQUEST,
				Errors: errorMessages,
			}),
		)
		return
	}

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		zap.L().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", nil))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		zap.L().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", nil))
		return
	}

	defer file.Close()

	//get data user_id from token
	adminID := ctx.GetString("user_id")

	//mapping from request to usecaseDTO
	documentation := usecase.DocumentationUseCaseDTO{
		BranchID:    documentationRequest.BranchID,
		AdminID:     adminID,
		Name:        documentationRequest.Name,
		PhotoName:   fileHeader.Filename,
		Date:        documentationRequest.Date,
		Location:    documentationRequest.Location,
		Description: documentationRequest.Description,
		Participant: documentationRequest.Participant,
	}

	//register event
	if err := h.documentationUseCase.Create(documentation, file, ctx); err != nil {
		if eventErr, ok := err.(*customerror.Err); ok {
			ctx.JSON(http.StatusBadRequest, res.JSON(false, "Failed to register event", eventErr))
			return
		}
		ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", nil))
		return
	}

	ctx.JSON(http.StatusCreated, res.JSON(true, "Register event successfully", nil))
}
