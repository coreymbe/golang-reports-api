package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Database interface {
	GetAllReports() ([]*Report, error)
	GetReport(r_ID int) (*Report, error)
	AddReport(certname string, environment string, status string, time string, transaction_uuid string) error
	RemoveReport(transaction_uuid string) error
}

type database struct {
	db *sql.DB
}

func InitializeDB(user, password, dbname string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("user=%s password= %s dbname=%s sslmode=disable", user, password, dbname)
	pdb, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return pdb, nil
}

func newDatabase(db *sql.DB) Database {
	return &database{
		db: db,
	}
}
