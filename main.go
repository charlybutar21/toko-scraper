package main

import (
	"github.com/charlybutar21/toko-scraper/app"
	"github.com/charlybutar21/toko-scraper/command"
	"github.com/charlybutar21/toko-scraper/repository"
	"github.com/charlybutar21/toko-scraper/service"
	"github.com/spf13/cobra"

	"os"
)

func main() {

	config := app.New()
	database := app.NewPostgresqlDatabase(config)

	productRepository := repository.NewProductRepository(database)
	scraperService := service.NewScraperService(config)
	productService := service.NewProductService(productRepository, database)
	csvService := service.NewCSVService()

	rootCmd := &cobra.Command{
		Use:   "command-cli",
		Short: "Example CLI command in Go using cobra.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	rootCmd.AddCommand(
		command.PopulateProducts(scraperService, productService, csvService),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
