package graph

import (
	"github.com/IsJordanBraz/go-clean-architecture/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUsecase usecase.CreateOrderUseCase
	ListOrderUsecase   usecase.ListOrderUseCase
}
