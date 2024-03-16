package notes

import (
	"context"
	"github.com/crewblade/notes_service/internal/domain/models"
	pb "github.com/crewblade/notes_service/protos/gen/go/notes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Notes interface {
	CreateNote(ctx context.Context, title string, content string) (id string, err error)
	GetNoteById(ctx context.Context, id string) (note models.Note, err error)
	GetNotes(ctx context.Context, limit int32, offset_id string) (notes []models.Note, next_offset_id string, err error)
	UpdateNote(ctx context.Context, id string, title string, content string) (note models.Note, err error)
	DeleteNote(ctx context.Context, id string) (note models.Note, err error)
}

type serverAPI struct {
	pb.UnimplementedNotesServer
	notes Notes
}

func Register(gRPC *grpc.Server, notes Notes) {
	pb.RegisterNotesServer(gRPC, &serverAPI{notes: notes})
}
func (s *serverAPI) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	if req.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "title is required")
	}
	id, err := s.notes.CreateNote(ctx, req.GetTitle(), req.GetContent())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &pb.CreateNoteResponse{
		Id: id,
	}, nil

}
func (s *serverAPI) GetNoteById(ctx context.Context, req *pb.GetNoteByIdRequest) (*pb.Note, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	var note models.Note
	note, err := s.notes.GetNoteById(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.Note{
		Id:      note.Id,
		Title:   note.Title,
		Content: note.Content,
	}, nil
}
func (s *serverAPI) GetNotes(ctx context.Context, req *pb.GetNotesRequest) (*pb.GetNotesResponse, error) {
	if req.GetOffsetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "get_offset_id is required")
	}
	var notesData []models.Note
	notesData, next_offset_id, err := s.notes.GetNotes(ctx, req.GetLimit(), req.GetOffsetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	var notes []*pb.Note
	for _, note := range notesData {
		notes = append(notes, &pb.Note{
			Id:      note.Id,
			Title:   note.Title,
			Content: note.Content,
		})
	}
	return &pb.GetNotesResponse{
		Notes:        notes,
		NextOffsetId: next_offset_id,
	}, nil
}

func (s *serverAPI) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*pb.Note, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	var note models.Note
	note, err := s.notes.UpdateNote(ctx, req.GetId(), req.GetTitle(), req.GetContent())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.Note{
		Id:      note.Id,
		Title:   note.Title,
		Content: note.Content,
	}, nil

}
func (s *serverAPI) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.Note, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	var note models.Note
	note, err := s.notes.DeleteNote(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.Note{
		Id:      note.Id,
		Title:   note.Title,
		Content: note.Content,
	}, nil
}
