package routes

import (
	controllers "inventory_management/Controllers"

	"github.com/gin-gonic/gin"
)

func Unser_routes(router *gin.Engine, user controllers.User_controller) {

	router.POST("/Signup", user.Signup)
	router.POST("/Login", user.Login)

}
