# GoCRUD - API Manajemen Produk

## Pengenalan

GoCRUD adalah REST API yang dibangun dengan Go untuk mengelola produk. Proyek ini mendemonstrasikan operasi CRUD (Create, Read, Update, Delete) dengan menggunakan **Clean Architecture** menggunakan repository pattern, service layer, dan data transfer objects (DTOs).

## Technology Stack

- **Bahasa**: Go 1.25.6
- **Web Framework**: [Chi Router](https://github.com/go-chi/chi) - Lightweight HTTP routing
- **Database**: SQLite dengan [GORM](https://gorm.io/) - Golang ORM
- **Library Tambahan**:
  - `google/uuid` - Untuk generate UUID
  - `go-sqlite3` - SQLite driver

## Struktur Proyek

```
gocrud/
├── main.go # Entry point aplikasi
├── go.mod # Go module definition
├── README.md # Dokumentasi
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

Proyek ini mengikuti pola **layered architecture**:

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

### Layer-layer

1. **Handlers** (`handlers.go`) - Menangani HTTP request/response
2. **Services** (`services.go`) - Business logic dan validasi
3. **Repositories** (`repositories.go`) - Database operations
4. **Models** (`models.go`) - Database entity definitions
5. **DTOs** (`dto.go`) - Data serialization/deserialization

## Database Schema

### Tabel Products

| Column     | Type      | Constraints                      |
|-----------|-----------|----------------------------------|
| ID        | TEXT      | Primary Key (UUID), Auto-generated |
| Name      | TEXT      | Not Null                         |
| Price     | INTEGER   | Not Null                         |
| Stock     | INTEGER   | Not Null, Default: 0             |
| CreatedAt | TIMESTAMP | Otomatis diatur                  |
| UpdatedAt | TIMESTAMP | Otomatis diatur                  |
| DeletedAt | TIMESTAMP | Soft delete support              |

## API Endpoints

Semua endpoint dimulai dengan `/products`

### Daftar Semua Produk

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

### Lihat Detail Produk

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

### Buat Produk Baru

```
POST /products
```

**Request Body**:

```json
{
  "name": "Produk Baru",
  "price": 15000,
  "stock": 10
}
```

**Response** (201 Created):

```json
{
  "id": "uuid-string",
  "name": "Produk Baru",
  "price": 15000,
  "stock": 10,
  "createdAt": "2026-02-10T00:00:00Z",
  "updatedAt": "2026-02-10T00:00:00Z"
}
```

### Update Produk

```
PUT /products/{id}
```

**Request Body**:

```json
{
  "name": "Nama Produk Diperbarui",
  "price": 20000,
  "stock": 8
}
```

**Response** (200 OK):

```json
{
  "id": "uuid-string",
  "name": "Nama Produk Diperbarui",
  "price": 20000,
  "stock": 8,
  "createdAt": "2026-02-10T00:00:00Z",
  "updatedAt": "2026-02-10T10:30:00Z"
}
```

### Hapus Produk

```
DELETE /products/{id}
```

**Response** (204 No Content):

(Body kosong)

## Mulai Menggunakan

### Prerequisite

- Go 1.25.6 atau lebih tinggi
- SQLite3
- **MinGW64** (diperlukan untuk Windows agar bisa compile SQLite driver)
- Git

### Install MinGW64 (Windows Aja)

SQLite butuh di-compile di Windows, jadi kamu perlu MinGW64:

**Step 1: Download MinGW64**

1. Buka: https://winlibs.com/#download-release
2. Download versi terbaru **UCRT runtime** (recommended: `winlibs-x86_64-posix-seh-gcc-XX.X.X-llvm-XX.X.X-mingw-w64-ucrt-release.zip`)

**Step 2: Extract dan Setup**

1. Extract ZIP file ke path yang simpel (contoh: `C:\mingw64`)
2. Tambahkan MinGW64 ke Windows PATH:
   - Buka **Environment Variables** (cari di Windows Start menu)
   - Klik tombol "Environment Variables"
   - Di bawah "User variables" atau "System variables", klik "New"
   - Variable name: `PATH`
   - Variable value: `C:\mingw64\bin` (atau path tempat kamu extract-nya)
   - Klik OK dan restart terminal

**Step 3: Verifikasi Instalasi**

```bash
gcc --version
g++ --version
```

Kamu seharusnya melihat informasi versi untuk kedua command tersebut.

### Install Go

1. Download dari: https://golang.org/dl/
2. Install Go 1.25.6 atau lebih tinggi
3. Verifikasi instalasi:

```bash
go version
```

### Instalasi

1. **Clone repository**:

```bash
git clone https://github.com/yourusername/gocrud.git
cd gocrud
```

2. **Install dependencies** (MinGW64 harus sudah diinstall):

```bash
set CGO_ENABLED=1
go mod download
```

Ini akan compile `go-sqlite3` driver menggunakan MinGW64's `gcc`.

3. **Build aplikasi**:

Untuk Windows (dengan CGO enabled):

```bash
set CGO_ENABLED=1
go build -o gocrud.exe
```

Untuk Linux/Mac:

```bash
CGO_ENABLED=1 go build -o gocrud
```

**Note**: `CGO_ENABLED=1` diperlukan karena GORM pakai SQLite yang butuh C bindings (cgo) untuk compile.

## Menjalankan Aplikasi

### Cek Prerequisite

Sebelum jalanin, pastikan semua prerequisite sudah terinstall:

```bash
go version        # Harus Go 1.25.6 atau lebih tinggi
gcc --version     # MinGW64 - Harus menunjukkan GCC version
```

### Quick Start (3 Langkah)

**Step 1: Buka folder proyek**

```bash
cd c:\Users\Bootrix\Documents\Bootrix\Learn\gocrud
```

**Step 2: Install dependencies** (cuma pertama kali)

```bash
set CGO_ENABLED=1
go mod download
```

**Step 3: Jalankan aplikasi**

**Opsi A: Jalankan langsung dengan CGO enabled**

```bash
set CGO_ENABLED=1
go run main.go
```

**Opsi B: Build dulu baru jalanin**

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

### Akses API

Setelah aplikasi jalan, kamu bisa akses di:

```
http://localhost:8080
```

**Base URL untuk semua endpoints**:

```
http://localhost:8080/products
```

### Informasi Port

- **Default Port**: `8080`
- **Base URL**: `http://localhost:8080`
- **API Prefix**: `/products`

#### Contoh endpoint lengkap:

- Daftar produk: `http://localhost:8080/products`
- Lihat produk: `http://localhost:8080/products/abc123`
- Buat produk: `POST http://localhost:8080/products`
- Update produk: `PUT http://localhost:8080/products/abc123`
- Hapus produk: `DELETE http://localhost:8080/products/abc123`

### Testing dengan Browser atau Postman

**Untuk GET requests (Daftar & Detail)**, kamu bisa test langsung di browser:

1. Buka browser
2. Pergi ke: `http://localhost:8080/products`
3. Kamu akan lihat JSON response dengan semua produk

**Untuk POST, PUT, DELETE requests**, pakai tools seperti:

- [Postman](https://www.postman.com/) - GUI tool (recommended)
- cURL - Command line tool
- Insomnia - API testing tool lainnya

### Build untuk Production

Kalau mau buat executable file:

```bash
# Windows - Build executable dengan CGO enabled
set CGO_ENABLED=1
go build -o gocrud.exe

# Jalanin executable
./gocrud.exe

# Linux/Mac - Build executable
CGO_ENABLED=1 go build -o gocrud
./gocrud
```

Executable akan membuat SQLite database file (`gocrud.db`) di folder yang sama.

## Environment Variables

### CGO_ENABLED

**Fungsi**: Mengaktifkan C bindings (cgo) yang diperlukan untuk compile SQLite.

- Kalau pakai GORM dengan SQLite, kamu **harus** set `CGO_ENABLED=1`
- Ini memungkinkan Go untuk pakai C libraries (dalam hal ini SQLite)
- Tanpa ini, kamu akan dapat error saat compile yang berkaitan dengan SQLite

**Cara set CGO_ENABLED:**

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

## Contoh Penggunaan

### Pakai cURL

**1. Buat produk**:

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "price": 10000000,
    "stock": 5
  }'
```

**2. Daftar semua produk**:

```bash
curl http://localhost:8080/products
```

**3. Lihat produk by ID** (ganti `{id}` dengan UUID asli):

```bash
curl http://localhost:8080/products/550e8400-e29b-41d4-a716-446655440000
```

**4. Update produk**:

```bash
curl -X PUT http://localhost:8080/products/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop Diperbarui",
    "price": 11000000,
    "stock": 3
  }'
```

**5. Hapus produk**:

```bash
curl -X DELETE http://localhost:8080/products/550e8400-e29b-41d4-a716-446655440000
```

### Pakai Postman

1. Buka Postman
2. Buat request baru
3. Pilih method-nya (GET, POST, PUT, DELETE)
4. Input URL: `http://localhost:8080/products`
5. Untuk POST/PUT: Pergi ke "Body" → Pilih "raw" → Pilih "JSON" → Masukkin JSON data-nya
6. Klik "Send"

## Troubleshooting

### MinGW64 Tidak Ketemu (Windows)

Kalau dapat error seperti `gcc: command not found`:

1. Download MinGW64 dari: https://winlibs.com/#download-release
2. Extract ke `C:\mingw64`
3. Tambahkan `C:\mingw64\bin` ke Windows PATH
4. Restart terminal/VS Code
5. Verifikasi dengan: `gcc --version`

### Error SQLite Compilation

Kalau dapat error saat `go mod download` atau `go build`:

```bash
# Aktifin CGO dan coba lagi
set CGO_ENABLED=1
go clean -modcache
go mod download
go build -o gocrud.exe
```

Pesan error umum:
- `gcc: command not found` → Install MinGW64
- `cgo: not enabled` → Set `CGO_ENABLED=1`
- Symbol errors during linking → Pastikan MinGW64 ada di PATH

### Port Sudah Terpakai

Kalau dapat error seperti "address already in use 8080":

```bash
# Windows - Cari proses yang pakai port 8080
netstat -ano | findstr :8080

# Kill proses-nya (ganti PID dengan angka asli)
taskkill /PID <PID> /F

# Linux/Mac
lsof -i :8080
kill -9 <PID>
```

### Dependencies Tidak Terinstall

```bash
# Aktifin CGO, clear cache dan reinstall
set CGO_ENABLED=1
go clean -modcache
go mod download
```

### Masalah Database

SQLite database (`gocrud.db`) otomatis dibuat saat first run. Kalau mau reset:

```bash
# Hapus database file
del gocrud.db  # Windows
rm gocrud.db   # Linux/Mac

# Jalanin aplikasi lagi
go run main.go
```

## Error Handling

API mengembalikan standard HTTP status codes:

- `200 OK` - Successful GET, PUT request
- `201 Created` - Successful POST request
- `204 No Content` - Successful DELETE request
- `400 Bad Request` - Invalid request body atau parameters
- `404 Not Found` - Resource tidak ketemu
- `500 Internal Server Error` - Server error

## License

Proyek ini open source dan tersedia di bawah MIT License.

