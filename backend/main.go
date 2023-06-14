package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB_USER     = os.Getenv("POSTGRES_USER")
	DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	DB_NAME     = os.Getenv("POSTGRES_DATABASE")
	DB_HOST     = os.Getenv("POSTGRES_HOST")
)

type Root struct {
	Message string `json:"message"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(&Root{"Hello world!"})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(json)
}

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", getRoot)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
