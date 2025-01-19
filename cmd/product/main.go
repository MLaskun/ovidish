package main

import (
	"fmt"
	"os"

	"github.com/MLaskun/ovidish/internal/product"
	"github.com/MLaskun/ovidish/internal/product/config"
)

const (
	databaseDsnEnv = "DATABASE_DSN"
	addressEnv     = "PRODUCT_ADDRESS"
)

func main() {
	cfg, err := makeConfig()
	if err != nil {
		panic(err)
	}

	if err := product.NewServer(cfg).Run(); err != nil {
		panic(err)
	}
}

func makeConfig() (*config.Config, error) {
	cfg := &config.Config{}

	if err := unmarshalConfigFromEnv(cfg); err != nil {
		return nil, fmt.Errorf("config creation failed: %w", err)
	}

	return cfg, nil
}

func unmarshalConfigFromEnv(cfg *config.Config) error {
	str := os.Getenv(databaseDsnEnv)
	if len(str) == 0 {
		return fmt.Errorf("missing variable: %s", databaseDsnEnv)
	}
	cfg.Database.Dsn = str

	str = os.Getenv(addressEnv)
	if len(str) == 0 {
		return fmt.Errorf("missing variable: %s", addressEnv)
	}
	cfg.Address = str

	return nil
}
