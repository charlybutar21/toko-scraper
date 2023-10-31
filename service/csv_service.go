package service

import "gitlab.com/pinvest/toko-scraper/entity"

type CSVService interface {
	WriteProductCSV(products []entity.Product) error
}
