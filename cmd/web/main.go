package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/turamant/Go-Blog/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"

)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	posts *mysql.PostModel
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}


func main(){
	addr := flag.String("addr", ":8000", "http network address")
	dsn := flag.String("dsn", "web:pass@/goblog?parseTime=true", "MySQL data source name")
	flag.Parse()

	// f, err := os.OpenFile("./tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
    
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		infoLog: infoLog,
		errorLog: errorLog,
		posts: &mysql.PostModel{DB: db},
	}
	
	server := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),

	}

	infoLog.Printf("Server start port: %s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}