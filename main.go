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

	// Add attribute that should be included in all the logs being generated.
	slogHandler = slogHandler.WithAttrs([]slog.Attr{
		slog.String("app-name", "product-api"),
		slog.String("app-version", "v0.0.1"),
	})

	logger := slog.New(slogHandler)

	rootHandler := handlers.NewRootHandler(logger)
	helloHandler := handlers.NewHelloHandler(logger)
	goodbyeHandler := handlers.NewGoodbyeHandler(logger)
	productHandler := handlers.NewProductsHandler(logger)

	sm := http.NewServeMux()

	sm.Handle("/", rootHandler)
	sm.Handle("/hello", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)
	sm.Handle("/products", productHandler)

	// Config for server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		logger.Info("Starting server on port 9090")
		if err := s.ListenAndServe(); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()

	// Implement gracefully shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan

	logger.Info("received signal, starting graceful shutdown", "sig", sig)

	ctx, ctxCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer ctxCancel()

	s.Shutdown(ctx)
}
