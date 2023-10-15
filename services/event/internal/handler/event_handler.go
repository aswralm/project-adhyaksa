package handler

import (
	"net/http"
	"project-adhyaksa/pkg/pagination"
	recovery "project-adhyaksa/pkg/recover"
	"project-adhyaksa/pkg/res"
	"project-adhyaksa/services/event/domain/handler"
	"project-adhyaksa/services/event/domain/usecase"
	"project-adhyaksa/services/event/internal/customerror"
	"project-adhyaksa/services/event/internal/handler/mapping"
	"project-adhyaksa/services/event/internal/handler/request"
	"project-adhyaksa/services/event/internal/handler/response"

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
	eventUsecaseDTO := mapping.EventRequestToUsecaseDTO(&eventRequest)
	eventUsecaseDTO.AdminID = adminID

	//usecase event
	if err := c.eventUseCase.Create(eventUsecaseDTO, ctx); err != nil {
		if eventErr, ok := err.(*customerror.Err); ok {
			ctx.JSON(http.StatusBadRequest, res.JSON(false, "Failed to register event", eventErr))
			return
		}
		ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", nil))
		return
	}

	ctx.JSON(http.StatusCreated, res.JSON(true, "Register event successfully", nil))
}

func (c *eventHandler) GetListEventPaginated(ctx *gin.Context) {
	defer recovery.Recover(ctx)

	var filters request.EventQueryPaginated
	if err := ctx.Bind(&filters); err != nil {
		ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", err.Error()))
		return
	}

	// Create pagination instance
	var paging pagination.Paginator
	paging.Create(filters.Limit, filters.Page, ctx.Request.URL.Path)

	queryFilter := filters.QueryParamMapping()

	events, err := c.eventUseCase.GetListPaginated(&paging, &queryFilter)
	if err != nil {
		if customErr, ok := err.(*customerror.Err); ok {
			ctx.JSON(http.StatusBadRequest, res.JSON(false, "Failed to get events", customErr))
			return
		}
		ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", nil))
		return
	}

	result := response.ListMapping(&events)
	ctx.JSON(http.StatusOK, res.JSON(
		true,
		"Success get events",
		result,
		paging.Cursor(),
		map[string]any{
			"count": paging.Count,
		},
	))

}
