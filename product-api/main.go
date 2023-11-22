package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codewithhoa/building-microservice-with-go/product-api/handlers"
	"github.com/codewithhoa/building-microservice-with-go/product-api/pkg/config"
	"github.com/codewithhoa/building-microservice-with-go/product-api/pkg/logger"
	"github.com/gorilla/mux"
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

	productHandler := handlers.NewProductsHandler(logger)

	r := mux.NewRouter()

	getProducts := r.Methods(http.MethodGet).Subrouter()
	getProducts.HandleFunc("/products", productHandler.GetProducts)

	postProducts := r.Methods(http.MethodPost).Subrouter()
	postProducts.HandleFunc("/products", productHandler.Create)
	postProducts.Use(productHandler.MiddlewareProductValidation)

	putProducts := r.Methods(http.MethodPut).Subrouter()
	putProducts.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProduct)
	putProducts.Use(productHandler.MiddlewareProductValidation)

	// Config for server
	srv := &http.Server{
		Addr: ":9090",
		// Good practice to set timeouts to avoid Slowloris attacks.
		Handler:      r,
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		logger.Info(fmt.Sprintf("Starting server on port %s", cf.ServerAddress))
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(err.Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	sig := <-sigChan

	logger.Info("received signal, starting graceful shutdown", "sig", sig)

	ctx, ctxCancel := context.WithTimeout(context.Background(), cf.ServerGracefulTimeout)
	defer ctxCancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		logger.Error("error from shutdown")
		os.Exit(0)
	}

	logger.Info("shutdown successfully")
	os.Exit(0)
}
