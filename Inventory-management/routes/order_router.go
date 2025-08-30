package routes

import (
	controllers "inventory_management/Controllers"
	"inventory_management/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, order_controll controllers.OrderController) {

	rolls := router.Group("/")

	rolls.Use(middleware.AuthMiddleware("admin"))

	rolls.POST("/orders", order_controll.CreateOrder)
	rolls.GET("/orders", order_controll.GetAllOrders)
	rolls.PUT("/orders/:id", order_controll.UpdateOrder)
	rolls.GET("/orders/:id", order_controll.GetOrder)
	rolls.DELETE("orders/:id", order_controll.DeleteOrder)
}
