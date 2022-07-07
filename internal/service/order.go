package service

import "context"

type (
	OrderService interface {
		GetOrders(ctx context.Context)
	}

	orderService struct{}
)

func NewOrderService() OrderService {
	return &orderService{}
}

func (s *orderService) GetOrders(_ context.Context) {
	panic("ORDERS SERVICE!")
}
