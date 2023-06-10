package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karomahlailatul/go_simple_ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/register", controllers.Register())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductAddAdmin())

	incomingRoutes.GET("/users/searchproduct", controllers.SearchProduct())
	incomingRoutes.GET("/users/productview", controllers.ProductView())
}
