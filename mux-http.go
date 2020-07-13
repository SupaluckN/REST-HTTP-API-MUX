package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type FullName struct {
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
}

type Address1 struct {
	Address string `json: "address"`
	City    string `json: "city"`
	Provice string `json: "provice"`
	Zipcode string `json: "zipcode"`
}

type Customer struct {
	Id       int      `json: id`
	FullName FullName `json:"fullName"`
	Address  Address1 `json:"address"`
	Email    string   `json: "email"`
	Phone    string   `json: "phone"`
}

var data []Customer = []Customer{}

func main() {
	r := mux.NewRouter() // r = router
	r.HandleFunc("/customers", addItem).Methods("POST")
	http.ListenAndServe(":8888", r)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Customers Details"))
	var newData Customer
	json.NewDecoder(r.Body).Decode(&newData)
	data = append(data, newData)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}
