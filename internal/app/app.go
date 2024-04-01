package app

import (
	grpcapp "github.com/crewblade/notes_service/internal/app/grpc"
	"github.com/crewblade/notes_service/internal/services/notes"
	"github.com/crewblade/notes_service/internal/storage/postgres"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.App
	Storage *postgres.Storage
}

func New(log *slog.Logger, grpcPort int, connectionString string) *App {
	storage, err := postgres.New(connectionString)
	if err != nil {
		panic(err)
	}

	notesService := notes.New(log, storage, storage, storage, storage, storage)
	grpcApp := grpcapp.New(log, notesService, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
		Storage: storage,
	}
}
