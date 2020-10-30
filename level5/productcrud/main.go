// main.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id    string  `json:"Id"`
	Code  string  `json:"Code"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
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
		}
	}

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory", getInventory)
	router.HandleFunc("/inventory/product", addProductInventory).Methods("POST")
	router.HandleFunc("/inventory/product/{id}", deleteProductInventory).Methods("DELETE")
	router.HandleFunc("/inventory/product/{id}", getProductInventoryById)
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
