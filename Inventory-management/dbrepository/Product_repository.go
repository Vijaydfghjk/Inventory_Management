package dbrepository

import (
	"errors"
	models "inventory_management/Models"
	"log"
	"time"

	"gorm.io/gorm"
)

type Product_repository interface {
	Create(product *models.Product) error
	GetByID(id uint) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Update(product *models.Product) error
	Create_borrower(product *models.Borrower) (*models.Borrower, error)
	Delete(id uint) error
	Fetching_data_serialnumber(serialnumber string) (*models.Borrower, error)
	View_byname(username string) ([]models.Borrower, error)
	CheckSerialnumber(serialnumber string) (bool, error)
	Borrower_stock_status(serialNumber string) (bool, error)
	In_stock(serialNumber string) (bool, error)
	GetUsername(serialNumber string) (string, error)
}

type Product_db struct {
	db *gorm.DB
}

func Product_repo(db *gorm.DB) *Product_db {
	return &Product_db{db: db}
}

func (p *Product_db) Create(product *models.Product) error {

	return p.db.Create(product).Error
}

func (p *Product_db) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := p.db.Preload("Assets").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product_db) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Preload("Assets").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Product_db) Update(product *models.Product) error {

	return p.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&models.Product{}).Where("id=?", product.ID).Updates(

			map[string]interface{}{

				"model_id":           product.ModelID,
				"category":           product.Category,
				"quantity":           product.Quantity,
				"price":              product.Price,
				"warehouse_id":       product.WarehouseID,
				"warehouse_location": product.WarehouseLocation,
				"updated_at":         time.Now(),
			}).Error; err != nil {

			return err
		}

		log.Println("checking", product)
		if err := tx.Where("product_id = ?", product.ID).Delete(models.ProductItemInput{}).Error; err != nil {

			return err
		}

		for i := range product.Assets {

			product.Assets[i].ProductID = product.ID
		}

		log.Println("After", product)
		if len(product.Assets) > 0 {

			if err := tx.Create(&product.Assets).Error; err != nil {

				return err
			}
		}

		return nil
	})

}

func (p *Product_db) Delete(id uint) error {
	return p.db.Delete(&models.Product{}, id).Error
}

func (p *Product_db) Fetching_data_serialnumber(serialnumber string) (*models.Borrower, error) {

	dropquerry := "DROP VIEW IF EXISTS borrower_view;"

	if err := p.db.Exec(dropquerry).Error; err != nil {

		return &models.Borrower{}, err
	}

	createViewQuery := `
	
	       create view  borrower_view as

		   select 
		   products.model_id,
           product_item_inputs.serial_number,
           products.name,
           products.category,
		   products.warehouse_id,
           products.warehouse_location
           from products
           inner join product_item_inputs on products.id = product_item_inputs.product_id;`

	if err := p.db.Exec(createViewQuery).Error; err != nil {

		return &models.Borrower{}, err
	}

	var borrowers *models.Borrower

	select_querry := "select * from borrower_view where serial_number = ?;"

	if err := p.db.Raw(select_querry, serialnumber).Scan(&borrowers).Error; err != nil {

		return borrowers, err
	}

	return borrowers, nil
}

func (a *Product_db) Create_borrower(product *models.Borrower) (*models.Borrower, error) {

	if err := a.db.Create(product).Error; err != nil {

		return nil, err
	}

	return product, nil
}

func (a *Product_db) View_byname(username string) ([]models.Borrower, error) {

	var filerter_borrower []models.Borrower

	if err := a.db.Where("useby=?", username).Find(&filerter_borrower).Error; err != nil {

		return filerter_borrower, err
	}

	return filerter_borrower, nil
}

func (a *Product_db) CheckSerialnumber(serialnumber string) (bool, error) {

	var temp models.ProductItemInput

	if err := a.db.Where("serial_number=?", serialnumber).Find(&temp).Error; err != nil {

		return true, err
	}

	if temp.SerialNumber != "" {

		return true, nil
	}

	return false, nil
}

func (a *Product_db) Borrower_stock_status(serialNumber string) (bool, error) {

	var temp models.Borrower
	if err := a.db.Where("serial_number=?", serialNumber).Find(&temp).Error; err != nil {

		return true, err
	}

	if temp.Serial_number != "" {

		return true, nil
	}

	return false, nil
}

func (a *Product_db) In_stock(serialNumber string) (bool, error) {

	result := a.db.Where("serial_number=?", serialNumber).Delete(&models.Borrower{})

	if result.Error != nil {

		return false, result.Error
	}

	if result.RowsAffected == 0 {

		return false, nil
	}

	return true, nil
}

func (a *Product_db) GetUsername(serialNumber string) (string, error) {

	var temp models.Borrower
	err := a.db.Where("serial_number=?", serialNumber).Take(&temp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "", err
	}

	return temp.Useby, nil
}

/*

1. func (*gorm.DB).Find(dest interface{}, conds ...interface{})

Difference between  First Vs Find

 First(&user, id) - Retrieves only one record (LIMIT 1).
 Find(&users, []int{...}) -  Retrieves	All matching records


About Find Parameter

dest - Variable to store results (struct or slice, by pointer)
conds - Optional conditions — primary key, SQL WHERE clause, or list of IDs

example - db.Find(&users, "age > ?", 25) || db.Find(&user, 1)


2. Preload


  type product struct{

    Assets   []ProductItemInput
   }

 That means instead of doing a separate query to get the assets later,
 GORM will join or do a separate query immediately and populate the Assets field of the product model.


3. What is Transactions ?

 A transaction is a group of one or more database operations that are executed together.
It ensures that either all changes happen successfully, or none of them do at all.

Example
They help protect your data from becoming incomplete or corrupted when:

An operation fails

A crash happens

A business rule isn't met

About RollBack() - You start a transaction → make some changes → something goes wrong → you call Rollback() → all changes are canceled.

 ACID (Core Principles of Transactions):
Property	Simple Meaning
Atomicity	All or nothing — no partial updates
Consistency	Keeps data valid (e.g. no negative stock)
Isolation	Other users don’t see incomplete work
Durability	Once saved (committed), it's permanent

*/
