# GoCRUD - Product Management API

## Overview

GoCRUD is a RESTful API built with Go for managing products. It demonstrates CRUD (Create, Read, Update, Delete) operations with a **Clean Architecture** using the repository pattern, service layer and data transfer objects (DTOs).

## Technology Stack

- **Language**: Go 1.25.6
- **Web Framework**: [Chi Router](https://github.com/go-chi/chi) - Lightweight HTTP routing
- **Database**: SQLite with [GORM](https://gorm.io/) - Golang ORM
- **Additional Libraries**:
  - `google/uuid` - UUID generation
  - `go-sqlite3` - SQLite driver

## Project Structure

```
gocrud/
├── main.go # Application entry point
├── go.mod # Go module definition
├── README.md # Documentation
├── core/ # Core/shared utilities
│ ├── dto.go # Common data transfer objects
│ └── types.go # Common types
└── products/ # Product module
├── models.go # Database models
├── dto.go # Product DTOs
├── handlers.go # HTTP request handlers
├── services.go # Business logic layer
├── repositories.go # Data access layer
└── urls.go # Route definitions
```


## Architecture

The project follows a **layered architecture** pattern:

```
HTTP Handlers
    ↓
Business Logic (Services)
    ↓
Data Access (Repositories)
    ↓
Entity (Gorm Model)
    ↓
Database (SQLite)
```


### Layers

1. **Handlers** (`handlers.go`) - HTTP request/response handling
2. **Services** (`services.go`) - Business logic and validation
3. **Repositories** (`repositories.go`) - Database operations
4. **Models** (`models.go`) - Database entity definitions
5. **DTOs** (`dto.go`) - Data serialization/deserialization

## Database Schema

### Products Table

| Column     | Type      | Constraints                      |
|-----------|-----------|----------------------------------|
| ID        | TEXT      | Primary Key (UUID), Auto-generated |
| Name      | TEXT      | Not Null                         |
| Price     | INTEGER   | Not Null                         |
| Stock     | INTEGER   | Not Null, Default: 0             |
| CreatedAt | TIMESTAMP | Automatically set                |
| UpdatedAt | TIMESTAMP | Automatically set                |
| DeletedAt | TIMESTAMP | Soft delete support              |

## API Endpoints

All endpoints are prefixed with `/products`

### List All Products

```
GET /products
```

**Response** (200 OK):

```json
[
  {
    "id": "uuid-string",
    "name": "Product Name",
    "price": 10000,
    "stock": 5,
    "createdAt": "2026-02-10T00:00:00Z",
    "updatedAt": "2026-02-10T00:00:00Z"
  }
]
```

### Get Product Detail

```
GET /products/{id}
```

**Response** (200 OK):

```json
{
  "id": "uuid-string",
  "name": "Product Name",
  "price": 10000,
  "stock": 5,
  "createdAt": "2026-02-10T00:00:00Z",
  "updatedAt": "2026-02-10T00:00:00Z"
}
```

**Response** (404 Not Found):

```json
{
  "message": "Product not found"
}
```

### Create Product

```
POST /products
```

**Request Body**:

```json
{
  "name": "New Product",
  "price": 15000,
  "stock": 10
}
```

**Response** (201 Created):

```json
{
  "id": "uuid-string",
  "name": "New Product",
  "price": 15000,
  "stock": 10,
  "createdAt": "2026-02-10T00:00:00Z",
  "updatedAt": "2026-02-10T00:00:00Z"
}
```

### Update Product

```
PUT /products/{id}
```

**Request Body**:

```json
{
  "name": "Updated Product Name",
  "price": 20000,
  "stock": 8
}
```

**Response** (200 OK):

```json
{
  "id": "uuid-string",
  "name": "Updated Product Name",
  "price": 20000,
  "stock": 8,
  "createdAt": "2026-02-10T00:00:00Z",
  "updatedAt": "2026-02-10T10:30:00Z"
}
```

### Delete Product

```
DELETE /products/{id}
```

**Response** (204 No Content):

(Empty body)

## Getting Started

### Prerequisites

- Go 1.25.6 or higher
- SQLite3
- Git

### Installation

1. **Clone the repository**:

```bash
git clone https://github.com/yourusername/gocrud.git
cd gocrud
```

2. **Install dependencies**:

```bash
go mod download
```

3. **Build the application**:

```bash
go build -o gocrud
```

## Running the Application

### Prerequisites Check

Before running, make sure you have:

```bash
go version  # Should be Go 1.25.6 or higher
sqlite3 --version  # Should be installed
```

### Quick Start (3 Steps)

**Step 1: Navigate to project directory**

```bash
cd c:\Users\Bootrix\Documents\Bootrix\Learn\gocrud
```

**Step 2: Install dependencies** (first time only)

```bash
go mod download
```

**Step 3: Run the application**

```bash
go run main.go
```

**Expected Output**:

```
Server is running on port 8080
Press Ctrl+C to stop the server
```

### Accessing the API

Once the application is running, you can access it at:

```
http://localhost:8080
```

**Base URL for all endpoints**:

```
http://localhost:8080/products
```

### Port Information

- **Default Port**: `8080`
- **Base URL**: `http://localhost:8080`
- **API Prefix**: `/products`

#### Full Endpoint Examples:

- List products: `http://localhost:8080/products`
- Get product: `http://localhost:8080/products/abc123`
- Create product: `POST http://localhost:8080/products`
- Update product: `PUT http://localhost:8080/products/abc123`
- Delete product: `DELETE http://localhost:8080/products/abc123`

### Testing with Browser or Postman

**For GET requests (List & Detail)**, you can test directly in your browser:

1. Open your browser
2. Go to: `http://localhost:8080/products`
3. You should see a JSON response with all products

**For POST, PUT, DELETE requests**, use tools like:

- [Postman](https://www.postman.com/) - GUI tool (recommended)
- cURL - Command line tool
- Insomnia - Another API testing tool

### Building for Production

If you want to create an executable file:

```bash
# Build the executable
go build -o gocrud.exe

# Run the executable
./gocrud.exe

# On Linux/Mac
go build -o gocrud
./gocrud
```

The executable will create a SQLite database file (`gocrud.db`) in the same directory.

## Example Usage

### Using cURL

**1. Create a product**:

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "price": 10000000,
    "stock": 5
  }'
```

**2. List all products**:

```bash
curl http://localhost:8080/products
```

**3. Get a product by ID** (replace `{id}` with actual UUID):

```bash
curl http://localhost:8080/products/550e8400-e29b-41d4-a716-446655440000
```

**4. Update a product**:

```bash
curl -X PUT http://localhost:8080/products/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Laptop",
    "price": 11000000,
    "stock": 3
  }'
```

**5. Delete a product**:

```bash
curl -X DELETE http://localhost:8080/products/550e8400-e29b-41d4-a716-446655440000
```

### Using Postman

1. Open Postman
2. Create a new request
3. Set the method (GET, POST, PUT, DELETE)
4. Enter the URL: `http://localhost:8080/products`
5. For POST/PUT: Go to "Body" → Select "raw" → Choose "JSON" → Enter your JSON data
6. Click "Send"

## Troubleshooting

### Port Already in Use

If you get an error like "address already in use 8080":

```bash
# Windows - Find process using port 8080
netstat -ano | findstr :8080

# Kill the process (replace PID with actual number)
taskkill /PID <PID> /F

# Linux/Mac
lsof -i :8080
kill -9 <PID>
```

### Dependencies Not Installing

```bash
# Clear cache and reinstall
go clean -modcache
go mod download
```

### Database Issues

The SQLite database (`gocrud.db`) is automatically created on first run. If you need to reset it:

```bash
# Delete the database file
rm gocrud.db  # or del gocrud.db on Windows

# Run the application again
go run main.go
```

## Error Handling

The API returns standard HTTP status codes:

- `200 OK` - Successful GET, PUT request
- `201 Created` - Successful POST request
- `204 No Content` - Successful DELETE request
- `400 Bad Request` - Invalid request body or parameters
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

## License

This project is open source and available under the MIT License.

