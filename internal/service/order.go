package service

import (
	"context"
	"dockerized-api/internal/repository"
)

type (
	OrderService interface {
		GetOrders(ctx context.Context)
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

func (s *orderService) GetOrders(ctx context.Context) {
	s.repository.GetOrders(ctx)
}
