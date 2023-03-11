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
	db     *model.Database
}

func NewServer() *Server {
	router := mux.NewRouter()

	// Init Sqlx
	sqlx, err := model.InitSqlx(config.GetDBConfig())
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Init DB instance
	db := model.NewDatabase(sqlx)

	// Init middleware
	middleware := middleware.New(db)

	// Init api handler
	apiHandler := apihandler.New(router, middleware, db)

	// Register all routes
	apiHandler.RegisterAllRoutes()

	// Set up CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(router)

	// Config server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.GetServerPort()),
		Handler:      corsHandler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{
		router: router,
		srv:    srv,
		db:     db,
	}
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
	defer s.db.Close()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Errorf("Error shutting down server: %v", err)
		os.Exit(1)
	}

	log.Info("Server gracefully stopped.")
}
