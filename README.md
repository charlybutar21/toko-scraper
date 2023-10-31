# toko-scraper
This is a simple command-line tool that is used for scraping a number of [products of category Mobile phones/Handphone from Tokopedia](https://www.tokopedia.com/p/handphone-tablet/handphone?ob=5&page=) and stores it in a csv file and in a PostgreSQL
database.

## Technology
* Golang 1.21.1
* PostgresSQL
* Docker
* [Go-Colly](https://go-colly.org/docs/examples/basic/)

## How to Setup & Run Application
**1. Clone repository**
```bash
git clone https://github.com/charlybutar21/toko-scraper.git && cd toko-scraper
```
**2. Run docker compose**
```bash
docker compose up
```
**3. Open another terminal and run command line tool (populate-products)**
```bash
go run main.go populate-products
```
**4. To see the result, open file product.csv or database console by following these steps:**
- Load environment variables
```bash
source .env
```
- Access the PostgreSQL database console
```bash
psql "host=localhost port=5433 user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=disable"
```
- View list of tables
```bash
\dt 
```
- Switch to tabular output
```bash
\x
```
- Query and view the products table
```bash
select * from products;
```

## Architecture

Command -> Service -> Repository