package service

import (
	models "inventory_management/Models"
	"inventory_management/dbrepository"
)

type OrderService interface {
	CreateOrder(order *models.Order) error
	GetOrder(id uint) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
}

type orderservice struct {
	repo dbrepository.OrderRepository
}

func NewOrderService(repo dbrepository.OrderRepository) *orderservice {

	return &orderservice{repo: repo}
}

func (r *orderservice) CreateOrder(order *models.Order) error {

	return r.repo.Create(order)
}

func (r *orderservice) GetOrder(id uint) (*models.Order, error) {

	return r.repo.GetByID(id)
}

func (r *orderservice) GetAllOrders() ([]models.Order, error) {

	return r.repo.GetAll()
}

func (r *orderservice) UpdateOrder(order *models.Order) error {
	return r.repo.Update(order)
}

func (r *orderservice) DeleteOrder(id uint) error {
	return r.repo.Delete(id)
}
