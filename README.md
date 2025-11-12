# TemuLokal Microservices

Monorepo untuk microservices TemuLokal: Gateway, Auth Service, dan UMKM Service. Dokumen ini menjelaskan cara menjalankan masing-masing service, menyiapkan environment (.env), migrasi database, serta daftar endpoint beserta host dan port yang digunakan.

## Arsitektur & Port Organization

Gunakan rentang port terorganisir agar mudah diingat dan mencegah konflik:

- Gateway: `8000`
- Auth Service: `8001`
- UMKM Service: `8002`

Jika perlu menambah service baru, lanjutkan urutan berikutnya (`8003`, `8004`, ...). Hindari mencampur dengan port acak agar konfigurasi terstandarisasi.

> Catatan: Sesuaikan `.env` UMKM Service ke `8002` bila ingin mengikuti pola ini secara ketat.

## Prasyarat

1. Sudah ter-install Go (versi sesuai project) dan MySQL server berjalan.
2. Buat database kosong untuk masing-masing service (atau satu database terpisah sesuai kebutuhan saat ini):
	- Auth: `temulokal_auth_db` (contoh pada `.env`)
	- UMKM: `temulokal_umkm_db`
3. Clone repository dan masuk ke direktori root.

## Menyiapkan Environment (.env)

Setiap service memiliki file `.env.example`. Copy menjadi `.env` dan sesuaikan nilai variabelnya.

### Gateway `.env`
Variabel penting:
- `APP_PORT` (default 8000)
- `JWT_SECRET`
- `AUTH_SERVICE_ORIGIN` (misal `http://localhost:8001`)
- `UMKM_SERVICE_ORIGIN` (tambahkan bila belum, contoh `http://localhost:8002`)

Tambahkan baris berikut jika belum ada untuk origin UMKM:
```
UMKM_SERVICE_ORIGIN="http://localhost:8002"
```

### Auth Service `.env`
Variabel penting:
- `APP_NAME`, `APP_PORT=8001`, `APP_ENV`, `APP_ORIGIN`
- `DB_HOST`, `DB_USER`, `DB_PASS`, `DB_NAME`
- SMTP config (`SMTP_HOST`, `SMTP_PORT`, `SMTP_PASSWORD`, `SMTP_SENDER`) jika fitur email digunakan
- `JWT_SECRET`

Contoh minimal (sesuaikan password & host DB):
```
APP_NAME="Authentication Service"
APP_PORT=8001
APP_ENV=development
APP_ORIGIN="http://localhost:8001"

DB_HOST="localhost:3306"
DB_USER=root
DB_PASS=your_password
DB_NAME=temulokal_auth_db

JWT_SECRET="replace-with-secure-random"
```

### UMKM Service `.env`
Variabel penting:
- `UMKM_APP_NAME`, `UMKM_APP_PORT=8002` (ubah dari 8082 agar konsisten), `APP_ENV`, `APP_ORIGIN`
- `DB_HOST`, `DB_USER`, `DB_PASS`, `DB_NAME`

Contoh:
```
UMKM_APP_NAME="UMKM Service"
UMKM_APP_PORT=8002
APP_ENV=development
APP_ORIGIN="http://localhost:8002"

DB_HOST="localhost:3306"
DB_USER=root
DB_PASS=your_password
DB_NAME=temulokal_umkm_db
```

## Migrasi Database

Jalankan migrasi setelah `.env` dibuat dan database tersedia.

### Auth Service Migration
```
cd auth-service
go run main.go migrate
```

### UMKM Service Migration
```
cd umkm-service
go run main.go migrate
```

Pastikan tidak ada error. Tabel akan dibuat sesuai model.

## Seeding Data (UMKM)

Gunakan file `umkm-service/seeds/umkm_seed.sql` untuk mengisi data contoh:
```
mysql -h <host> -u <user> -p temulokal_umkm_db < umkm-service/seeds/umkm_seed.sql
```

## Menjalankan Service

Jalankan masing-masing service (disarankan urutan: auth, umkm, gateway).

### Auth Service Start
```
cd auth-service
go run main.go start
```

### UMKM Service Start
```
cd umkm-service
go run main.go start
```

### Gateway Start
```
cd gateway
go run main.go start
```

### Jalankan dengan Makefile (semua service sekaligus)

Cara cepat menjalankan seluruh service sekaligus:
```
make run
```

- Untuk menghentikan: tekan `Ctrl + C` pada terminal yang menjalankan `make run` (menghentikan proses `wait`), lalu jalankan:
```
make stop
```
Perintah `make stop` akan mencoba menghentikan proses `gateway`, `auth-service`, dan `umkm-service` yang masih berjalan.

### Menjalankan dengan Docker Compose

Alternatif orkestrasi container menggunakan Docker:
```
docker compose build
docker compose up -d
```

Service yang berjalan:
- MySQL: port 3306 (container `db`)
- Auth Service: port 8001 (container `auth-service`)
- UMKM Service: port 8002 (container `umkm-service`)
- Gateway: port 8000 (container `gateway`)

Log salah satu service:
```
docker compose logs -f auth-service
```

Hentikan semua:
```
docker compose down
```

Database otomatis dibuat lewat skrip `docker/initdb/01-create-databases.sql`. Jalankan migrasi di dalam container bila tabel belum terbentuk:
```
docker compose exec auth-service /app/auth-service migrate
docker compose exec umkm-service /app/umkm-service migrate
```

## Daftar Endpoint

### Gateway (Port 8000)
Semua request diproxy ke service terkait berdasarkan prefix.
- Auth: `http://localhost:8000/auth/*`
- UMKM: `http://localhost:8000/umkm/*`

### Auth Service (Port 8001)
Prefix: `/auth`
- `POST /auth/login`
- `POST /auth/register`

### UMKM Service (Port 8002)
Prefix: `/umkm`
- `GET /umkm/` — list semua UMKM dengan struktur nested (owner, gallery, location, opening_hours, rating aggregate)

## Build Semua Service Sekaligus

Gunakan target build pada Makefile untuk menghasilkan binary masing-masing service ke dalam folder `bin/`:
```
make build
```

Output:
- `bin/gateway`
- `bin/auth-service`
- `bin/umkm-service`

Menjalankan binary yang sudah dibuild (pastikan `.env` setiap service sudah benar):
```
./bin/auth-service
./bin/umkm-service
./bin/gateway
```

Catatan: Setiap service adalah binary terpisah (bukan digabung jadi satu proses). Target `make build` mempermudah build “semua jadi satu langkah” namun tetap menghasilkan tiga executable terpisah.

## Struktur Rating & Perhitungan

Tabel `umkm_ratings` menyimpan data individual `(id, umkm_id, value)` dengan value 1–5. Aggregasi `count` dan `star` (rata-rata) dihitung di layer usecase setiap kali endpoint list dipanggil.

## Konvensi Tambahan

- Tambah service baru: pilih port berikutnya dalam rentang 8000+.
- Hindari port 8080/3000 default untuk menjaga konsistensi internal.
- Simpan semua seed SQL di folder `seeds/` masing-masing service.

## Troubleshooting Cepat

- Error migrasi: cek koneksi DB (`DB_HOST`, user/password).
- Endpoint kosong: pastikan seed sudah dijalankan dan preload di repository sudah benar.
- Nilai rating 0: berarti belum ada data di `umkm_ratings` untuk UMKM tersebut.

## Next Steps (Opsional)

- Tambah endpoint detail UMKM: `GET /umkm/:id`.
- Tambah endpoint buat rating: `POST /umkm/:id/ratings`.
- Pagination & filtering di list UMKM.

---
Dokumen ini bisa diperbarui saat fitur baru ditambahkan.
