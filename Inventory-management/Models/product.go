package models

import "time"

type Product struct {
	ID                uint               `json:"id" gorm:"primaryKey"`
	ModelID           string             `json:"model_id" binding:"required"`
	Name              string             `json:"name" binding:"required"`
	Category          string             `json:"category" binding:"required"`
	Quantity          int                `json:"quantity" binding:"required"`
	Price             float64            `json:"price" binding:"required"`
	WarehouseID       uint               `json:"warehouse_id" binding:"required"`
	WarehouseLocation string             `json:"warehouse_location" binding:"required"`                          // New Field: Physical location of warehouse                                      // // "In Stock" or "In Use"
	Assets            []ProductItemInput `json:"assets" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"` // Assigned user (nullable)
	UpdatedAt         time.Time          `json:"updated_at"`
}

type ProductItemInput struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	ProductID    uint   `json:"product_id"`
	SerialNumber string `gorm:"index" json:"serialnumber" binding:"required"`
}

/*

meansing of constraint:OnDelete:CASCADE"`

 When a parent Product is deleted, all its related ProductItems are automatically deleted from the database.


*/
