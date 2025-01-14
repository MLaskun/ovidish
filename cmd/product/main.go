package main

import (
	"github.com/MLaskun/ovidish/internal/product"
	"github.com/MLaskun/ovidish/internal/product/config"
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
	cfg := &config.Config{
		Address: ":3000",
	}

	return cfg, nil
}
