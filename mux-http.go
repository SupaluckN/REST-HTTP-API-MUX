package main

import (
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()	
	router.HandleFunc("/customers", addItem)
	http.ListenAndServe(":8888", router)
}

func addItem(w http.ResponseWriter, r*http.Request) {
	w.Write([]byte("TEST"))
}