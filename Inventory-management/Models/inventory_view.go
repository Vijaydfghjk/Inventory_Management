package models

type Inventory struct {
	ModelID           string `json:"model_id"`
	SerialNumber      string `json:"serialnumber"`
	Name              string `json:"name" `
	Category          string `json:"category" `
	WarehouseID       uint   `json:"warehouse_id"`
	WarehouseLocation string `json:"warehouse_location"`
	Status            string `json:"status"`
	Borrower_by       string `json:"borrower_by"`
}
