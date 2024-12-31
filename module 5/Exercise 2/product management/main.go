package main

import (
	"fmt"
	"inventory_server/config"
	"inventory_server/controller"
	"inventory_server/middleware"
	"inventory_server/repository"
	"inventory_server/service"
	"log"
	"net/http"
)

func main() {
	config.InitialiseDatabase()

	productRepo := repository.NewProductRepository(config.GetDB())
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	http.HandleFunc("/products", productController.GetAllProducts)
	http.HandleFunc("/products/", productController.GetProduct)
	http.HandleFunc("/products/create", productController.CreateProduct)
	http.HandleFunc("/products/update/", productController.UpdateProduct)
	http.HandleFunc("/products/delete/", productController.DeleteProduct)

	http.HandleFunc("/products", middleware.LoggingMiddleware(http.HandlerFunc(productController.GetAllProducts)))
	http.HandleFunc("/products/", middleware.LoggingMiddleware(http.HandlerFunc(productController.GetProduct)))

	port := ":8000"
	fmt.Printf("Inventory Service is running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
