package product

import (
	"net/http"

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
	srv := &http.Server{
		Addr:    s.cfg.Address,
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
