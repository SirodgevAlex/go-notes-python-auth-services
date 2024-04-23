package db

import (
	"database/sql"
	"fmt"
	"log"
    "errors"

	_ "github.com/lib/pq"
	"mynotes/internal/models"
)

var db *sql.DB

func ConnectPostgresDB() error {
    connStr := "postgres://postgres:1234@host.docker.internal:5432/jet-style?sslmode=disable"
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        return err
    }
    err = db.Ping()
    if err != nil {
        return err
    }
    log.Println("Connected to PostgreSQL database")
    return nil
}

func ClosePostgresDB() {
    if db != nil {
        db.Close()
        log.Println("Disconnected from PostgreSQL database")
    }
}

func GetPostgresDB() (*sql.DB, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetNoteByID(id string) (*models.Note, error) {
	var note models.Note
	err := db.QueryRow("SELECT id, created_at, author_id, text, is_public FROM notes WHERE id = $1", id).Scan(
		&note.ID, &note.CreatedAt, &note.AuthorID, &note.Text, &note.IsPublic,
	)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("note not found")
	case err != nil:
		return nil, err
	}
	return &note, nil
}

func CreateNote(note *models.Note) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO notes(created_at, author_id, text, is_public) VALUES($1, $2, $3, $4) RETURNING id",
        note.CreatedAt.Format("2006-01-02 15:04:05"), note.AuthorID, note.Text, note.IsPublic,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to insert note into database: %v", err)
	}
	if id == 0 {
		return 0, errors.New("failed to get ID of inserted note")
	}
	return id, nil
}

func GetAllPublicNotes() ([]models.Note, error) {
    rows, err := db.Query("SELECT id, created_at, author_id, text, is_public FROM notes WHERE is_public = true")
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    defer rows.Close()

    notes := []models.Note{}

    for rows.Next() {
        var note models.Note
        err := rows.Scan(&note.ID, &note.CreatedAt, &note.AuthorID, &note.Text, &note.IsPublic)
        if err != nil {
            log.Fatal(err)
            return nil, err
        }
        notes = append(notes, note)
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
        return nil, err
    }

    return notes, nil
}

func GetUserPrivateNotes(userID int) ([]models.Note, error) {
    rows, err := db.Query("SELECT id, created_at, author_id, text FROM notes WHERE author_id = $1 AND is_public = false", userID)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    defer rows.Close()

    notes := []models.Note{}

    for rows.Next() {
        var note models.Note
        err := rows.Scan(&note.ID, &note.CreatedAt, &note.AuthorID, &note.Text)
        if err != nil {
            log.Fatal(err)
            return nil, err
        }
        notes = append(notes, note)
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
        return nil, err
    }

    return notes, nil
}

func UpdateNoteByID(id string, updatedNote *models.Note) error {
    query := `
        UPDATE notes
        SET text = COALESCE(NULLIF($1, ''), text),
            is_public = COALESCE($2, is_public)
        WHERE id = $3
    `

    _, err := db.Exec(query, updatedNote.Text, updatedNote.IsPublic, id)
    if err != nil {
        return fmt.Errorf("failed to update note: %v", err)
    }

    return nil
}

func DeleteNoteByID(id string) (error) {
	result, err := db.Exec("DELETE FROM notes WHERE id=$1", id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("no note found with the provided ID")
	}

    return nil
}

func GetUserIDFromNote(id string) (string, error) {
    var authorID string
    err := db.QueryRow("SELECT author_id FROM notes WHERE id = $1", id).Scan(&authorID)
    if err != nil {
        return "", err
    }
    return authorID, nil
}