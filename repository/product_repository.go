package repository

import (
	"context"
	"database/sql"
	"github.com/charlybutar21/toko-scraper/entity"
)

type ProductRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, product entity.Product) error
}
