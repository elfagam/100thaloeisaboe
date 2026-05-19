# 100 Tahun RSUD Prof. Dr. H. Aloei Saboe (RSAS)

Proyek ini adalah portal aplikasi web seremonial dan interaktif yang dibangun secara khusus untuk memperingati **Satu Abad (100 Tahun)** berdirinya RSUD Prof. Dr. H. Aloei Saboe. Aplikasi ini dirancang untuk ditampilkan pada layar besar (videotron/proyektor) saat acara puncak perayaan, dengan mengusung antarmuka sinematik, interaktif, dan premium.

## 🌟 Fitur Utama

### 1. Preshow & Pemutar Video Ucapan (`/preshow`)
Sesi khusus yang diputar sebelum acara inti dimulai. Menampilkan kompilasi video ucapan selamat dari para tokoh.
*   **Efek Layar Lebar Estetik**: Video tidak akan terpotong meskipun rasionya berbeda, berkat efek *Blurred Background* (latar belakang video yang diburamkan).
*   **Audio Ducking Otomatis**: Suara musik latar (`suara_latar.mp4`) otomatis mengecil saat video ucapan berbicara, dan mengeras kembali saat jeda transisi antar video.
*   **Panel Kontrol Rahasia**: Panel pengatur video yang hanya muncul saat kursor diarahkan, menjaga layar tetap bersih untuk penonton.
*   **Transisi Profesional**: Jeda 3 detik dengan animasi memuat (*loading*) elegan antar video.

### 2. Peluncuran VVIP Sinematik (`/launching`)
Halaman puncak peresmian yang interaktif, dioptimalkan untuk layar raksasa (Rasio 2:1, contoh 4x2 meter).
*   **Animasi Detak Jantung**: Representasi dedikasi RSAS dalam menyelamatkan nyawa, dengan animasi EKG.
*   **Staggered SVG Drawing**: Animasi memukau di mana garis EKG, angka 100, dan siluet gedung digambar secara berurutan dan *overlapping*.
*   **Aset Resolusi Tinggi (SVG)**: Logo-logo RSAS diubah menjadi format vektor murni agar tetap tajam sempurna meski ditampilkan di layar videotron raksasa.

### 3. Studio Tema Seremoni
Panel kontrol (*floating button* bergambar palet) untuk mengubah nuansa visual secara *real-time*:
*   **Preset Premium**: Tersedia berbagai opsi warna mulai dari *Centennial Gold*, *Royal Gold*, *Emerald Royal*, hingga *Dark Obsidian*.
*   **Penyesuaian Kecerahan**: Mode gelapkan layar (*Cinema Mode*) untuk meningkatkan fokus saat memutar konten penting.

### 4. Galeri, Refleksi, & Dashboard
Halaman-halaman penunjang untuk menampilkan pencapaian, sejarah panjang, serta momen-momen penting RSAS selama seabad mengabdi.

---

## 🛠️ Arsitektur Teknologi

Proyek ini menggunakan pemisahan arsitektur (*microservices*):

*   **Frontend (Antarmuka):** Dibangun dengan **Vue.js 3** + Vite. Menggunakan SCSS untuk penataan gaya yang kompleks dan animasi CSS yang efisien (hardware-accelerated).
*   **Backend (API):** Dibangun menggunakan bahasa **Go (Golang)**. Mengelola logika bisnis, autentikasi, serta berinteraksi dengan basis data.
*   **Basis Data & Penyimpanan:** Terintegrasi dengan layanan Google Cloud, seperti Firestore untuk data dan Google Cloud Storage untuk galeri.

---

## 🚀 Panduan Menjalankan Proyek Secara Lokal

Pastikan Anda telah menginstal `Node.js` (untuk Frontend) dan `Go` (untuk Backend) di komputer Anda.

### 1. Menjalankan Backend (Go API)
```bash
cd backend-rs100
# Unduh semua dependensi
go mod tidy
# Jalankan server
go run ./cmd/api
```
Server backend akan berjalan di `http://localhost:8081`.

### 2. Menjalankan Frontend (Vue.js)
Buka terminal/tab baru:
```bash
cd frontend-rs100
# Instal dependensi (hanya perlu dilakukan sekali)
npm install
# Jalankan server pengembangan
npm run dev
```
Aplikasi web akan dapat diakses melalui browser pada `http://localhost:5173`.

---

## 📂 Struktur Direktori Utama
*   `frontend-rs100/` : Berisi kode antarmuka aplikasi.
    *   `src/pages/` : Halaman-halaman utama (`Launching.vue`, `Preshow.vue`, `Home.vue`, dll).
    *   `src/components/` : Komponen antarmuka yang dapat digunakan kembali.
    *   `public/` : Aset-aset statis (Video, SVG, MP3).
*   `backend-rs100/` : Berisi kode sistem server dan API.
    *   `cmd/api/` : Titik masuk (*entry point*) aplikasi backend.
    *   `internal/` : Logika inti aplikasi (domain, repository, usecase).

---
*Dibuat untuk merayakan dedikasi tanpa batas RSUD Prof. Dr. H. Aloei Saboe (1926 - 2026).*
