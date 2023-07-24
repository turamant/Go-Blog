package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is home page"))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8000", mux))
}