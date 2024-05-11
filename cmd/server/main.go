package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/sportgroup-hq/api/internal/bootstrap"
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/common-lib/logger"
)

func main() {
	deps, err := bootstrap.Up()
	if err != nil {
		panic(err)
	}

	slog.SetDefault(slog.New(logger.NewLogger(os.Stdout, config.Get().Log.Level)))

	slog.Info("Starting up...")

	go func() {
		if err := deps.GRPCServer.Start(); err != nil {
			panic(fmt.Errorf("failed to start grpc server: %w", err))
		}
	}()

	if err = deps.HTTPServer.Start(); err != nil {
		panic(err)
	}
}
