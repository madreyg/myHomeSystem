package router

import (
	"github.com/gorilla/mux"
	"log"
)

func InitRouting() {
	router := mux.NewRouter().StrictSlash(true)

	log.Println(router)
}
