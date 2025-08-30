package main

import (
	controllers "inventory_management/Controllers"
	models "inventory_management/Models"
	service "inventory_management/Service"
	"inventory_management/dbrepository"
	"inventory_management/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

/*

The Go runtime automatically calls all init() functions in a package before your main() function runs.

*/

func main() {

	dsn := "root:Vijay@123@tcp(localhost:3306)/inventory_management?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Order{}, &models.OrderItem{}, &models.Product{}, &models.ProductItemInput{}, &models.Borrower{}, &models.User{})
	log.Println("Hello")
	//db.Migrator().DropColumn(&models.Product{}, "status") removing the existing table in db

	if err != nil {

		log.Fatal(err.Error())
	}

	server := gin.New()

	server.Use(gin.Logger())

	unser_db := dbrepository.User_repo(db)

	order_db := dbrepository.OrderRepo(db)

	order_service := service.NewOrderService(order_db)

	order_controll := controllers.Order_controll(order_service)

	productdb := dbrepository.Product_repo(db)

	product_service := service.New_Product_service(productdb)

	product_controller := controllers.Product_con(product_service)

	user_controll := controllers.User_con(unser_db)

	routes.Unser_routes(server, *user_controll)

	routes.RegisterRoutes(server, *order_controll)

	routes.Register_product_routes(server, *product_controller)

	port := os.Getenv("PORT")
	server.Run("localhost:" + port)
}

/*

http://localhost:8080/

*/
