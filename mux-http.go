package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Address1 struct {
	Address string `json: "address"`
	City    string `json: "city"`
	Provice string `json: "provice"`
	Zipcode string `json: "zipcode"`
}

type Customer struct {
	Id       int      `json: id`
	FullName string   `json:"fullName"`
	Address  Address1 `json:"address"`
}

var data []Customer = []Customer{}

func main() {
	router := mux.NewRouter() // r = router
	router.HandleFunc("/customers", addItem).Methods("POST")
	http.ListenAndServe(":8888", router)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("TEST"))
	var newData Customer
	json.NewDecoder(r.Body).Decode(&newData)
	data = append(data, newData)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}
