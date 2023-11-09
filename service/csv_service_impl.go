package service

import (
	"encoding/csv"
	"github.com/charlybutar21/toko-scraper/entity"
	"os"
)

type CSVServiceImpl struct {
}

func NewCSVService() CSVService {
	return &CSVServiceImpl{}
}

func (service *CSVServiceImpl) WriteProductCSV(products []entity.Product) error {
	file, err := os.Create("product.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Product Name", "Description", "Image Link", "Price", "Rating", "Merchant Name"}
	writer.Write(headers)

	for _, product := range products {
		record := []string{
			product.Name,
			product.Description,
			product.ImageLink,
			product.Price,
			product.Rating,
			product.Merchant,
		}

		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil

}
