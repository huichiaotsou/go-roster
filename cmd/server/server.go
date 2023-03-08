package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/cmd/api"
	"github.com/huichiaotsou/go-roster/config"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	router *mux.Router
	srv    *http.Server
}

func NewServer() *Server {
	s := &Server{}
	s.router = mux.NewRouter()

	// Register all routes
	api.RegisterAllRoutes(s.router)

	// Set up CORS
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(s.router)

	addr := fmt.Sprintf(":%s", config.GetServerPort())

	s.srv = &http.Server{
		Addr:         addr,
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s
}

func (s *Server) Start() {
	go func() {
		log.Infof("Starting server on port %s", config.GetServerPort())
		if err := s.srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Errorf("Error shutting down server: %v", err)
		os.Exit(1)
	}

	log.Info("Server gracefully stopped.")
}

func (s *Server) UseMiddleware(middleware func(http.Handler) http.Handler) {
	s.router.Use(middleware)
}

func (s *Server) HandleFunc(method, path string, handlerFunc http.HandlerFunc) {
	s.router.HandleFunc(path, handlerFunc).Methods(method)
}
