package models

import "time"

type Order struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	CustomerName string      `json:"customer_name"`
	Status       string      `json:"status"` // e.g., "created", "shipped", "cancelled"
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	Items        []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}
