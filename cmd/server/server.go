package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/huichiaotsou/go-roster/apihandler"
	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/middleware"
	"github.com/huichiaotsou/go-roster/model"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	router *mux.Router
	srv    *http.Server
	logger *log.Logger
}

func NewServer() *Server {
	logger := log.New()
	// logger.SetFormatter(&log.JSONFormatter{})

	s := &Server{
		router: mux.NewRouter(),
		logger: logger,
	}

	// Init Sqlx
	sqlx, err := model.InitSqlx(config.GetDBConfig())
	if err != nil {
		logger.Fatalf("Failed to initialize database: %v", err)
	}

	// Init Database instance
	db := model.NewDatabase(sqlx)

	// Init middleware
	middleware := middleware.New(db, logger)

	// Init api handler
	apiHandler := apihandler.New(s.router, logger, middleware, db)

	// Register all routes
	apiHandler.RegisterAllRoutes()

	// Set up CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(s.router)

	s.srv = &http.Server{
		Addr:         fmt.Sprintf(":%s", config.GetServerPort()),
		Handler:      corsHandler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s
}

func (s *Server) Start() {
	go func() {
		s.logger.Infof("Starting server on port %s", config.GetServerPort())
		if err := s.srv.ListenAndServe(); err != nil {
			s.logger.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.Errorf("Error shutting down server: %v", err)
		os.Exit(1)
	}

	s.logger.Info("Server gracefully stopped.")
}

func (s *Server) UseMiddleware(middleware func(http.Handler) http.Handler) {
	s.router.Use(middleware)
}

func (s *Server) HandleFunc(method, path string, handlerFunc http.HandlerFunc) {
	s.router.HandleFunc(path, handlerFunc).Methods(method)
}
