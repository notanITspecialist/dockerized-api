package service

import (
	"context"
	"dockerized-api/internal/repository"
)

type (
	OrderService interface {
		GetOrders(ctx context.Context) ([]repository.Order, error)
	}

	orderService struct {
		repository repository.OrderRepository
	}
)

func NewOrderService(
	repository repository.OrderRepository,
) OrderService {
	return &orderService{
		repository: repository,
	}
}

func (s *orderService) GetOrders(ctx context.Context) ([]repository.Order, error) {
	orders, err := s.repository.GetOrders(ctx)
	if err != nil {
		// log error
		return nil, err
	}

	return orders, err
}
