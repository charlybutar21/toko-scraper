FROM golang:1.21.4

# set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download and install the dependencies
RUN go mod download

# Build Go app
RUN go build -o toko-scraper .

# Jalankan perintah saat container dimulai
CMD ["./toko-scraper"]