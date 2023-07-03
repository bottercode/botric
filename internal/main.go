package main

import (
	"log"
	"net/http"
)

func main(){
	
	log.Println("Listening on localhost")
	log.Fatal(http.ListenAndServe(":8080", nil))
}