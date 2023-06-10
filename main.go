package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/karomahlailatul/go_simple_ecommerce/controllers"
	"github.com/karomahlailatul/go_simple_ecommerce/database"
	"github.com/karomahlailatul/go_simple_ecommerce/middlewares"
	"github.com/karomahlailatul/go_simple_ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Product"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.use(middlewares.Authenticate())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/checkout", app.BuyFromCart())
	router.GET("/instanbuy", app.InstanBuy())

	log.Fatal(router.run(":" + port))
}
