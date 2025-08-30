package dbrepository

import (
	models "inventory_management/Models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetByID(id uint) (*models.Order, error)
	GetAll() ([]models.Order, error)
	Update(order *models.Order) error
	Delete(id uint) error
}

type Orderdb struct {
	db *gorm.DB
}

func OrderRepo(db *gorm.DB) *Orderdb {
	return &Orderdb{db: db}
}

func (r *Orderdb) Create(order *models.Order) error {

	return r.db.Create(order).Error
}

func (r *Orderdb) GetByID(id uint) (*models.Order, error) {

	var order models.Order
	err := r.db.Preload("Items").First(&order, id).Error
	return &order, err
}

func (r *Orderdb) GetAll() ([]models.Order, error) {

	var orders []models.Order
	err := r.db.Preload("Items").Find(&orders).Error

	return orders, err
}

func (r *Orderdb) Update(order *models.Order) error {

	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Update the basic order fields
		if err := tx.Model(&models.Order{}).Where("id = ?", order.ID).
			Updates(map[string]interface{}{
				"customer_name": order.CustomerName,
				"status":        order.Status,
			}).Error; err != nil {
			return err
		}

		// 2. Delete existing items for the order
		if err := tx.Where("order_id = ?", order.ID).Delete(&models.OrderItem{}).Error; err != nil {
			return err
		}

		// 3. Set OrderID for each new item and insert
		for i := range order.Items {
			order.Items[i].OrderID = order.ID
		}

		if len(order.Items) > 0 {
			if err := tx.Create(&order.Items).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *Orderdb) Delete(id uint) error {

	return r.db.Delete(&models.Order{}, id).Error
}

/*

We use Preload("Items") when querying Order records to ensure that the related Items are fetched in the same operation.
This avoids the need for separate queries for each order's items,
improving efficiency and reducing DB load


*/
