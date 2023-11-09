package command

import (
	"context"
	"fmt"
	"github.com/charlybutar21/toko-scraper/service"
	"github.com/spf13/cobra"
)

func PopulateProducts(scraperService service.ScraperService, productService service.ProductService, csvService service.CSVService) *cobra.Command {
	return &cobra.Command{
		Use:   "populate-products",
		Short: "Populate products",
		Long:  "Populate products",
		RunE: func(cmd *cobra.Command, args []string) error {
			return populateProducts(scraperService, productService, csvService)
		},
	}
}

func populateProducts(scraperService service.ScraperService, productService service.ProductService, csvService service.CSVService) error {
	products, err := scraperService.ScrapeTokopediaProducts()
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = productService.InsertProducts(ctx, products)
	if err != nil {
		return err
	}

	err = csvService.WriteProductCSV(products)
	if err != nil {
		return err
	}

	fmt.Println("Populate products has been successfully completed!")

	return nil
}
