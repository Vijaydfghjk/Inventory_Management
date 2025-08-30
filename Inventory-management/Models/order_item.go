package models

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"` // Foreign key
	Product   string  `json:"product"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}
