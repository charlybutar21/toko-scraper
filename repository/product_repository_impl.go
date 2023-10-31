package repository

import (
	"context"
	"database/sql"
	"gitlab.com/pinvest/toko-scraper/entity"
)

type productRepositoryImpl struct {
}

func NewProductRepository(database *sql.DB) ProductRepository {
	return &productRepositoryImpl{}
}

func (repository *productRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, product entity.Product) error {
	query := "INSERT INTO products (name, description, image_link, price, rating, merchant) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := tx.ExecContext(ctx, query, product.Name, product.Description, product.ImageLink, product.Price, product.Rating, product.Merchant)
	if err != nil {
		return err
	}

	return nil
}
