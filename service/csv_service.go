package service

import "github.com/charlybutar21/toko-scraper/entity"

type CSVService interface {
	WriteProductCSV(products []entity.Product) error
}
