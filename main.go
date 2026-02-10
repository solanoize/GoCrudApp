package main

import (
	"fmt"
	"gocrud/products"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database", err.Error())
	}

	db.AutoMigrate(&products.Product{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	var productRepository products.IProductRepository = products.ProductRepositoryImpl(db)
	var productService *products.ProductService = products.NewProductService(productRepository)

	r.Mount("/products", products.URLPattern(productService))

	fmt.Println("Server is running on port 8080")
	fmt.Println("Press Ctrl+C to stop the server")
	http.ListenAndServe(":8080", r)

}
