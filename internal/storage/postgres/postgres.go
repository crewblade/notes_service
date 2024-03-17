package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/crewblade/notes_service/internal/domain/models"
	"github.com/crewblade/notes_service/internal/storage"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"time"
)

type Storage struct {
	db *sql.DB
}

func New(connectionString string) (*Storage, error) {
	const op = "storage.postgres.New"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS notes (
		id UUID PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT,
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	)
`)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}
	return &Storage{db: db}, nil

}

func (s *Storage) CreateNote(ctx context.Context, title string, content string) (id string, err error) {
	const op = "storage.postgres.CreateNote"
	id = uuid.NewString()
	createdAt := time.Now()
	stmt, err := s.db.Prepare("INSERT INTO notes(id, title, content, created_at) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	_, err = stmt.ExecContext(ctx, id, title, content, createdAt)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return id, nil

}
func (s *Storage) GetNoteById(ctx context.Context, id string) (models.Note, error) {
	const op = "storage.postgres.GetNoteById"
	stmt, err := s.db.Prepare("SELECT id, title, content FROM notes WHERE id = $1")
	if err != nil {
		return models.Note{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRowContext(ctx, id)
	var note models.Note
	err = row.Scan(&note.Id, &note.Title, &note.Content)
	if err != nil {
		return models.Note{}, fmt.Errorf("%s: %w", op, storage.IdNotFound)
	}

	return note, nil
}
func (s *Storage) UpdateNote(ctx context.Context, id, title, content string) (models.Note, error) {
	const op = "storage.postgres.UpdateNote"

	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM notes WHERE id = $1)", id).Scan(&exists)
	if err != nil || !exists {
		return models.Note{}, fmt.Errorf("%s: %w", op, storage.IdNotFound)
	}

	_, err = s.db.ExecContext(ctx, "UPDATE notes SET title = $1, content = $2 WHERE id = $3", title, content, id)
	if err != nil {
		return models.Note{}, fmt.Errorf("%s: %w", op, err)
	}

	var updatedNote models.Note
	err = s.db.QueryRowContext(ctx, "SELECT id, title, content FROM notes WHERE id = $1", id).Scan(&updatedNote.Id, &updatedNote.Title, &updatedNote.Content)
	if err != nil {
		return models.Note{}, fmt.Errorf("%s: %w", op, err)
	}

	return updatedNote, nil
}
func (s *Storage) DeleteNote(ctx context.Context, id string) (models.Note, error) {
	const op = "storage.postgres.DeleteNote"

	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM notes WHERE id = $1)", id).Scan(&exists)
	if err != nil || !exists {
		return models.Note{}, fmt.Errorf("%s: %w", op, storage.IdNotFound)
	}
	var deletedNote models.Note
	err = s.db.QueryRowContext(ctx, "SELECT id, title, content FROM notes WHERE id = $1", id).Scan(&deletedNote.Id, &deletedNote.Title, &deletedNote.Content)
	if err != nil {
		return models.Note{}, fmt.Errorf("%s: %w", op, err)
	}

	_, err = s.db.ExecContext(ctx, "DELETE FROM notes WHERE id = $1", id)
	if err != nil {
		return models.Note{}, fmt.Errorf("%s: %w", op, err)
	}

	return deletedNote, nil
}

func (s *Storage) GetNotes(ctx context.Context, limit int32, offsetID string) ([]models.Note, string, error) {
	const op = "storage.postgres.GetNotes"

	var offsetTime time.Time
	err := s.db.QueryRowContext(ctx, "SELECT created_at FROM notes WHERE id = $1", offsetID).Scan(&offsetTime)
	if err != nil {
		//if errors.Is(err, sql.ErrNoRows) {
		//	return nil, "", fmt.Errorf("%s: %w", op, storage.IdNotFound)
		//}
		//return nil, "", fmt.Errorf("%s: %w", op, err)
		return nil, "", fmt.Errorf("%s: %w", op, storage.IdNotFound)
	}
	rows, err := s.db.QueryContext(ctx, "SELECT id, title, content FROM notes WHERE created_at >= $1 ORDER BY created_at LIMIT $2", offsetTime, limit+1)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	var notes []models.Note
	var nextOffsetID string
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.Id, &note.Title, &note.Content); err != nil {
			fmt.Println(err.Error())
			return nil, "", fmt.Errorf("%s: %w", op, err)
		}
		notes = append(notes, note)
		nextOffsetID = note.Id
	}
	// (1 2 3) 4| 5
	// limit 3 + 1
	if len(notes) == int(limit)+1 {
		notes = notes[:len(notes)-1]
	} else {
		nextOffsetID = ""
	}
	return notes, nextOffsetID, nil
}
func (s *Storage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
