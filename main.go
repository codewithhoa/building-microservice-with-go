package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codewithhoa/building-microservice-with-go/handlers"
)

func main() {
	slogOpts := slog.HandlerOptions{
		AddSource: false,
		Level:     slog.Level(slog.LevelInfo),
	}

	var slogHandler slog.Handler = slog.NewTextHandler(os.Stdout, &slogOpts)

	slogHandler = slogHandler.WithAttrs([]slog.Attr{
		{
			Key:   "app-name",
			Value: slog.AnyValue("todo-api"),
		},
		{
			Key:   "app-version",
			Value: slog.AnyValue("v0.0.1"),
		},
	})

	logger := slog.New(slogHandler)

	rootHandler := handlers.NewRootHandler(logger)
	helloHandler := handlers.NewHelloHandler(logger)
	goodbyeHandler := handlers.NewGoodbyeHandler(logger)

	sm := http.NewServeMux()

	sm.Handle("/", rootHandler)
	sm.Handle("/hello", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		logger.Info("Starting server on port 9090")
		if err := s.ListenAndServe(); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan

	logger.Info("received signal, starting graceful shutdown", "sig", sig)

	ctx, ctxCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer ctxCancel()

	s.Shutdown(ctx)
}
