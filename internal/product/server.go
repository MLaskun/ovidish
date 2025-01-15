package product

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MLaskun/ovidish/internal/infrastructure/database"
	"github.com/MLaskun/ovidish/internal/product/config"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() error {
	db, err := database.Init(*s.cfg)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	fmt.Println("database connection pool established")

	repo := NewProductRepository(s.cfg, db)
	svc := NewProductService(repo)
	handler := NewProductHandler(svc)
	routes := routes(handler)

	httpServer := &http.Server{
		Addr:         s.cfg.Address,
		Handler:      routes,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	shutdownError := make(chan error, 1)

	go func() {
		sig := <-quit
		fmt.Println("Shutting down server gracefully...",
			"signal", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			shutdownError <- fmt.Errorf("server shutdown error: %w", err)
			return
		}

		fmt.Println("server gracefully stopped")
		shutdownError <- nil
	}()

	fmt.Println("starting server on port", httpServer.Addr)
	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}

	if err := <-shutdownError; err != nil {
		return err
	}

	return nil
}
