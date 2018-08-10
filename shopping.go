package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Items struct{
	ID 			int 		`json: "id"`
	Name 		string 		`json: "name"`
	Price		string		`json: "price"`
	Quantity 	int			`json: "quantity"`

}

func showItems(w http.Responsewriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encoder(books)
	
}

func addItems(w http.Responsewriter,r *http.Request){

}
func updateItem(w http.Responsewriter,r *http.Request) {
	
}
func deleteItem(w http.Responsewriter,r *http.Request) {
	
}

func main() {
	mx := mux.NewRouter()
	
	mx.HandleFunc("/api/items",showItems).Methods("GET")
	mx.HandleFunc("/api/items",addItems).Methods("POST")
	mx.HandleFunc("/api/items/{id}",updateItem).Methods("PUT")
	mx.HandleFunc("/api/items/{id}",deleteItem).Methods("DELETE")

	http.ListenAndServe(":8080",mx)
}

