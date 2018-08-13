package main

import(
	//"fmt"
	//"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
)


type Items struct{
	ID 			int 		`json: "id"`
	Name 		string 		`json: "name"`
	Price		int			`json: "price"`
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
	json.NewDecoder(r.Body).Decode(&tmp)
	itemNo++;
	tmp.ID = itemNo
	items = append(items,tmp)
	json.NewEncoder(w).Encode(tmp)

}
func updateItem(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	in := mux.Vars(r)
	for idx, tmp := range items{
		var cnv string
		cnv = strconv.Itoa(tmp.ID)
		if(in["id"] == cnv){
			items = append(items[:idx] , items[idx + 1:]...)
			w.Header().Set("Content-Type","application/json")
			var tmp2 Items
			json.NewDecoder(r.Body).Decode(&tmp2)
			tmp2.ID = tmp.ID

			//if the user wishes to edit 1 or 2 variables of an instead of 
			// all of them
			
			if(tmp2.Name == "") {
				tmp2.Name = tmp.Name
			}
			if(tmp2.Price == 0){
				tmp2.Price = tmp.Price;
			}
			if(tmp2.Quantity == 0){
				tmp2.Quantity = tmp.Quantity
			}
			items = append(items,tmp2)
			json.NewEncoder(w).Encode(tmp2)
			return
		}
	}
	json.NewEncoder(w).Encode(items)
	
	
}
func deleteItem(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	in := mux.Vars(r)
	for idx, tmp := range items{
		var cnv string
		cnv = strconv.Itoa(tmp.ID)
		if(in["id"] == cnv){
			items = append(items[:idx], items[idx + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)

	
}

func HomePage(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte("Welcome to the Shopping Server.\n"))
}

func main() {
	mx := mux.NewRouter()
	
	items = append(items,Items{ID: 1 , Name : "Shirt" , Price: 1150 , Quantity: 2})
	items = append(items,Items{ID: 2 , Name : "Pant" , Price: 1250 , Quantity: 2})

	mx.HandleFunc("/",HomePage)
	mx.HandleFunc("/shop/items",showItems).Methods("GET")
	mx.HandleFunc("/shop/items",addItems).Methods("POST")
	mx.HandleFunc("/shop/items/{id}",updateItem).Methods("PUT")
	mx.HandleFunc("/shop/items/{id}",deleteItem).Methods("DELETE")

	http.ListenAndServe(":8080",mx)
}

