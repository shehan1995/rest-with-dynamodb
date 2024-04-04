package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-with-dynamodb/config"
	"rest-with-dynamodb/repositories/dynamodb"
	"syscall"
	"time"
)

type Server struct {
	router   *chi.Mux
	config   config.ServerConfig
	dynamoDB *dynamodb.DynamoHandle
}

func NewServer(cfg config.Config, dynamodb *dynamodb.DynamoHandle) *Server {
	srv := &Server{
		router:   chi.NewRouter(),
		config:   cfg.Server,
		dynamoDB: dynamodb,
	}

	srv.httpRouter()

	return srv
}

func (s *Server) Start(ctx context.Context) {
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", s.config.Port),
		Handler:      s.router,
		IdleTimeout:  time.Duration(s.config.IdleTimeout) * time.Second,
		ReadTimeout:  time.Duration(s.config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.WriteTimeout) * time.Second,
	}

	shutdownComplete := handleShutdown(func() {
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("server.Shutdown failed: %v\n", err)
		}
	})

	initRepositories(s.dynamoDB)

	log.Printf("Starting Server on port: %v\n", s.config.Port)

	if err := server.ListenAndServe(); err == http.ErrServerClosed {
		<-shutdownComplete
	} else {
		log.Printf("http.ListenAndServe failed: %v\n", err)
	}

	log.Println("Shutdown gracefully")

}

func handleShutdown(onShutdownSignal func()) <-chan struct{} {
	shutdown := make(chan struct{})

	go func() {
		shutdownSignal := make(chan os.Signal, 1)
		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

		<-shutdownSignal

		onShutdownSignal()
		close(shutdown)
	}()

	return shutdown
}

func initRepositories(dynamoDBHandler *dynamodb.DynamoHandle) {
	dynamodb.DynamoDBRepo = dynamodb.NewRepository(dynamoDBHandler, "books")
}
