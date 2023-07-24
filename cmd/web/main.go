package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}



func main(){
	addr := flag.String("addr", ":8000", "http network address")
	flag.Parse()
	// f, err := os.OpenFile("./tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
    
	app := &application{
		infoLog: infoLog,
		errorLog: errorLog,
	}
	
	server := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),

	}

	infoLog.Printf("Server start port: %s", *addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}