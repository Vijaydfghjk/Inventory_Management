package service

import (
	models "inventory_management/Models"
	"inventory_management/dbrepository"
)

type Product_service interface {
	Create_Product(product *models.Product) error
	Getproduct(id uint) (*models.Product, error)
	GetALLproduct() ([]models.Product, error)
	Updateproduct(product *models.Product) error
	Deleteproduct(id uint) error
	Inuse(serialnumber string) (*models.Borrower, error)
	Save_borrower(product *models.Borrower) (*models.Borrower, error)
	View_by_user(Username string) ([]models.Borrower, error)
	CheckSerial_number(serialnumber string) (bool, error)
	Borrower_stock(serialnumber string) (bool, error)
	Do_Instock(serialnumber string) (bool, error)
	Getby_username(serialnumber string) (string, error)
}

type Product_struct struct {
	product_repo dbrepository.Product_repository
}

func New_Product_service(val dbrepository.Product_repository) *Product_struct {

	return &Product_struct{product_repo: val}
}
func (a *Product_struct) Create_Product(product *models.Product) error {

	return a.product_repo.Create(product)
}

func (a *Product_struct) Getproduct(id uint) (*models.Product, error) {

	return a.product_repo.GetByID(id)
}

func (a *Product_struct) GetALLproduct() ([]models.Product, error) {

	return a.product_repo.GetAll()
}

func (a *Product_struct) Updateproduct(product *models.Product) error {

	return a.product_repo.Update(product)
}

func (a *Product_struct) Deleteproduct(id uint) error {

	return a.product_repo.Delete(id)
}

func (a *Product_struct) Inuse(serialnumber string) (*models.Borrower, error) {

	borrower, err := a.product_repo.Fetching_data_serialnumber(serialnumber)

	return borrower, err
}

func (a *Product_struct) Save_borrower(product *models.Borrower) (*models.Borrower, error) {

	return a.product_repo.Create_borrower(product)
}

func (a *Product_struct) View_by_user(Username string) ([]models.Borrower, error) {

	return a.product_repo.View_byname(Username)
}

func (a *Product_struct) CheckSerial_number(serialnumber string) (bool, error) {

	return a.product_repo.CheckSerialnumber(serialnumber)
}

func (a *Product_struct) Borrower_stock(serialnumber string) (bool, error) {

	return a.product_repo.Borrower_stock_status(serialnumber)
}

func (a *Product_struct) Do_Instock(serialnumber string) (bool, error) {

	return a.product_repo.In_stock(serialnumber)
}

func (a *Product_struct) Getby_username(serialnumber string) (string, error) {

	return a.product_repo.GetUsername(serialnumber)
}
