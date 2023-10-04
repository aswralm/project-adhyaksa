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

type eventHandler struct {
	eventUseCase usecase.EventUseCase
}

func NewEventHandler(eventUseCase usecase.EventUseCase) handler.EventHandler {
	return &eventHandler{eventUseCase: eventUseCase}
}

func (c *eventHandler) RegisterEvent(ctx *gin.Context) {
	defer recovery.Recover(ctx)
	var (
		eventRequest  request.RegisterEventRequest
		errorMessages = make(map[string]string)
	)
	//validate token

	// Validate request
	if err := ctx.ShouldBind(&eventRequest); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", nil))
			return
		}

		for _, e := range validationErr {
			fieldJSONName := eventRequest.GetJsonFieldName(e.Field())
			errorMessages[fieldJSONName] = eventRequest.ErrMessages()[fieldJSONName][e.ActualTag()]
		}
	}
	if len(errorMessages) > 0 {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			res.JSON(false, "Failed to register event", &customerror.Err{
				Code:   customerror.ERROR_INVALID_REQUEST,
				Errors: errorMessages,
			}),
		)
		return
	}

	//get data user_id from token
	adminID := ctx.GetString("user_id")

	//mapping from request to usecaseDTO
	event := usecase.EventUseCaseDTO{
		BranchID:    eventRequest.BranchID,
		AdminID:     adminID,
		Name:        eventRequest.Name,
		StartTime:   eventRequest.StartTime,
		EndTime:     eventRequest.EndTime,
		Location:    eventRequest.Location,
		Description: eventRequest.Description,
	}

	//usecase event
	if err := c.eventUseCase.Create(event, ctx); err != nil {
		if eventErr, ok := err.(*customerror.Err); ok {
			ctx.JSON(http.StatusBadRequest, res.JSON(false, "Failed to register event", eventErr))
			return
		}
		ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", nil))
		return
	}

	ctx.JSON(http.StatusCreated, res.JSON(true, "Register event successfully", nil))
}
