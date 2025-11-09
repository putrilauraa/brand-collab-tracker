# brand-collab-tracker

# 📑 Brand Collaboration Tracker API

Aplikasi backend ini dirancang untuk membantu content creator (reviewer skincare & makeup) mengelola, melacak, dan memusatkan semua detail kerjasama (kolaborasi) dengan brand, dari SOW hingga status pembayaran.

**Project Overview:**
- Menggunakan Bahasa Go (Golang).
- Menerapkan JSON Web Token (JWT) Middleware untuk authorization.
- Menggunakan Database Relasional (PostgreSQL) dengan GORM.
- Memiliki 6 tabel berelasi.
- Di-deploy dengan Railway.

## 💡 Fitur Utama
- **Manajemen Relasi:** Mengaitkan Projects dengan Brand dan Category yang relevan.
- **Tracking Checklist:** Melacak status Task per proyek (misalnya, `Penerimaan Produk`, `Approval Konten`, `Pembayaran Fee`).
- **Keamanan:** Mengamankan semua endpoint menggunakan **JSON Web Token (JWT)**.

---

## 🛠️ Tech Stack & Deployment

| Komponen | Teknologi | Keterangan |
| :--- | :--- | :--- |
| **Bahasa** | **Go (Golang)** | Bahasa inti backend. |
| **Web Framework** | **Gin Gonic** | Routing dan HTTP handling. |
| **Database** | **PostgreSQL** | Database Relasional (Di-deploy di Railway). |
| **ORM** | **GORM** | Object-Relational Mapping. |
| **Authentication** | **JWT** | Otorisasi stateless. |

**URL Domain Publik (Railway):** `brand-collab-tracker-production.up.railway.app`

---

## 🏗️ Struktur Database (6 Tabel)

Aplikasi ini menggunakan 6 tabel utama yang saling berelasi:

| Tabel | Deskripsi | Relasi Kunci |
| :--- | :--- | :--- |
| `users` | Akun content creator. | — |
| `category_masters` | Data master kategori produk. | — |
| `brands` | Detail brand klien. | 1:1 ke `category_masters` |
| `projects` | Detail setiap SOW/kerjasama. | **1:M** ke `tasks` dan `attachments` |
| `tasks` | Item checklist (Tracking status). | M:1 ke `projects` |
| `project_attachments` | Link ke brief atau dokumen penting. | M:1 ke `projects` |

---

## 💻 Setup Lokal & Run Project

1.  **Clone Repository:**
    ```bash
    git clone https://github.com/putrilauraa/brand-collab-tracker.git
    cd brand-collab-tracker
    ```
2.  **Inisiasi Modules:**
    ```bash
    go mod tidy
    ```
3.  **Konfigurasi Database Lokal:**
    * Pastikan server **PostgreSQL** lokal Anda berjalan.
    * Buat database kosong baru.
4.  **Buat File `.env`:**
    Buat file `.env` di root proyek dan isi dengan kredensial lokal Anda:
    ```ini
    # Kredensial Lokal (Digunakan oleh config/database.go)
    PGHOST=localhost
    PGPORT=5432
    PGUSER=postgres
    PGPASSWORD=[PASSWORD DB LOKAL]
    PGDATABASE=[NAMA DB LOKAL]

    # Secret Wajib (Harus sama dengan di Railway Variables)
    JWT_SECRET=[SECRET_KEY_PANJANG_DAN_UNIK_ANDA]
    SERVER_PORT=8080
    ```
5.  **Jalankan Aplikasi:**
    ```bash
    go run main.go
    ```

---

## 🔑 Dokumentasi API Endpoints

Semua endpoint diuji menggunakan domain publik Railway Anda. Gunakan **Postman** untuk pengujian. Endpoint yang terlindungi **wajib** menyertakan **`Authorization: Bearer [TOKEN]`**.

### 1. Authentication (Public)

| Method | Path | Keterangan |
| :--- | :--- | :--- |
| `POST` | `/api/users/register` | Membuat akun baru. |
| `POST` | `/api/users/login` | Mendapatkan **JWT Token** (Wajib). |

### 2. Features (Protected)

| Method | Path | Keterangan | Relasi |
| :--- | :--- | :--- | :--- |
| **`POST`** | `/api/categories` | CREATE Category. | — |
| **`GET`** | `/api/categories` | READ ALL Categories. | — |
| **`PUT`** | `/api/categories/:id` | **UPDATE Category** (Memperbarui nama). | — |
| **`DELETE`** | `/api/categories/:id` | DELETE Category. | — |
| **`POST`** | `/api/brands` | CREATE Brand. | Requires `category_id` |
| **`GET`** | `/api/brands` | READ ALL Brands. | — |
| **`GET`** | `/api/brands/:id` | READ Brand by ID. | — |
| **`PUT`** | `/api/brands/:id` | **UPDATE Brand**. | Requires `category_id` |
| **`DELETE`** | `/api/brands/:id` | DELETE Brand. | — |
| **`POST`** | `/api/projects` | CREATE Project. | Requires `brand_id` |
| **`GET`** | `/api/projects` | GET ALL Projects. | - |
| **`GET`** | `/api/projects/:id` | GET Project by ID. | - |
| **`PUT`** | `/api/projects/:id` | **UPDATE Project** (Mengubah status, fee, dll.). | Requires `brand_id` |
| **`DELETE`** | `/api/projects/:id` | DElETE Project. | - |
| **`GET`** | `/api/:id/tasks` | GET Task by Project ID. | - |
| **`GET`** | `/api/:id/attachments` | GET Attachement by Project ID. | - |
| **`POST`** | `/api/tasks` | CREATE Task/Checklist. | Requires `project_id` |
| **`GET`** | `/api/tasks/:id` | GET Task by ID. | - |
| **`PUT`** | `/api/tasks/:id` | **UPDATE Task** (Mengubah status `is_completed`, notes, dll.). | Requires `project_id` |
| **`DELETE`** | `/api/tasks/:id` | DElETE Task. | - |
| **`POST`** | `/api/attachments` | CREATE Attachment (Link Brief/SOW). | Requires `project_id` |
| **`GET`** | `/api/attachments/:id` | GET Attachment by ID. | - |
| **`PUT`** | `/api/attachments/:id` | **UPDATE Attachment** (Mengubah URL file). | Requires `project_id` |
| **`DELETE`** | `/api/attachments/:id` | DElETE Attachement. | - |
