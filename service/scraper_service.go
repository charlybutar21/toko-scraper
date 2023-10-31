package service

import "gitlab.com/pinvest/toko-scraper/entity"

type ScraperService interface {
	ScrapeTokopediaProducts() ([]entity.Product, error)
}
