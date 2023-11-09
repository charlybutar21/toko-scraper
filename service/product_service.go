package service

import (
	"context"
	"github.com/charlybutar21/toko-scraper/entity"
)

type ProductService interface {
	InsertProducts(ctx context.Context, products []entity.Product) error
}
