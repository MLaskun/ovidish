package product

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MLaskun/ovidish/internal/product/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config) *Server {
	repo := NewProductRepository()
	svc := NewProductService(repo)
	handler := NewProductHandler(svc)
	routes := routes(handler)

	httpServer := &http.Server{
		Addr:         cfg.Address,
		Handler:      routes,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	shutdownError := make(chan error, 1)

	go func() {
		sig := <-quit
		fmt.Println("Shutting down server gracefully...",
			"signal", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := s.httpServer.Shutdown(ctx); err != nil {
			shutdownError <- fmt.Errorf("server shutdown error: %w", err)
			return
		}

		fmt.Println("server gracefully stopped")
		shutdownError <- nil
	}()

	fmt.Println("starting server on port", s.httpServer.Addr)
	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}

	if err := <-shutdownError; err != nil {
		return err
	}

	return nil
}
