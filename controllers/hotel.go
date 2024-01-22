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

type HotelController interface {
	//admin
	GetAllHotels(c echo.Context) error
	GetHotelByID(c echo.Context) error
	CreateHotel(c echo.Context) error
	UpdateHotel(c echo.Context) error
	DeleteHotel(c echo.Context) error
	SearchHotelAvailable(c echo.Context) error
}

type hotelController struct {
	hotelUsecase usecases.HotelUsecase
}

func NewHotelController(hotelUsecase usecases.HotelUsecase) HotelController {
	return &hotelController{hotelUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *hotelController) GetAllHotels(ctx echo.Context) error {
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

	minimumPrice, _ := strconv.Atoi(ctx.QueryParam("minimum_price"))
	maximumPrice, _ := strconv.Atoi(ctx.QueryParam("maximum_price"))
	ratingClass, _ := strconv.Atoi(ctx.QueryParam("rating_class"))

	addressParam := ctx.QueryParam("address")
	nameParam := ctx.QueryParam("name")

	sortByPriceParam := ctx.QueryParam("sort_by_price")

	recomendationParam := ctx.QueryParam("recomendation")
	recomendation, _ := strconv.ParseBool(recomendationParam)
	hotels, count, err := c.hotelUsecase.GetAllHotels(page, limit, minimumPrice, maximumPrice, ratingClass, addressParam, nameParam, sortByPriceParam, recomendation)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all hotel",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all hotels",
			hotels,
			page,
			limit,
			count,
		),
	)
}

func (c *hotelController) GetHotelByID(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				"Unauthorized",
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

	id, _ := strconv.Atoi(ctx.Param("id"))
	hotel, err := c.hotelUsecase.GetHotelByID(userId, uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusNotFound,
			helpers.NewErrorResponse(
				http.StatusNotFound,
				"Failed to get hotel by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get hotel by id",
			hotel,
		),
	)

}

func (c *hotelController) CreateHotel(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				"Unauthorized",
			),
		)
	}
	var hotelDTO dtos.HotelInput
	if err := ctx.Bind(&hotelDTO); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding hotel",
				helpers.GetErrorData(err),
			),
		)
	}

	hotel, err := c.hotelUsecase.CreateHotel(&hotelDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a hotel",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a hotel",
			hotel,
		),
	)
}

func (c *hotelController) UpdateHotel(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				"Unauthorized",
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
	var hotelInput dtos.HotelInput
	if err := ctx.Bind(&hotelInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding hotel",
				helpers.GetErrorData(err),
			),
		)
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	hotel, err := c.hotelUsecase.GetHotelByID(userId, uint(id))
	if hotel.HotelID == 0 {
		return ctx.JSON(
			http.StatusNotFound,
			helpers.NewErrorResponse(
				http.StatusNotFound,
				"Failed to get hotel by id",
				helpers.GetErrorData(err),
			),
		)
	}

	hotelResp, err := c.hotelUsecase.UpdateHotel(uint(id), hotelInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to updated a hotel",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated hotel",
			hotelResp,
		),
	)
}

func (c *hotelController) DeleteHotel(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				"Unauthorized",
			),
		)
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.hotelUsecase.DeleteHotel(uint(id))
	if err != nil {
		return ctx.JSON(
			http.StatusNotFound,
			helpers.NewErrorResponse(
				http.StatusNotFound,
				"Failed to delete hotel",
				helpers.GetErrorData(err),
			),
		)
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted hotel",
			nil,
		),
	)
}

// =============================== USER ================================== \\

func (c *hotelController) SearchHotelAvailable(ctx echo.Context) error {
	userId := uint(1)
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		userId = 1
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

	minimumPrice, _ := strconv.Atoi(ctx.QueryParam("minimum_price"))
	maximumPrice, _ := strconv.Atoi(ctx.QueryParam("maximum_price"))
	ratingClass, _ := strconv.Atoi(ctx.QueryParam("rating_class"))

	addressParam := ctx.QueryParam("address")
	nameParam := ctx.QueryParam("name")

	sortByPriceParam := ctx.QueryParam("sort_by_price")
	hotels, count, err := c.hotelUsecase.SearchHotelAvailable(int(userId), page, limit, minimumPrice, maximumPrice, ratingClass, addressParam, nameParam, sortByPriceParam)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all hotel",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all hotels",
			hotels,
			page,
			limit,
			count,
		),
	)
}
