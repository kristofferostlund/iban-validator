package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	router     *http.ServeMux
	httpServer *http.Server
}

func New(address string, port int) *Server {
	addr := fmt.Sprintf("%s:%d", address, port)

	return &Server{
		router:     http.NewServeMux(),
		httpServer: &http.Server{Addr: addr},
	}
}

func (s *Server) Serve() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGTERM)

	go s.startServer()

	<-stop

	s.close()
}

func (s *Server) startServer() {
	s.initRoutes()
	s.httpServer.Handler = s.router

	log.Printf("Server running on %s", s.httpServer.Addr)

	if err := s.httpServer.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatalf("Error running serer")
		}
	}
}

func (s *Server) close() {
	log.Println("Shutting down server...")

	defer log.Println("Server successfully shut down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shut down server: %v", err)
	}

	os.Exit(0)
}
