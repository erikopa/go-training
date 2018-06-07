package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func main(){

	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Println("GET Home")
}
