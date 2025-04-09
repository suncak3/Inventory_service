package main

import (
	"github.com/gin-gonic/gin"
	"inventory-service/db"
	"inventory-service/handler"
)

func main() {
	router := gin.Default()
	productHandler := handler.NewHandler()

	defer db.CloseConnection()

	router.GET("/products", productHandler.GetAllProducts)
	router.GET("/products/:id", productHandler.GetProductByID)
	router.POST("/products", productHandler.CreateProduct)
	router.PATCH("/products/:id", productHandler.UpdateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	router.Run("0.0.0.0:8082")
}
