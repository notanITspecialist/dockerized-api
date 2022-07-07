package endpoints

import (
	"dockerized-api/internal/service"
	"encoding/json"
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

func (h *orderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.GetOrders(r.Context())
	if err != nil {
		// log error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		// log error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
