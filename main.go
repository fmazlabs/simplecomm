package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
)

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

var products []Product

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, item := range products {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Product{})
}

func main() {
	products = append(products, Product{ID: "1", Name: "Product 1", Price: 100})
	products = append(products, Product{ID: "2", Name: "Product 2", Price: 200})

	router := mux.NewRouter()
	router.HandleFunc("/api/products", GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", GetProduct).Methods("GET")

	htmlBytes, err := ioutil.ReadFile("web/dist/index.html")
	if err != nil {
		log.Fatal(err)
	}

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(htmlBytes)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
