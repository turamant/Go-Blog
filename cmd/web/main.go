package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)



func main(){
	addr := flag.String("addr", ":8000", "http network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/post", showPost)
	mux.HandleFunc("/post/create", createPost)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Server start port: %s", *addr)
	errorLog.Fatal(http.ListenAndServe(*addr, mux))
}