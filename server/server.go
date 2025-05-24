package server

import (
	"context"
	"fmt"
	"gotth/assets"
	"gotth/modules/landing"
	"gotth/store"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	logger *log.Logger
	port int
	httpServer *http.Server
	store *store.BBoltDatabase

	useEmbedded bool
}

func MustNewServer(
	logger *log.Logger,
	port int,
	db *store.BBoltDatabase,
	useEmbedded bool,
) *Server {
	return &Server{
		logger: logger,
		port: port,
		store: db,
		useEmbedded: useEmbedded,
	}
}

func (s *Server) Start() error {
	s.logger.Printf("Starting server on port %d", s.port)

	// If useEmbedded, we will serve the static files from the embedded filesystem, else we will serve from the filesystem ./assets
	router := http.NewServeMux()

	if s.useEmbedded {
		s.logger.Println("Using embedded filesystem")
		router.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assets.Assets))))
	} else {
		s.logger.Println("Using filesystem")
		router.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	}

	landing.NewLandingController(router)

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: router,
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func () {
		if err := s.httpServer.ListenAndServe()
		err != http.ErrServerClosed {
			log.Fatalf("Error running server: %v", err)
		}
	}()

	<- stopChan

	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		s.logger.Fatalf("Error shutting down server: %v", err)
		return err
	}
	return nil
}
