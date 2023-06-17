package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/middlewares"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HistorySearchStationController interface {
	GetAllHistorySearchStation(c echo.Context) error
	GetHistorySearchStationByID(c echo.Context) error
	CreateHistorySearchStation(c echo.Context) error
	UpdateHistorySearchStation(c echo.Context) error
}

type historySearchStationController struct {
	historySearchStationUsecase usecases.HistorySearchStationUseCase
}

func NewHistorySearchStationController(historySearchStationUsecase usecases.HistorySearchStationUseCase) HistorySearchStationController {
	return &historySearchStationController{historySearchStationUsecase}
}

func (c *historySearchStationController) GetAllHistorySearchStation(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(nil),
			),
		)
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	if err != nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(err),
			),
		)
	}

	pageParam := ctx.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}

	limitParam := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 1000
	}

	historySearchStation, count, err := c.historySearchStationUsecase.GetAllHistorySearchStation(userId, page, limit)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get a History Search Station",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully to get History Search Stations",
			historySearchStation,
			page,
			limit,
			count,
		),
	)
}

func (c *historySearchStationController) GetHistorySearchStationByID(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(nil),
			),
		)
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	if err != nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(err),
			),
		)
	}


	IDParam := ctx.QueryParam("ID")
	ID, _ := strconv.Atoi(IDParam)

	historySearchStation, err := c.historySearchStationUsecase.GetHistorySearchStationByID(userId, uint(ID))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get a History Search Station",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get History Search Station",
			historySearchStation,
		),
	)
}

func (c *historySearchStationController) CreateHistorySearchStation(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(nil),
			),
		)
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	if err != nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(err),
			),
		)
	}

	var historySearchStationInput dtos.HistorySearchStationInput
	if err := ctx.Bind(&historySearchStationInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding History Search Station",
				helpers.GetErrorData(err),
			),
		)
	}

	historySearchStation, err := c.historySearchStationUsecase.CreateHistorySearchStation(userId, historySearchStationInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a History Search Station",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a History Search Station",
			historySearchStation,
		),
	)
}

func (c *historySearchStationController) UpdateHistorySearchStation(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(nil),
			),
		)
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	if err != nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(err),
			),
		)
	}

	var historySearchStationInput dtos.HistorySearchStationInput
	if err := ctx.Bind(&historySearchStationInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding History Search Station",
				helpers.GetErrorData(err),
			),
		)
	}

	IDParam := ctx.QueryParam("id")
	ID, _ := strconv.Atoi(IDParam)

	historySearchStation, err := c.historySearchStationUsecase.UpdateHistorySearchStation(userId, uint(ID), historySearchStationInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to update a History Search Station",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to update a History Search Station",
			historySearchStation,
		),
	)

}
