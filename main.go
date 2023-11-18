package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codewithhoa/building-microservice-with-go/handlers"
	"github.com/codewithhoa/building-microservice-with-go/pkg/config"
	"github.com/codewithhoa/building-microservice-with-go/pkg/logger"
)

func main() {

	// Load config environment
	cf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("can not get config")
		return
	}

	// Init logger
	logger, err := logger.NewLogger(cf)
	if err != nil {
		log.Fatal("can not run logger")
		return
	}

	rootHandler := handlers.NewRootHandler(logger)
	helloHandler := handlers.NewHelloHandler(logger)
	goodbyeHandler := handlers.NewGoodbyeHandler(logger)
	productHandler := handlers.NewProductsHandler(logger)

	sm := http.NewServeMux()

	sm.Handle("/", rootHandler)
	sm.Handle("/hello", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)
	sm.Handle("/products/", productHandler)

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

	err = s.Shutdown(ctx)
	if err != nil {
		panic("error from shutdown")
	}
}
