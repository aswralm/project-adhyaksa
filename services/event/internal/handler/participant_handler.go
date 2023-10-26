package handler

import (
	"net/http"
	recovery "project-adhyaksa/pkg/recover"
	"project-adhyaksa/pkg/res"
	"project-adhyaksa/services/event/domain/handler"
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/customerror"
	"project-adhyaksa/services/event/internal/handler/mapping"
	"project-adhyaksa/services/event/internal/handler/request"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type participantHandler struct {
	participantUseCase usecase.ParticipantUseCase
}

func NewParticipantHandler(participantUseCase usecase.ParticipantUseCase) handler.ParticipantHandler {
	return &participantHandler{participantUseCase: participantUseCase}
}

func (h *participantHandler) ConfirmEvent(ctx *gin.Context) {
	defer recovery.Recover(ctx)
	var (
		participantRequest request.ConfirmEventRequest
		errorMessages      = make(map[string]string)
	)
	//validate token

	// Validate request
	if err := ctx.ShouldBind(&participantRequest); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", nil))
			return
		}

		for _, e := range validationErr {
			fieldJSONName := participantRequest.GetJsonFieldName(e.Field())
			errorMessages[fieldJSONName] = participantRequest.ErrMessages()[fieldJSONName][e.ActualTag()]
		}
	}
	if len(errorMessages) > 0 {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			res.JSON(false, "Failed to confirm event", &customerror.Err{
				Code:   customerror.ERROR_INVALID_REQUEST,
				Errors: errorMessages,
			}),
		)
		return
	}

	//get data user_id from token
	userID := ctx.GetString("user_id")

	useCaseDTO := mapping.ParticipantRequestToUsecaseDTO(&participantRequest)
	useCaseDTO.UserID = userID

	if err := h.participantUseCase.ConfirmEvent(useCaseDTO, ctx); err != nil {
		if eventErr, ok := err.(*customerror.Err); ok {
			ctx.JSON(http.StatusBadRequest, res.JSON(false, "Failed to confirm event", eventErr))
			return
		}
		ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", nil))
		return
	}

	ctx.JSON(http.StatusCreated, res.JSON(true, "Confirm event successfully", nil))

}
