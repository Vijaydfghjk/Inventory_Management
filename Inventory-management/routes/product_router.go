package routes

import (
	controllers "inventory_management/Controllers"
	"inventory_management/middleware"

	"github.com/gin-gonic/gin"
)

func Register_product_routes(router *gin.Engine, pdu_controll controllers.Product_controll) {

	rolls := router.Group("/")

	rolls.Use(middleware.AuthMiddleware("member"))

	rolls.POST("/Products", pdu_controll.Add_product)
	rolls.GET("/Products/:id", pdu_controll.Viewby_id)
	rolls.GET("/Products", pdu_controll.View)
	rolls.DELETE("/Products/:id", pdu_controll.Delete_product)
	rolls.PUT("/Products/:id", pdu_controll.Update_Product)
	rolls.DELETE("/Instock/", pdu_controll.Make_instock)

	rolls.POST("/Inuse", pdu_controll.In_use)
	rolls.GET("/Username/:use_by", pdu_controll.Filer_by_user)
	rolls.GET("/inventory_view", pdu_controll.Inventory_view)

	// http://localhost:8080/inventory_view?status=inuse

}
