package endpoints

import (
	"dockerized-api/internal/service"
	"net/http"
)

type (
	OrderEndpoints interface {
		GetOrders(http.ResponseWriter, *http.Request)
	}

	orderHandler struct {
		service service.OrderService
	}
)

func NewOrderEndpoints(
	service service.OrderService,
) OrderEndpoints {
	return &orderHandler{
		service: service,
	}
}

func (h *orderHandler) GetOrders(_ http.ResponseWriter, r *http.Request) {
	h.service.GetOrders(r.Context())
}
