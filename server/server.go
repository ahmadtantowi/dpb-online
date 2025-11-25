package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	srv    *http.Server
	logger *slog.Logger
}

func NewServer(logger *slog.Logger, addr string, opt ...Option) *Server {
	server := &Server{
		srv: &http.Server{
			Addr:    addr,
			Handler: nil,
		},
		logger: logger,
	}

	for _, o := range opt {
		o(server)
	}

	return server
}

type Option func(*Server)

func WithRouter(handler http.Handler) Option {
	return func(s *Server) {
		s.srv.Handler = handler
	}
}

func (s *Server) StartAndWait() {
	go func() {
		s.logger.Info("starting server", "port", s.srv.Addr)
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Warn("failed to start server", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.logger.Info("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.Warn("server shutdown", "error", err)
	}
	s.logger.Info("server exiting")
	os.Exit(0)
}
