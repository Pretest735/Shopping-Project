package main

import (
	//"fmt"
	//"log"
	"time"
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

type User struct{
	UserId		int   	`json : "userid"`
	UserName	string 	`json : "Username"`
	PassWord	string	`json : "Password"`
}

var exist_user = make(map[string]User)

var items []Items
var itemNo int = 2
var userNo int = 0
var users []User

func alreadyLoggedIn(r *http.Request) bool {
	cookie , err := r.Cookie("username")
	if err == nil{
		_,flag := exist_user[cookie.Value]
		return flag;
	} else{
		return false
	}
}

func showItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","application/json")
	if alreadyLoggedIn(r) == false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok : 0, Message : "Login first or register to create account"})
		return
	}
	err := json.NewEncoder(w).Encode(items)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Failed to show Information."})
		return
	}

}

func addItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","application/json")
	if alreadyLoggedIn(r) == false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok : 0, Message : "Login first or register to create account"})
		return
	}
	var tmp Items
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err == nil {
		if tmp.Name == "" || tmp.Price <= 0 || tmp.Quantity <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Invalid Information."})
			return

		}
		

		for  idx := range items{
			if tmp.Name == items[idx].Name {
				json.NewEncoder(w).Encode(Response{Ok : 1, Message : "Item already exists"})
				return
			}
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
	w.Header().Set("content-type","application/json")
	if alreadyLoggedIn(r) == false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok : 0, Message : "Login first or register to create account"})
		return
	}
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

					if tmp2.Name == "" || tmp2.Price <= 0 || tmp2.Quantity <= 0 {
						w.WriteHeader(http.StatusBadRequest)
						json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Invalid Information"})
						return
					}

					items[idx] = tmp2
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
	w.Header().Set("content-type","application/json")
	if alreadyLoggedIn(r) == false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok : 0, Message : "Login first or register to create account"})
		return
	}
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
	w.Header().Set("content-type","application/json")
	if alreadyLoggedIn(r) == false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok : 0, Message : "Login first or register to create account"})
		return
	}
	err := json.NewEncoder(w).Encode(Response{Ok : 1, Message : "Welcome to the server"})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok: 0, Message: "Failed to process Request."})
		return
	}
}


func registerUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	if alreadyLoggedIn(r) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok : 0, Message : "Logout first to register"})
		return
	}

	var tmp User
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err == nil{
		_,flag := exist_user[tmp.UserName]
		if flag == true{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Ok : 0, Message : "User already exists!!!"})
			return
		} else{
			userNo++
			tmp.UserId = userNo
			users = append(users,tmp)
			exist_user[tmp.UserName] = tmp
			json.NewEncoder(w).Encode(Response{Ok : 1, Message : "Registration successful."})
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok : 0, Message : "Failed to process Request"})
	}
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","application/json")
	if alreadyLoggedIn(r)  {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok : 0, Message : "Please logut to login with another account"})
		return
	}
	var tmp User
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err == nil {
		tmp2,flag := exist_user[tmp.UserName]
		if flag == false {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Ok : 0, Message : "Invalid UserName"})
			return
		} else {
			if tmp2.PassWord != tmp.PassWord {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(Response{Ok : 0, Message : "Incorrect password."})
				return
			} else{
				c := http.Cookie{Name : "username" , Value : tmp.UserName}
				http.SetCookie(w, &c)
				json.NewEncoder(w).Encode(Response{Ok : 1, Message : "Login successful."})
				return
			}
		}
	} else {
		json.NewEncoder(w).Encode(Response{Ok : 0, Message : "Failed to process Request."})
		return
	}
}

func logoutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","application/json")
	if alreadyLoggedIn(r) == false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Ok : 0, Message : "You are not logged in."})
		return
	}

	logout := http.Cookie{Name : "username" , Value : "" , Expires : time.Now()}
	http.SetCookie(w,&logout)
	json.NewEncoder(w).Encode(Response{Ok : 1, Message : "Successfully logged out"})

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
	mx.HandleFunc("/shop/register", registerUser).Methods("POST")
	mx.HandleFunc("/shop/login", loginUser).Methods("POST")
	mx.HandleFunc("/shop/logout", logoutUser).Methods("GET")
	


	http.ListenAndServe(":8080", mx)
}
