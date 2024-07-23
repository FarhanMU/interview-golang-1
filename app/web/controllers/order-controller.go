package controllers

import (
	"net/http"

	"zen-test/app/helpers"
	"zen-test/app/middleware"
	"zen-test/app/web"
	"zen-test/app/web/models"
	"zen-test/app/web/services"
)

type OrderController interface {
	FindAllOrder(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
}

type OrderControllerImpl struct {
	OrderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

// FindAll Order godoc
// @Summary FindAll Order from the store
// @Description FindAll Order from the store
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=models.OrderResponse}
// @Failure 401 {object} web.WebResponse
// @Router /orders [get]
// @Security BearerAuth
func (c *OrderControllerImpl) FindAllOrder(w http.ResponseWriter, r *http.Request) {

	data, err := c.OrderService.FindAllOrder(r.Context())
	helpers.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	helpers.WriteResponseBody(w, webResponse)
}

// Create Order godoc
// @Summary create Order for the store
// @Description create Order for the store
// @Tags Order
// @Accept json
// @Produce json
// @Param Order body models.OrderCreate true "Order create"
// @Success 200 {object} web.WebResponse{data=models.OrderResponse}
// @Failure 401 {object} web.WebResponse
// @Router /orders [post]
// @Security BearerAuth
func (c *OrderControllerImpl) CreateOrder(w http.ResponseWriter, r *http.Request) {
	createOrderRequest := models.OrderItemCreateUpdate{}
	helpers.ToRequestBody(r, &createOrderRequest)

	userId := middleware.GetUserID(r)

	OrderResponse, err := c.OrderService.CreateOrder(r.Context(), createOrderRequest, userId)
	helpers.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   OrderResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}
