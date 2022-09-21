package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Server struct {
	DB     Database
	Router *mux.Router
}

func main() {
	dbName := os.Getenv("dbName")
	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")

	sqldb, err := InitializeDB(dbUser, dbPass, dbName)
	if err != nil {
		log.Fatal(err)
	}

	db := newDatabase(sqldb)

	server := initServer(db)

	log.Println("Starting Report Server...")
	log.Fatal(http.ListenAndServe(":2754", server.Router))
}

func initServer(db Database) *Server {

	server := &Server{
		DB:     db,
		Router: mux.NewRouter(),
	}
	// Project Routes
	server.Router.HandleFunc(allReportsRoute, server.allReportsHandler).Methods("GET")
	server.Router.HandleFunc(reportRoute, server.reportHandler).Methods("GET")
	server.Router.HandleFunc(addReportRoute, server.addReportHandler).Methods("POST")
	server.Router.HandleFunc(removeReportRoute, server.removeReportHandler).Methods("POST")

	return server
}
