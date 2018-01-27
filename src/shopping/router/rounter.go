package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func InitRouting() * mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products", home).Methods("GET")
	return router
}


func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Run list of products.")
}