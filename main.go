package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
}

var products []Product

func init() {
	products = []Product{
		{
			ID:          "1",
			Name:        "Product 1",
			Description: "Description for product 1",
			Price:       100.0,
		},
		{
			ID:          "2",
			Name:        "Product 2",
			Description: "Description for product 2",
			Price:       200.0,
		},
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/api/products", handleProducts)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
