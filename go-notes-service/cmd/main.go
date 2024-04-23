package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mynotes/internal/handlers"
	"mynotes/internal/db"
)

func main() {
	db.ConnectPostgresDB()
	defer db.ClosePostgresDB()

	r := mux.NewRouter()

	r.HandleFunc("/notes", handlers.GetAllNotes).Methods("GET")
	r.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	r.HandleFunc("/notes/{id}", handlers.GetNoteByID).Methods("GET")
	r.HandleFunc("/notes/{id}", handlers.UpdateNoteByID).Methods("PATCH")
	r.HandleFunc("/notes/{id}", handlers.DeleteNoteByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}