package main

import (
	"fmt"
	"net/http"
	"valid-cards/pkg"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", pkg.PostAndValidateCards).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Error happened, %v\n", err.Error())
		return
	}
}
