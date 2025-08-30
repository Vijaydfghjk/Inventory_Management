package models

import "time"

type Borrower struct {
	Useby             string    `json:"useby"`
	Using_location    string    `json:"using_location"`
	ModelID           string    `json:"model_id"`
	Name              string    `json:"name"`
	Category          string    `json:"category"`
	WarehouseID       uint      `json:"warehouse_id"`
	WarehouseLocation string    `json:"warehouse_location"`
	Status            string    `json:"status"`
	Serial_number     string    `gorm:"index" json:"serial_number" binding:"required"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Borrower_Request struct {
	Use_by         string      `json:"use_by" binding:"required"`
	Using_location string      `json:"using_location" binding:"required"`
	View           []*Borrower `json:"view"`
}
