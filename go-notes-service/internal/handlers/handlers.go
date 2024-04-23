package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
    "strconv"

	"mynotes/internal/auth"
	"mynotes/internal/db"
	"mynotes/internal/models"

	"github.com/gorilla/mux"
)

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")
    if bearerToken == "" {
        publicNotes, err := db.GetAllPublicNotes()
        if err != nil {
            http.Error(w, "Failed to get public notes from database", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(publicNotes)
        return
    }

    token := strings.Split(bearerToken, " ")
    if len(token) != 2 {
        http.Error(w, "Invalid authorization token", http.StatusBadRequest)
        return
    }

    userID, err := auth.GetUserIdFromToken(token[1])
    if err != nil {
        publicNotes, err := db.GetAllPublicNotes()
        if err != nil {
            http.Error(w, "Failed to get public notes from database", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(publicNotes)
        return
    }

    publicNotes, err := db.GetAllPublicNotes()
    if err != nil {
        http.Error(w, "Failed to get public notes from database", http.StatusInternalServerError)
        return
    }

    privateNotes, err := db.GetUserPrivateNotes(int(userID))
    if err != nil {
        http.Error(w, "Failed to get user's private notes from database", http.StatusInternalServerError)
        return
    }

    allNotes := append(publicNotes, privateNotes...)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(allNotes)
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	noteID := params["id"]
	
	note, err := db.GetNoteByID(noteID)
	if err != nil {
		http.Error(w, "Failed to fetch note", http.StatusInternalServerError)
		return
	}
	
	if note == nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

    if note.IsPublic == false {
        bearerToken := r.Header.Get("Authorization")
        if bearerToken == "" {
            http.Error(w, "Authorization token required", http.StatusUnauthorized)
            return
        }
    
        token := strings.Split(bearerToken, " ")
        if len(token) != 2 {
            http.Error(w, "Invalid authorization token", http.StatusBadRequest)
            return
        }

        userID, err := auth.GetUserIdFromToken(token[1])
        if err != nil {
            http.Error(w, "Failed to authenticate user", http.StatusUnauthorized)
            return
        }

        if note.AuthorID != userID {
            http.Error(w, "Forbidden", http.StatusForbidden)
        }
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
    bearerToken := r.Header.Get("Authorization")
    if bearerToken == "" {
        http.Error(w, "Authorization token required", http.StatusUnauthorized)
        return
    }

    token := strings.Split(bearerToken, " ")
    if len(token) != 2 {
        http.Error(w, "Invalid authorization token", http.StatusBadRequest)
        return
    }

    userID, err := auth.GetUserIdFromToken(token[1])
    if err != nil {
        http.Error(w, "Failed to authenticate user", http.StatusUnauthorized)
        return
    }

    var note models.Note
    err = json.NewDecoder(r.Body).Decode(&note)
    if err != nil {
        http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
        return
    }

    fmt.Println(note)

    note.CreatedAt = time.Now()
    note.AuthorID = int(userID)

    noteID, err := db.CreateNote(&note)
    if err != nil {
        fmt.Println(err)
        http.Error(w, "Failed to create note", http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "message": "Note created successfully",
        "note_id": noteID,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func UpdateNoteByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    noteID := params["id"]

    var updatedNote models.Note
    err := json.NewDecoder(r.Body).Decode(&updatedNote)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    note, err := db.GetNoteByID(noteID)
    if err != nil {
        http.Error(w, "Note not found", http.StatusNotFound)
        return
    }

    bearerToken := r.Header.Get("Authorization")
    if bearerToken == "" {
        http.Error(w, "Authorization token required", http.StatusUnauthorized)
        return
    }

    token := strings.Split(bearerToken, " ")
    if len(token) != 2 {
        http.Error(w, "Invalid authorization token", http.StatusBadRequest)
        return
    }

    userID, err := auth.GetUserIdFromToken(token[1])
    if err != nil {
        http.Error(w, "Failed to authenticate user", http.StatusUnauthorized)
        return
    }

    noteAuthorID := note.AuthorID

    if userID != noteAuthorID {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    err = db.UpdateNoteByID(noteID, &updatedNote)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}


func DeleteNoteByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    noteID := params["id"]

    userID, err := db.GetUserIDFromNote(noteID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    bearerToken := r.Header.Get("Authorization")
    if bearerToken == "" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    token := strings.Split(bearerToken, " ")
    if len(token) != 2 {
        http.Error(w, "Invalid authorization token format", http.StatusBadRequest)
        return
    }

    userIDFromToken, err := auth.GetUserIdFromToken(token[1])
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    stringUserIDFromToken := strconv.Itoa(userIDFromToken)

    if userID != stringUserIDFromToken {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    err = db.DeleteNoteByID(noteID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
