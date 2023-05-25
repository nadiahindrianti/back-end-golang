package controllers

import (
	"back-end-golang/dtos"
	"back-end-golang/helpers"
	"back-end-golang/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentController interface {
	GetPaymentByID(c echo.Context) error
	CreatePayment(c echo.Context) error
	UpdatePayment(c echo.Context) error
	DeletePayment(c echo.Context) error
}

type paymentController struct {
	paymentUsecase usecases.PaymentUsecase
}

func NewPaymentController(paymentUsecase usecases.PaymentUsecase) PaymentController {
	return &paymentController{paymentUsecase}
}

func (c *paymentController) GetPaymentByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	payment, err := c.paymentUsecase.GetPaymentByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get payment by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get payment by id",
			payment,
		),
	)

}

func (c *paymentController) CreatePayment(ctx echo.Context) error {
	var paymentDTO dtos.PaymentInput
	if err := ctx.Bind(&paymentDTO); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding payment",
				helpers.GetErrorData(err),
			),
		)
	}

	payment, err := c.paymentUsecase.CreatePayment(&paymentDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a payment",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a payment",
			payment,
		),
	)
}

func (c *paymentController) UpdatePayment(ctx echo.Context) error {
	var paymentInput dtos.PaymentInput
	if err := ctx.Bind(&paymentInput); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding payment",
				helpers.GetErrorData(err),
			),
		)
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	payment, err := c.paymentUsecase.GetPaymentByID(uint(id))
	if payment.PaymentID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get payment by id",
				helpers.GetErrorData(err),
			),
		)
	}

	paymentResp, err := c.paymentUsecase.UpdatePayment(uint(id), paymentInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding payment",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated payment",
			paymentResp,
		),
	)
}

func (c *paymentController) DeletePayment(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.paymentUsecase.DeletePayment(uint(id))
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed binding payment",
				helpers.GetErrorData(err),
			),
		)
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted payment",
			nil,
		),
	)
}
