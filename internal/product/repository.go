package product

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/MLaskun/ovidish/internal/product/model"
	"github.com/lib/pq"
)

var (
	ErrNoRecordFound = errors.New("record not found")
	ErrEditConflict  = errors.New("edit conflict")
)

type ProductRepository struct {
	DB *sql.DB
}

func (pr ProductRepository) Insert(product *model.Product) error {
	query := `
        INSERT INTO products (name, description, categories, quantity, price)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, version`

	args := []any{product.Name, product.Description,
		pq.Array(product.Categories), product.Quantity, product.Price}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return pr.DB.QueryRowContext(ctx, query, args...).
		Scan(&product.ID, &product.Version)
}

func (pr ProductRepository) Get(id int64) (*model.Product, error) {
	if id < 1 {
		return nil, ErrNoRecordFound
	}

	query := `
        SELECT id, name, description, quantity, price, version
        FROM products
        WHERE id = $1`

	var product model.Product

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := pr.DB.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		pq.Array(&product.Categories),
		&product.Quantity,
		&product.Price,
		&product.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoRecordFound
		default:
			return nil, err
		}
	}

	return &product, nil
}

func (pr ProductRepository) Update(product *model.Product) error {
	query := `
        UPDATE products
        SET name = $1, description = $2, categories = $3,
        quantity = $4, price = $5, version = version + 1
        WHERE id = $6 AND version = $7`

	args := []any{
		product.Name,
		product.Description,
		pq.Array(product.Categories),
		product.Quantity,
		product.Price,
		product.ID,
		product.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := pr.DB.QueryRowContext(ctx, query, args...).
		Scan(&product.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (pr ProductRepository) Delete(id int64) error {
	if id < 1 {
		return ErrNoRecordFound
	}

	query := `
        DELETE FROM products
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := pr.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNoRecordFound
	}

	return nil
}

func (pr ProductRepository) GetAll(name string,
	categories []string) ([]*model.Product, error) {
	query := `
        SELECT id, name, description, categories, quantity, price, version
        FROM products
        WHERE (name ILIKE '%' || $1 || '%' OR $1 = '')
        AND (categories @> $2 OR $2 = '{}'
        ORDER BY id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{name, pq.Array(categories)}

	rows, err := pr.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	totalRecords := 0
	products := []*model.Product{}

	for rows.Next() {
		var product model.Product
		err := rows.Scan(
			&totalRecords,
			&product.ID,
			&product.Name,
			&product.Description,
			pq.Array(&product.Categories),
			&product.Quantity,
			&product.Price,
			&product.Version,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil

}
