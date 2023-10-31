package main

import (
	"github.com/spf13/cobra"
	"gitlab.com/pinvest/toko-scraper/app"
	"gitlab.com/pinvest/toko-scraper/command"
	"gitlab.com/pinvest/toko-scraper/repository"
	"gitlab.com/pinvest/toko-scraper/service"
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
