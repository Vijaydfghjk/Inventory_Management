package controllers

import (
	"fmt"
	models "inventory_management/Models"
	service "inventory_management/Service"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Multiple_serialnumber struct {
	Serial_number string `json:"serial_number" binding:"required"`
}

type Product_controll struct {
	service  service.Product_service
	validate *validator.Validate
}

func Product_con(service service.Product_service) *Product_controll {

	return &Product_controll{service: service,
		validate: validator.New(),
	}
}

func (a *Product_controll) Add_product(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	if valerr := a.validate.Struct(product); valerr != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": valerr.Error()})
		return
	}

	if product.Quantity != len(product.Assets) {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "Quantity mismatch"})
		return
	}

	for _, v := range product.Assets {

		check, err := a.service.CheckSerial_number(v.SerialNumber)

		if check || err != nil {

			c.JSON(http.StatusConflict, gin.H{"Message": "Given Serial Number alredy exist"})
			return
		}
	}

	err := a.service.Create_Product(&product)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (a *Product_controll) View(c *gin.Context) {

	products, err := a.service.GetALLproduct()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (a *Product_controll) Viewby_id(c *gin.Context) {

	vid := c.Param("id")

	id, err := strconv.Atoi(vid)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid_Request"})
		return
	}

	product, err := a.service.Getproduct(uint(id))

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to fetch the data"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (a *Product_controll) Update_Product(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "invalid_request"})
		return
	}

	if valerr := a.validate.Struct(product); valerr != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": valerr.Error()})
		return
	}

	if product.Quantity != len(product.Assets) {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "Quantity mismatch"})
		return
	}

	pid := c.Param("id")

	id, err := strconv.Atoi(pid)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid_id"})
		return
	}

	exisist_product, _ := a.service.Getproduct(uint(id))

	if product.Category != "" {

		exisist_product.Category = product.Category
	}

	if product.ModelID != "" {

		exisist_product.ModelID = product.ModelID
	}

	if product.Name != "" {

		exisist_product.Name = product.Name
	}

	if product.Price != 0 {

		exisist_product.Price = exisist_product.Price
	}

	if product.Quantity != 0 {

		exisist_product.Quantity = product.Quantity
	}

	if product.WarehouseID != 0 {

		exisist_product.WarehouseID = product.WarehouseID
	}

	if product.WarehouseLocation != "" {

		exisist_product.WarehouseLocation = product.WarehouseLocation
	}

	for i := range product.Assets {

		exisist_product.Assets[i].SerialNumber = product.Assets[i].SerialNumber
	}

	dberr := a.service.Updateproduct(exisist_product)

	if dberr != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to update"})
		return
	}

	c.JSON(http.StatusOK, exisist_product)
}

func (a *Product_controll) Delete_product(c *gin.Context) {

	pid := c.Param("id")

	id, err := strconv.Atoi(pid)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"Message": "Id not found"})
		return
	}

	dberr := a.service.Deleteproduct(uint(id))

	if dberr != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable_to_delete_the record"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Record has been deleted"})
}

func (a *Product_controll) In_use(c *gin.Context) {

	var borrower *models.Borrower_Request

	saved := []*models.Borrower{}

	emailto := ""
	if err := c.ShouldBindJSON(&borrower); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	if borrower.View == nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "Serialnumber required"})
		return
	}
	for _, v := range borrower.View {

		if valerr := a.validate.Struct(v); valerr != nil {

			c.JSON(http.StatusBadRequest, gin.H{"Message": valerr.Error()})
			return
		}

		check, err := a.service.Borrower_stock(v.Serial_number)

		if check {
			c.JSON(http.StatusConflict, gin.H{"Message": "This is Alreday in_use, hence Do in_stock"})
			return
		}

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Problem with Borrower_stock_status"})
			return
		}

		borrowers, err := a.service.Inuse(v.Serial_number)
		borrowers.Useby = borrower.Use_by
		borrowers.Using_location = borrower.Using_location
		borrowers.Status = "In_use"

		if borrowers.ModelID == "" {

			c.JSON(http.StatusNotFound, gin.H{"Message": "Following Serial_Number Not exist"})
			return

		}

		registerdata, err := a.service.Save_borrower(borrowers)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
			return
		}

		saved = append(saved, registerdata)
		emailto = registerdata.Useby
	}

	itemDetails := ""
	for _, item := range saved {

		itemDetails += fmt.Sprintf("SerialNumber :%v, Model: %v\n",
			item.Serial_number, item.Name)

	}
	subject := "Product In Use Confirmation"
	body := fmt.Sprintf("Hi, %v the following product has been assign to you.", emailto, itemDetails)
	go func() {
		if err := sendEmail(emailto, subject, body); err != nil {
			// Log the error but still return success
			fmt.Println("Failed to send email:", err)
		}
	}()
	c.JSON(http.StatusOK, saved)
}

func (a *Product_controll) Filer_by_user(c *gin.Context) {

	user_name := c.Param("use_by")

	users, err := a.service.View_by_user(user_name)

	if len(users) == 0 {

		c.JSON(http.StatusNotFound, gin.H{"Message": "No More recods found by the following user"})
		return
	}

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (a *Product_controll) Make_instock(c *gin.Context) {

	var temp []Multiple_serialnumber
	to := ""
	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "BadRequest"})
		return
	}

	item_details := ""
	for _, val := range temp {

		if err := a.validate.Struct(val); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		to, _ = a.service.Getby_username(val.Serial_number)

		status, err := a.service.Do_Instock(val.Serial_number)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
			return
		}

		if !status {

			c.JSON(http.StatusOK, gin.H{"Message": "This is Alredy in_stock."})
			return
		}

		item_details += fmt.Sprintf("Serial_number %v\n", val.Serial_number)
	}

	subject := "Product has been Moved to inventory"
	body := fmt.Sprintf("Hi %v Following assets moved to Instock", to, item_details)

	go func() {
		if err := sendEmail(to, subject, body); err != nil {

			fmt.Println("Failed to send email:", err)
		}
	}()
	c.JSON(http.StatusOK, gin.H{"Message": "Assest has been moved to In_stock"})
}

func (a *Product_controll) Inventory_view(c *gin.Context) {

	status := c.Query("status")

	var inventoryll []models.Inventory

	products, err := a.service.GetALLproduct()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	for i := 0; i < len(products); i++ {

		for j := 0; j < len(products[i].Assets); j++ {
			var inventory models.Inventory

			inventory.Category = products[i].Category
			inventory.ModelID = products[i].ModelID
			inventory.Name = products[i].Name
			inventory.SerialNumber = products[i].Assets[j].SerialNumber
			inventory.WarehouseID = products[i].WarehouseID
			inventory.WarehouseLocation = products[i].WarehouseLocation

			check, _ := a.service.Borrower_stock(inventory.SerialNumber)

			if !check {

				inventory.Status = "instock"

			} else {

				inventory.Status = "inuse"

				inventory.Borrower_by, _ = a.service.Getby_username(inventory.SerialNumber)
			}

			inventoryll = append(inventoryll, inventory)

		}

	}

	switch status {

	case "instock":

		for i := range inventoryll {

			if inventoryll[i].Status == "instock" {

				c.JSON(http.StatusOK, inventoryll[i])
			}

		}
		return
	case "inuse":

		for i := range inventoryll {

			if inventoryll[i].Status == "inuse" {

				c.JSON(http.StatusOK, inventoryll[i])
			}

		}
		return
	}

	c.JSON(http.StatusOK, inventoryll)
}

func sendEmail(to, subject, body string) error {

	from := "vijaytsk001@gmail.com"
	password := "eqflrzvkhhodeocn" // "pveg jitu hvza ikta"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" + // Added content type
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
}

// Visit: https://myaccount.google.com/apppasswords
