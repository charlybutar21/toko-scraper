package service

import (
	"github.com/gocolly/colly"
	"gitlab.com/pinvest/toko-scraper/app"
	"gitlab.com/pinvest/toko-scraper/entity"
	"strconv"
	"sync"
)

type ScraperServiceImpl struct {
	colly         *colly.Collector
	configuration app.Config
}

func NewScraperService(configuration app.Config) ScraperService {
	return &ScraperServiceImpl{
		colly:         colly.NewCollector(),
		configuration: configuration,
	}
}

func (service *ScraperServiceImpl) ScrapeTokopediaProducts() ([]entity.Product, error) {
	totalProductsScraped := service.configuration.Get("TOTAL_PRODUCTS_SCRAPED")
	totalRecord, err := strconv.Atoi(totalProductsScraped)
	if err != nil {
		return nil, err
	}

	service.colly.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:57.0) Gecko/20100101 Firefox/57.0")
	})

	var products []entity.Product
	productCh := make(chan entity.Product)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			service.colly.OnHTML(".e1nlzfl9", func(e *colly.HTMLElement) {
				productName := e.ChildText("div.css-11s9vse > span.css-20kt3o")
				description := e.ChildText("div.css-11s9vse > span.css-20kt3o")
				imageLink := e.ChildAttr("div.css-79elbk img", "src") // base64 because generated from html
				price := e.ChildText("div.css-pp6b3e > span.css-o5uqvq")

				var rating int
				e.ForEach(".css-177n1u3", func(index int, j *colly.HTMLElement) {
					rating++
				})

				var merchantName string
				counterIndexMerchant := 1
				e.ForEach("div.css-vbihp9 > span.css-ywdpwd", func(index int, k *colly.HTMLElement) {
					if counterIndexMerchant%2 == 0 {
						merchantName = k.Text
					}
					counterIndexMerchant++
				})

				productCh <- entity.Product{
					Name:        productName,
					Description: description,
					ImageLink:   imageLink,
					Price:       price,
					Rating:      strconv.Itoa(rating),
					Merchant:    merchantName,
				}

			})

			counterPage := 1
			for len(products) < totalRecord {
				service.colly.Visit("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=5&page=" + strconv.Itoa(counterPage))
				counterPage++
			}
		}()
	}

	go func() {
		wg.Wait()
		close(productCh)
	}()

	for product := range productCh {
		products = append(products, product)
		if len(products) >= totalRecord {
			break
		}
	}

	return products, nil
}
