package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Customers struct {
	FullName string `json:"fullName"`
}

var data []Customers = []Customers{}

func main() {
	router := mux.NewRouter() // r = router
	router.HandleFunc("/customers", addItem).Methods("POST")
	http.ListenAndServe(":8888", router)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("TEST"))
	var newData Customers
	json.NewDecoder(r.Body).Decode(&newData)
	data = append(data, newData)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	
}