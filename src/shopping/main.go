package main

import (
	"fmt"
	"log"
	"net/http"
	"shopping/router"
)

func main() {
	//http.HandleFunc("/", handler)
	rout := router.InitRouting()
	log.Println("Start HomeSystem")
	http.ListenAndServe(":8080", rout)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
