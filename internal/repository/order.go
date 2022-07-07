package repository

import "context"

type (
	OrderRepository interface {
		GetOrders(ctx context.Context) ([]Order, error)
	}

	orderRepository struct{}
)

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) GetOrders(_ context.Context) ([]Order, error) {
	return []Order{
		{ID: "1", Title: "Title 1"},
		{ID: "2", Title: "Title 3"},
	}, nil
}
