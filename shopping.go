package main

import (
	//"fmt"
	//"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Items struct {
	ID       int    `json: "id,omitempty"`
	Name     string `json: "name,omitempty"`
	Price    int    `json: "price,omitempty"`
	Quantity int    `json: "quantity,omitempty"`
}
type Response struct {
	Ok      int    `json : "ok"`
	Message string `json : "Message"`
}

var items []Items
var itemNo int = 2

func showItems(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(items)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Failed to show Information."})
		return
	}

}

func addItems(w http.ResponseWriter, r *http.Request) {
	var tmp Items
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err == nil {
		if tmp.Name == "" || tmp.Price < 0 || tmp.Quantity < 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Invalid Information."})
			return

		}
		itemNo++
		tmp.ID = itemNo
		items = append(items, tmp)
		err = json.NewEncoder(w).Encode(tmp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Failed to add item."})
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Failed to process Request."})
		return
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	in := mux.Vars(r)
	cnv, err := strconv.Atoi(in["id"])
	var sc bool = false
	if err == nil {
		for idx, tmp := range items {
			if tmp.ID == cnv {
				sc = true

				var tmp2 Items
				err = json.NewDecoder(r.Body).Decode(&tmp2)
				if err == nil {
					tmp2.ID = tmp.ID

					if tmp2.Name == "" || tmp2.Price < 0 || tmp2.Quantity < 0 {
						w.WriteHeader(http.StatusBadRequest)
						json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Invalid Information"})
						return
					}

					items[idx] = tmp2
					json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Succesfully updated Information."})
					break
				} else {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Failed to receive update Information."})
					return
				}
			}
		}
		if sc == false {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Ok: 0, Message: "The given entry is not found."})
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Failed to process Request."})
		return
	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Failed to update Information"})
		return
	}

}
func deleteItem(w http.ResponseWriter, r *http.Request) {

	in := mux.Vars(r)
	var sc bool = false
	cnv, err := strconv.Atoi(in["id"])
	if err == nil {
		for idx, tmp := range items {
			if tmp.ID == cnv {
				sc = true
				items = append(items[:idx], items[idx+1:]...)
				break
			}
		}
		if sc == false {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Ok: 0, Message: "The given entry is not found."})
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Failed to process Request."})
		return
	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Failed to delete items."})
		return
	}

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	pr := "Welcome to the Shopping Server."
	err := json.NewEncoder(w).Encode(pr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Failed to process Request."})
		return
	}
}

func main() {
	mx := mux.NewRouter()

	items = append(items, Items{ID: 1, Name: "Shirt", Price: 1150, Quantity: 2})
	items = append(items, Items{ID: 2, Name: "Pant", Price: 1250, Quantity: 2})

	mx.HandleFunc("/shop", HomePage)
	mx.HandleFunc("/shop/items", showItems).Methods("GET")
	mx.HandleFunc("/shop/items", addItems).Methods("POST")
	mx.HandleFunc("/shop/items/{id}", updateItem).Methods("PUT")
	mx.HandleFunc("/shop/items/{id}", deleteItem).Methods("DELETE")

	http.ListenAndServe(":8080", mx)
}
