package endpoints

import "net/http"

type (
	OrderEndpoints interface {
		GetOrders(http.ResponseWriter, *http.Request)
	}

	orderHandler struct{}
)

func NewOrderEndpoints() OrderEndpoints {
	return &orderHandler{}
}

func (h *orderHandler) GetOrders(http.ResponseWriter, *http.Request) {
	panic("ORDERS!!")
}
