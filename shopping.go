package main

import(
	//"fmt"
	//"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)


type Items struct{
	ID 			int 		`json: "id"`
	Name 		string 		`json: "name"`
	Price		int		`json: "price"`
	Quantity 	int			`json: "quantity"`

}

var items []Items
var itemNo int = 0

func showItems(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(items)
	
}

func addItems(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var tmp Items
	_ = json.NewDecoder(r.Body).Decode(&tmp)
	itemNo++;
	tmp.ID = itemNo
	items = append(items,tmp)
	json.NewEncoder(w).Encode(tmp)

}
func updateItem(w http.ResponseWriter,r *http.Request) {
	
}
func deleteItem(w http.ResponseWriter,r *http.Request) {
	
}

func HomePage(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte("Welcome to the Shopping Server.\n"))
}

func main() {
	mx := mux.NewRouter()
	
	items = append(items,Items{ID: 1 , Name : "Shirt" , Price: 1150 , Quantity: 2})
	items = append(items,Items{ID: 2 , Name : "Pant" , Price: 1250 , Quantity: 2})

	mx.HandleFunc("/",HomePage)
	mx.HandleFunc("/api/items",showItems).Methods("GET")
	mx.HandleFunc("/api/items",addItems).Methods("POST")
	mx.HandleFunc("/api/items/{id}",updateItem).Methods("PUT")
	mx.HandleFunc("/api/items/{id}",deleteItem).Methods("DELETE")

	http.ListenAndServe(":8080",mx)
}

