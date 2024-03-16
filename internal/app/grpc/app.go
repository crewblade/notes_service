package grpcapp

import (
	"fmt"
	notesrpc "github.com/crewblade/notes_service/internal/grpc/notes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, notesService notesrpc.Notes, port int) *App {
	gRPCServer := grpc.NewServer()
	reflection.Register(gRPCServer)
	notesrpc.Register(gRPCServer, notesService)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}
func (a *App) Run() error {
	const op = "grpcapp.Run"
	log := a.log.With(slog.String("op", op))
	log.Info("Starting gRPC server")
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", a.port))

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"
	log := a.log.With(slog.String("op", op))
	log.Info("Stopping gRPC server", slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()

}
