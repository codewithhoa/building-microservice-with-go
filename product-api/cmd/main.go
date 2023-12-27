package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codewithhoa/building-microservice-with-go/product-api/data"
	"github.com/codewithhoa/building-microservice-with-go/product-api/internal/handlers"
	"github.com/codewithhoa/building-microservice-with-go/product-api/pkg/config"
	"github.com/codewithhoa/building-microservice-with-go/product-api/pkg/logger"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Load config environment
	cfg := config.NewConfig()
	err := cfg.LoadConfig()
	if err != nil {
		log.Fatal("can not get config, err:", err)
	}

	// Init logger
	slh, err := logger.NewSlogHandler(cfg)
	if err != nil {
		log.Fatal(err)
	}

	sll := logger.NewSlogLogger(slog.New(slh))
	v := data.NewValidation()
	productHandler := handlers.NewProductsHandler(sll, v)
	r := mux.NewRouter()
	// middleware to set "Content-Type" to "application/json"
	r.Use(productHandler.MiddlewareContentTypeJSON)

	api := r.PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	v1 := api.PathPrefix("/v1").Subrouter()
	getProducts := v1.Methods(http.MethodGet).Subrouter()
	getProducts.HandleFunc("/products", productHandler.GetAll)

	getProductByID := v1.Methods(http.MethodGet).Subrouter()
	getProductByID.HandleFunc("/products/{id:[0-9]+}", productHandler.GetByID)

	postProducts := v1.Methods(http.MethodPost).Subrouter()
	postProducts.HandleFunc("/products", productHandler.Post)
	postProducts.Use(productHandler.MiddlewareProductValidation)

	putProducts := v1.Methods(http.MethodPut).Subrouter()
	putProducts.HandleFunc("/products/{id:[0-9]+}", productHandler.Put)
	putProducts.Use(productHandler.MiddlewareProductValidation)

	deleteProducts := v1.Methods(http.MethodDelete).Subrouter()
	deleteProducts.HandleFunc("/products/{id:[0-9]+}", productHandler.Delete)

	// CORS configuration
	ch := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{
			"*", // accept all host
			// "http://localhost:9999", // host for documentation
		}),
		gohandlers.AllowedHeaders(
			[]string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Accept", "Authorization"},
		),
		gohandlers.AllowedMethods([]string{"GET", "PATCH", "POST", "PUT", "OPTIONS", "DELETE"}),
		gohandlers.MaxAge(1),
	)

	// Config for server
	srv := &http.Server{
		Addr:     ":9090",
		ErrorLog: slog.NewLogLogger(slh, cfg.SlogLogLevel()),
		// Good practice to set timeouts to avoid Slowloris attacks.
		Handler:      ch(r),
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		sll.Info(fmt.Sprintf("Starting server on port %s", cfg.ServerAddress()))
		if err := srv.ListenAndServe(); err != nil {
			sll.Error(err.Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	sig := <-sigChan

	sll.Info("received signal, starting graceful shutdown", "sig", sig)

	ctx, ctxCancel := context.WithTimeout(context.Background(), cfg.ServerGracefulTimeout())
	defer ctxCancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		sll.Error("error from shutdown")
		os.Exit(0)
	}

	sll.Info("shutdown successfully")
	os.Exit(0)
}
