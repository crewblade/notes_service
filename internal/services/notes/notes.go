package notes

import (
	"context"
	"errors"
	"fmt"
	"github.com/crewblade/notes_service/internal/domain/models"
	"github.com/crewblade/notes_service/internal/storage"
	"log/slog"
)

type Notes struct {
	log            *slog.Logger
	noteCreator    NoteCreator
	noteGetterById NoteGetterById
	noteUpdater    NoteUpdater
	noteDeleter    NoteDeleter
	noteLister     NoteLister
}

type NoteCreator interface {
	CreateNote(ctx context.Context, title string, content string) (id string, err error)
}
type NoteGetterById interface {
	GetNoteById(ctx context.Context, id string) (models.Note, error)
}
type NoteUpdater interface {
	UpdateNote(ctx context.Context, id, title, content string) (models.Note, error)
}
type NoteDeleter interface {
	DeleteNote(ctx context.Context, id string) (models.Note, error)
}
type NoteLister interface {
	GetNotes(ctx context.Context, limit int32, offset_id string) (notes []models.Note, next_offset_id string, err error)
}

func New(
	log *slog.Logger,
	noteCreator NoteCreator,
	noteGetterById NoteGetterById,
	noteUpdater NoteUpdater,
	noteDeleter NoteDeleter,
	noteLister NoteLister,
) *Notes {
	return &Notes{
		log:            log,
		noteCreator:    noteCreator,
		noteGetterById: noteGetterById,
		noteUpdater:    noteUpdater,
		noteDeleter:    noteDeleter,
		noteLister:     noteLister,
	}
}

func (n *Notes) CreateNote(ctx context.Context, title string, content string) (id string, err error) {
	const op = "services.notes.CreateNote"
	log := n.log.With(slog.String("op", op))
	id, err = n.noteCreator.CreateNote(ctx, title, content)
	if err != nil {
		log.Warn("err:" + err.Error())
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("Note created", slog.Any("id", id))
	return

}
func (n *Notes) GetNoteById(ctx context.Context, id string) (models.Note, error) {
	const op = "services.notes.GetNoteById"
	log := n.log.With(slog.String("op", op))
	note, err := n.noteGetterById.GetNoteById(ctx, id)
	if err != nil {
		if errors.Is(err, storage.IdNotFound) {
			log.Warn("Id not found", slog.String("err", err.Error()))
		}
		return note, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("Note received", slog.Any("note", note))
	return note, nil

}
func (n *Notes) UpdateNote(ctx context.Context, id, title, content string) (models.Note, error) {
	const op = "services.notes.UpdateNote"
	log := n.log.With(slog.String("op", op))
	note, err := n.noteUpdater.UpdateNote(ctx, id, title, content)
	if err != nil {
		if errors.Is(err, storage.IdNotFound) {
			log.Warn("Id not found", slog.String("err", err.Error()))
		}
		return note, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("Note updated", slog.Any("note", note))
	return note, nil
}
func (n *Notes) DeleteNote(ctx context.Context, id string) (models.Note, error) {
	const op = "services.notes.DeleteNote"
	log := n.log.With(slog.String("op", op))
	note, err := n.noteDeleter.DeleteNote(ctx, id)
	if err != nil {
		if errors.Is(err, storage.IdNotFound) {
			log.Warn("Id not found", slog.String("err", err.Error()))
		}
		return note, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("Note deleted ", slog.Any("note", note))
	return note, nil
}
func (n *Notes) GetNotes(
	ctx context.Context,
	limit int32,
	offset_id string) (notes []models.Note, next_offset_id string, err error) {
	const op = "services.notes.GetNotes"
	log := n.log.With(slog.String("op", op))
	notes, next_offset_id, err = n.noteLister.GetNotes(ctx, limit, offset_id)
	if err != nil {
		if errors.Is(err, storage.IdNotFound) {
			log.Warn("Id not found", slog.String("err", err.Error()))
		}
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("Notes received ", slog.Any("notes", notes))
	return notes, next_offset_id, nil
}
