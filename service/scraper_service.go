package service

import "github.com/charlybutar21/toko-scraper/entity"

type ScraperService interface {
	ScrapeTokopediaProducts() ([]entity.Product, error)
}
