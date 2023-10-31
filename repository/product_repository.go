package repository

import (
	"context"
	"database/sql"
	"gitlab.com/pinvest/toko-scraper/entity"
)

type ProductRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, product entity.Product) error
}
