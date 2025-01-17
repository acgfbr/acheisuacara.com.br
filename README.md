# Achei Sua Cara - URL Shortener

A specialized URL shortener service that only accepts and processes marketplace URLs. Built with Go, MySQL, and Redis.

Production live: acheisuacara.com.br

## Features

- URL shortening for marketplace links only
- Rate limiting (60 requests per minute)
- Click tracking
- Input validation
- Duplicate URL detection
- Permanent redirects

## Supported Marketplaces

- Amazon
- MercadoLivre
- Americanas
- Magazine Luiza
- Shopee
- AliExpress

## Prerequisites

- Go 1.16 or higher
- MySQL 8 or higher ( i am using an external mysql xd )
- Redis 6.0 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/acheisuacara.com.br.git
cd acheisuacara.com.br
```

2. Install dependencies:
```bash
go mod download
```

3. Configure the application:
- Copy `.env.example` to the root directory and fill with your values
- Update the configuration values as needed

4. Create the MySQL database:
```sql
CREATE DATABASE acheisuacara;
```

5. Run the application:
```bash
go run cmd/api/main.go
```

## Testing

The project uses SQLite for testing to avoid requiring a MySQL database during tests. To run the tests:

1. Run all tests:
```bash
go test ./... -v
```

2. Run tests for specific packages:
```bash
# Test models
go test ./pkg/models -v

# Test services
go test ./pkg/services -v

# Test handlers
go test ./pkg/handlers -v
```

3. Run tests with coverage:
```bash
go test ./... -v -cover
```

4. Run tests and generate coverage report:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Note: The "record not found" messages during tests are expected and part of the test cases.

## API Endpoints

### Shorten URL
```
POST /api/shorten
Content-Type: application/json

{
    "url": "https://www.amazon.com/product/123"
}
```

### Access Shortened URL
```
GET /:shortCode
```

## Rate Limiting

The API is rate-limited to 60 requests per minute per IP address. If you exceed this limit, you'll receive a 429 Too Many Requests response.

## Error Handling

- 400 Bad Request: Invalid URL or non-marketplace URL
- 404 Not Found: Short URL not found
- 429 Too Many Requests: Rate limit exceeded
- 500 Internal Server Error: Server-side error

## License

MIT License 