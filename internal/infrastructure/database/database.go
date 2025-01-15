package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/MLaskun/ovidish/internal/product/config"
)

func Init(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Database.Dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
