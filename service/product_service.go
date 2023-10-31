package service

import (
	"context"
	"gitlab.com/pinvest/toko-scraper/entity"
)

type ProductService interface {
	InsertProducts(ctx context.Context, products []entity.Product) error
}
