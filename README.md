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
- **MinGW64** (required for Windows to compile SQLite driver)
- Git

### Installing MinGW64 (Windows Only)

SQLite requires compilation on Windows, so you need MinGW64:

**Step 1: Download MinGW64**

1. Go to: https://winlibs.com/#download-release
2. Download the latest **UCRT runtime** version (recommended: `winlibs-x86_64-posix-seh-gcc-XX.X.X-llvm-XX.X.X-mingw-w64-ucrt-release.zip`)

**Step 2: Extract and Setup**

1. Extract the downloaded ZIP file to a simple path (e.g., `C:\mingw64`)
2. Add MinGW64 to Windows PATH:
   - Open **Environment Variables** (search in Windows Start menu)
   - Click "Environment Variables" button
   - Under "User variables" or "System variables", click "New"
   - Variable name: `PATH`
   - Variable value: `C:\mingw64\bin` (or wherever you extracted it)
   - Click OK and restart your terminal

**Step 3: Verify Installation**

```bash
gcc --version
g++ --version
```

You should see version information for both commands.

### Installing Go

1. Download from: https://golang.org/dl/
2. Install Go 1.25.6 or higher
3. Verify installation:

```bash
go version
```

### Installation

1. **Clone the repository**:

```bash
git clone https://github.com/yourusername/gocrud.git
cd gocrud
```

2. **Install dependencies** (MinGW64 must be installed first):

```bash
go mod download
```

This will compile the `go-sqlite3` driver using MinGW64's `gcc`.

3. **Build the application**:

For Windows (with CGO enabled):

```bash
set CGO_ENABLED=1
go build -o gocrud.exe
```

For Linux/Mac:

```bash
CGO_ENABLED=1 go build -o gocrud
```

**Note**: `CGO_ENABLED=1` is required because GORM uses SQLite which needs C bindings (cgo) to compile.

## Running the Application

### Prerequisites Check

Before running, make sure you have all prerequisites installed:

```bash
go version        # Should be Go 1.25.6 or higher
gcc --version     # MinGW64 - Should show GCC version
```

### Quick Start (3 Steps)

**Step 1: Navigate to project directory**

```bash
cd c:\Users\Bootrix\Documents\Bootrix\Learn\gocrud
```

**Step 2: Install dependencies** (first time only)

```bash
set CGO_ENABLED=1
go mod download
```

**Step 3: Run the application**

**Option A: Direct run with CGO enabled**

```bash
set CGO_ENABLED=1
go run main.go
```

**Option B: Build then run**

```bash
set CGO_ENABLED=1
go build -o gocrud.exe
./gocrud.exe
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
# Windows - Build the executable with CGO enabled
set CGO_ENABLED=1
go build -o gocrud.exe

# Run the executable
./gocrud.exe

# Linux/Mac - Build the executable
CGO_ENABLED=1 go build -o gocrud
./gocrud
```

The executable will create a SQLite database file (`gocrud.db`) in the same directory.

## Environment Variables

### CGO_ENABLED

**Purpose**: Enables C bindings (cgo) which is required for SQLite compilation.

- When using GORM with SQLite, you **must** set `CGO_ENABLED=1`
- This allows Go to use C libraries (in this case, SQLite)
- Without this, you'll get compilation errors related to SQLite

**Setting CGO_ENABLED:**

Windows (Command Prompt):
```bash
set CGO_ENABLED=1
go build
```

Windows (PowerShell):
```powershell
$env:CGO_ENABLED = "1"
go build
```

Linux/Mac:
```bash
CGO_ENABLED=1 go build
```

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

### MinGW64 Not Found (Windows)

If you get error like `gcc: command not found`:

1. Download MinGW64 from: https://winlibs.com/#download-release
2. Extract to `C:\mingw64`
3. Add `C:\mingw64\bin` to Windows PATH
4. Restart your terminal/VS Code
5. Verify with: `gcc --version`

### SQLite Compilation Error

If you get errors during `go mod download` or `go build`:

```bash
# Enable CGO and try again
set CGO_ENABLED=1
go clean -modcache
go mod download
go build -o gocrud.exe
```

Common error messages:
- `gcc: command not found` → Install MinGW64
- `cgo: not enabled` → Set `CGO_ENABLED=1`
- Symbol errors during linking → Make sure MinGW64 is in PATH

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
# Enable CGO, clear cache and reinstall
set CGO_ENABLED=1
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

