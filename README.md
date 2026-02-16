# Backend POS Laundry

Backend sistem Point of Sale (POS) yang robust dan scalable untuk bisnis manajemen laundry yang dibangun dengan Go. REST API ini menyediakan fungsionalitas lengkap untuk mengelola pelanggan, inventori, layanan, transaksi, dan autentikasi pengguna.

## 📋 Daftar Isi

- [Fitur](#fitur)
- [Tech Stack](#tech-stack)
- [Prasyarat](#prasyarat)
- [Instalasi](#instalasi)
- [Konfigurasi](#konfigurasi)
- [Menjalankan Aplikasi](#menjalankan-aplikasi)
- [Dokumentasi API](#dokumentasi-api)
- [Struktur Proyek](#struktur-proyek)
- [Database](#database)
- [Berkontribusi](#berkontribusi)
- [Lisensi](#lisensi)

## ✨ Fitur

- **Manajemen Pengguna**: Registrasi pengguna, autentikasi, dan otorisasi
- **Manajemen Pelanggan**: Database pelanggan lengkap dengan informasi kontak
- **Manajemen Inventori**: Lacak persediaan dan bahan laundry
- **Manajemen Layanan**: Tentukan dan kelola layanan laundry dan harga
- **Pemrosesan Transaksi**: Tangani pesanan dan pembayaran pelanggan
- **Dashboard**: Analitik bisnis real-time dan statistik
- **Pencatatan Aktivitas**: Lacak semua aktivitas sistem untuk audit
- **Pencatatan Permintaan**: Pencatatan permintaan/respons komprehensif
- **Autentikasi**: Autentikasi aman berbasis JWT
- **Penanganan Error**: Respons error API standar
- **Paginasi**: Pengambilan data efisien dengan dukungan paginasi

## 🛠 Tech Stack

- **Bahasa**: Go 1.24.5
- **Framework**: Gin Gonic
- **Database**: PostgreSQL
- **Containerization**: Docker & Docker Compose
- **Logging**: Logrus
- **Autentikasi**: JWT (JSON Web Tokens)
- **Format Respons**: JSON

## 📦 Prasyarat

Sebelum memulai, pastikan Anda telah menginstal berikut ini:

- Go 1.24 atau lebih tinggi
- PostgreSQL 12 atau lebih tinggi
- Docker & Docker Compose (opsional, untuk setup containerized)
- Git

## 🚀 Instalasi

### 1. Clone Repositori

```bash
git clone <repository-url>
cd pos-laundry-be
```

### 2. Instal Dependensi

```bash
go mod download
go mod tidy
```

### 3. Atur Variabel Lingkungan

Buat file `.env` di direktori root dengan variabel berikut:

```env
# Konfigurasi Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=pos_laundry

# Konfigurasi Server
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# Konfigurasi JWT
JWT_SECRET=your_jwt_secret_key
JWT_EXPIRATION=24h

# Logging
LOG_LEVEL=info
```

## ⚙️ Konfigurasi

### Konfigurasi Database

Perbarui pengaturan koneksi database di `config/config.go`:

```go
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}
```

### Konfigurasi Server

Konfigurasi pengaturan server di `config/config.go`:

```go
type ServerConfig struct {
    Port string
    Host string
}
```

## ▶️ Menjalankan Aplikasi

### Opsi 1: Menggunakan Docker Compose (Direkomendasikan)

```bash
docker-compose up -d
```

Ini akan:
- Memulai database PostgreSQL
- Menjalankan migrasi database
- Memulai server API di port 8080

### Opsi 2: Menjalankan Secara Lokal

#### 1. Mulai PostgreSQL

```bash
# Menggunakan Docker
docker run -d \
  --name postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=pos_laundry \
  -p 5432:5432 \
  postgres:latest
```

#### 2. Jalankan Migrasi

```bash
go run cmd/api/main.go migrate
```

#### 3. Mulai Server

```bash
go run cmd/api/main.go
```

API akan tersedia di `http://localhost:8080`

## 📚 Dokumentasi API

### URL Dasar

```
http://localhost:8080/api
```

### Autentikasi

Sebagian besar endpoint memerlukan autentikasi JWT. Sertakan token di header Authorization:

```
Authorization: Bearer <your_jwt_token>
```

### Endpoint Utama

#### Pengguna
- `POST /users/register` - Daftar pengguna baru
- `POST /users/login` - Login pengguna
- `GET /users/{id}` - Dapatkan detail pengguna
- `PUT /users/{id}` - Update pengguna
- `DELETE /users/{id}` - Hapus pengguna

#### Pelanggan
- `GET /customers` - Daftar semua pelanggan (dengan paginasi)
- `POST /customers` - Buat pelanggan baru
- `GET /customers/{id}` - Dapatkan detail pelanggan
- `PUT /customers/{id}` - Update pelanggan
- `DELETE /customers/{id}` - Hapus pelanggan

#### Layanan
- `GET /services` - Daftar semua layanan
- `POST /services` - Buat layanan baru
- `GET /services/{id}` - Dapatkan detail layanan
- `PUT /services/{id}` - Update layanan
- `DELETE /services/{id}` - Hapus layanan

#### Inventori
- `GET /inventory` - Daftar item inventori
- `POST /inventory` - Tambahkan item inventori
- `PUT /inventory/{id}` - Update inventori
- `DELETE /inventory/{id}` - Hapus dari inventori

#### Transaksi
- `GET /transactions` - Daftar transaksi
- `POST /transactions` - Buat transaksi baru
- `GET /transactions/{id}` - Dapatkan detail transaksi
- `PUT /transactions/{id}` - Update transaksi

#### Dashboard
- `GET /dashboard/summary` - Dapatkan ringkasan bisnis
- `GET /dashboard/statistics` - Dapatkan data analitik

### Format Respons

**Respons Sukses (200)**
```json
{
  "status": "success",
  "data": {
    "id": "123",
    "name": "Item Name"
  }
}
```

**Respons Error**
```json
{
  "status": "error",
  "message": "Deskripsi error",
  "code": "ERROR_CODE"
}
```

## 📁 Struktur Proyek

```
pos-laundry-be/
├── cmd/
│   └── api/
│       └── main.go                 # Entry point aplikasi
├── config/
│   ├── config.go                   # Manajemen konfigurasi
│   └── types.go                    # Tipe konfigurasi
├── internal/
│   ├── entities/                   # Model data
│   ├── features/                   # Modul fitur
│   │   ├── customers/              # Fitur pelanggan (delivery, service, repository)
│   │   ├── dashboard/              # Fitur dashboard
│   │   ├── inventory/              # Fitur inventori
│   │   ├── services/               # Fitur layanan
│   │   ├── transactions/           # Fitur transaksi
│   │   └── users/                  # Fitur pengguna
│   ├── middleware/                 # Middleware HTTP
│   └── server/                     # Setup server
├── pkg/
│   ├── activitylogger/             # Utility pencatatan aktivitas
│   ├── contextutils/               # Utility context
│   ├── db/                         # Utility database
│   ├── httpErrors/                 # Penanganan error HTTP
│   ├── logger/                     # Logger aplikasi
│   ├── pagination/                 # Utility paginasi
│   └── utils/                      # Utility umum
├── docs/                           # Dokumentasi API
├── docker-compose.yml              # Konfigurasi Docker Compose
├── go.mod                          # File modul Go
└── README.md                       # File ini
```

### Pola Arsitektur

Proyek ini mengikuti pola **Clean Architecture** dengan pemisahan kepentingan yang jelas:

- **Delivery Layer** - Handler HTTP dan validasi permintaan
- **Service Layer** - Logika bisnis dan operasi domain
- **Repository Layer** - Persistensi data dan operasi database
- **Entity Layer** - Model data dan objek domain

## 🗄️ Database

### Skema PostgreSQL

Aplikasi menggunakan PostgreSQL dengan tabel-tabel utama berikut:

- **users** - Akun pengguna dan autentikasi
- **customers** - Informasi pelanggan
- **services** - Layanan laundry yang tersedia
- **inventory** - Item inventori dan tingkat stok
- **transactions** - Transaksi penjualan
- **transaction_items** - Item dalam transaksi
- **activity_logs** - Pencatatan aktivitas sistem
- **transaction_logs** - Riwayat transaksi

### Migrasi

Migrasi database dikelola secara otomatis. Jalankan migrasi dengan:

```bash
go run cmd/api/main.go migrate
```

## 🤝 Berkontribusi

1. Buat branch fitur: `git checkout -b feature/fitur-anda`
2. Buat perubahan Anda mengikuti struktur proyek
3. Uji secara menyeluruh
4. Commit dengan pesan yang jelas: `git commit -m 'Add: fitur anda'`
5. Push ke branch: `git push origin feature/fitur-anda`
6. Buka Pull Request

## 📝 Lisensi

Proyek ini dilisensikan di bawah Lisensi MIT - lihat file LICENSE untuk detail.

## 📧 Dukungan

Untuk dukungan dan pertanyaan, silakan hubungi tim pengembangan atau buka issue di repositori.

---

**Terakhir Diperbarui**: 16 Februari 2026  
**Versi**: 1.0.0
