package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id    string  `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductInventory struct {
	Product  Product
	Quantity int
}

var inventory []ProductInventory

func getInventory(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(inventory)
}

func getProductInventoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, productInventory := range inventory {
		if productInventory.Product.Id == id {
			json.NewEncoder(w).Encode(productInventory)
			break
		}
	}
}

func addProductInventory(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var productInventory ProductInventory
	json.Unmarshal(reqBody, &productInventory)
	inventory = append(inventory, productInventory)
	json.NewEncoder(w).Encode(productInventory)
}

func deleteProductInventory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, productInventory := range inventory {
		if productInventory.Product.Id == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			break
		}
	}

}

func updateProductInventory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newProductInventory ProductInventory
	json.Unmarshal(reqBody, &newProductInventory)

	for i, productInventory := range inventory {
		if productInventory.Product.Id == id {
			inventory[i] = newProductInventory
			break
		}
	}

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory/product", addProductInventory).Methods("POST")
	router.HandleFunc("/inventory/product/{id}", updateProductInventory).Methods("PUT")
	router.HandleFunc("/inventory/product/{id}", deleteProductInventory).Methods("DELETE")
	router.HandleFunc("/inventory/product/{id}", getProductInventoryById).Methods("GET")
	http.ListenAndServe(":3000", router)
}

func main() {

	inventory = []ProductInventory{
		ProductInventory{
			Product{"1", "p1", "Produc 1", 15.0},
			5,
		},
		ProductInventory{
			Product{},
			10,
		},
	}

	handleRequests()
}
