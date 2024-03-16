package main

import (
	"github.com/crewblade/notes_service/internal/app"
	"github.com/crewblade/notes_service/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	cfg := config.MustLoad()
	log.Info("starting application",
		slog.Any("cfg", cfg))
	application := app.New(
		log,
		cfg.GRPC.Port,
		cfg.ConnectionString,
	)
	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	signal := <-stop
	log.Info("application stopped with signal:" + signal.String())

	if err := application.CloseDB(); err != nil {
		log.Error("failed to close database connection", slog.String("err", err.Error()))
	}
	application.GRPCSrv.Stop()

}
