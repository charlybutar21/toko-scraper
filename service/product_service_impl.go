package service

import (
	"context"
	"database/sql"
	"github.com/charlybutar21/toko-scraper/entity"
	"github.com/charlybutar21/toko-scraper/repository"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
	database          *sql.DB
}

func NewProductService(productRepository repository.ProductRepository, database *sql.DB) ProductService {
	return &ProductServiceImpl{
		productRepository: productRepository,
		database:          database,
	}
}

func (service *ProductServiceImpl) InsertProducts(ctx context.Context, products []entity.Product) error {
	tx, err := service.database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, product := range products {
		err = service.productRepository.Insert(ctx, tx, product)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
