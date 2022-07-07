package repository

import "context"

type (
	OrderRepository interface {
		GetOrders(ctx context.Context)
	}

	orderRepository struct{}
)

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) GetOrders(_ context.Context) {
	panic("ORDERS REPOSITORY!!!")
}
