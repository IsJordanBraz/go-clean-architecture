package entity

type OrderRepositoryInterface interface {
	Save(order *Order) (Order, error)
	FindAll() ([]Order, error)
}
