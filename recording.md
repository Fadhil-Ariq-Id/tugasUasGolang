# Panduan & Script Video Presentasi UAS Matematika Diskrit
**Judul:** Optimasi Rute Pengiriman Ice Cream Momoyo Menggunakan Algoritma Dijkstra
**Tenggat Waktu:** 1 Februari 2026

## 1. Persiapan (Setup)
Sebelum merekam, pastikan hal berikut siap:
- **Code Editor (VS Code):** Buka folder project. Buka tab `main.go` dan `dijkstra/dijkstra.go`.
- **Terminal:** Buka terminal di VS Code (pastikan berada di folder root project).
- **Slide/PDF:** Buka file `Laporan Jurnal_ Optimasi Rute Pengiriman Ice Cream Momoyo.pdf.pdf` untuk bagian Latar Belakang & Metode.
- **Software Rekaman:** OBS Studio / Zoom (Record Meeting) / Google Meet.
- **Mic & Kamera:** Pastikan suara kedua anggota terdengar jelas.

> **Tips Cepat:** Karena waktu sempit, gunakan Zoom/Gmeet. Share screen, nyalakan kamera keduanya, lalu 'Record to Cloud/Local'. Tidak perlu edit video yang rumit.

---

## 2. Struktur Video (Running Order)
Video ini dibagi menjadi 5 segmen utama.

| Segmen | Durasi Est. | Pembicara | Visual |
| :--- | :--- | :--- | :--- |
| **1. Pembukaan & Latar Belakang** | 1-2 Menit | Anggota 1 | Slide/PDF (Bagian Pendahuluan) |
| **2. Metode (Teori Graph/Dijkstra)** | 1-2 Menit | Anggota 2 | Slide/PDF (Bagian Metode) |
| **3. Penjelasan Code (Line-by-Line)** | 3-5 Menit | Anggota 1 & 2 | VS Code (`main.go` & `dijkstra.go`) |
| **4. Demo & Hasil** | 1 Menit | Anggota 1 | Terminal (Output Program) |
| **5. Kesimpulan MatDis** | 1 Menit | Anggota 2 | Wajah Presenter / Slide Penutup |

---

## 3. Outline Script
Berikut adalah *poin-poin bicara* (bukan harus hafal mati, tapi intinya harus tersampaikan).

### Bagian 1: Pembukaan & Latar Belakang (Anggota 1)
*   **Salam:** "Halo Pak/Bu, kami dari kelompok [Sebutkan Nama Kelompok]. Saya [Nama A] dan rekan saya [Nama B]."
*   **Judul:** "Kami akan mempresentasikan Final Project Matelmatika Diskrit tentang **Optimasi Rute Pengiriman Ice Cream Momoyo**."
*   **Masalah (Latar Belakang):**
    *   Focus ke PDF Laporan.
    *   "Latar belakang kasus ini adalah adanya kendala waktu dan biaya dalam pengiriman logistik Ice Cream Momoyo."
    *   "Kita perlu mencari rute terpendek dari Gudang (Sumber) ke berbagai cabang outlet agar pengiriman efisien."

### Bagian 2: Metode (Anggota 2)
*   **Penyelesaian:**
    *   "Untuk menyelesaikan masalah ini, kami menggunakan **Teori Graf**."
    *   "Lokasi (Gudang/Outlet) direpresentasikan sebagai **Node**."
    *   "Jalan antar lokasi direpresentasikan sebagai **Edge** dengan bobot berupa **Waktu Tempuh (menit)**."
*   **Algoritma:**
    *   "Metode spesifik yang kami pakai adalah **Algoritma Dijkstra**. Algoritma ini menjamin penemuan jalur terpendek dalam graf berbobot positif, yang sangat cocok untuk kasus rute jalan raya."

### Bagian 3: Penjelasan Code (Bergantian) - *Inti Penilaian*

**Transisi ke VS Code.**

**(Anggota 1 - Menjelaskan Struktur & Main):**
*   **Buka `main.go`**.
*   (Line 11-13) "Di fungsi `main`, pertama kita membangun Graph. `data.DataSample()` ini memuat data node dan bobot jalan yang sudah kami survei."
*   (Line 16-17) "Kemudian kita jalankan fungsi `dijkstra.Dijkstra`. Kita masukkan Graph `g`, Node Sumber (`SumberNode`), dan total Node (`JumlahNode`). Fungsi ini mengembalikan `dist` (jarak terpendek) dan `prev` (untuk melacak jalur)."
*   (Line 23-32) "Setelah perhitungan selesai, kita lakukan *looping* untuk mencetak hasil ke setiap tujuan. Kita pakai `ReconstructPath` untuk mengurutkan kembali jalurnya dari data `prev` tadi."

**(Anggota 2 - Menjelaskan Inti Algoritma):**
*   **Buka `dijkstra/dijkstra.go`**.
*   "Sekarang masuk ke logika intinya di paket `dijkstra`."
*   (Line 10-25) "Pertama, kami membuat Struct `Item` dan `PriorityQueue`. Ini boilerplate standar Go untuk mengimplementasikan Heap, supaya pencarian node dengan jarak terpendek bisa cepat (O(log V))."
*   (Line 45 - Fungsi Dijkstra) "Di fungsi `Dijkstra` ini algo bekerja:"
    *   (Line 51-55) "Inisialisasi `dist` dengan Infinity (MaxInt), kecuali source yang 0."
    *   (Line 63) "Masuk ke loop utama selama Queue tidak kosong."
    *   (Line 64-70) "`Pop` node dengan jarak terkecil. Jika sudah visited, skip."
    *   (Line 73-86) "**Relaksasi (Relaxation):** Kita cek semua tetangga. Jika jarak ke tetangga lewat node sekarang lebih pendek dari jarak yang tercatat sebelumnya (`alt < dist[e.To]`), maka kita update jaraknya dan simpan node sekarang sebagai `prev` (sebelumnya) dari tetangga tersebut. Lalu push ke Queue."
*   (Line 92 - ReconstructPath) "Fungsi ini cuma *backtracking*. Dari tujuan, kita tarik mundur lewat array `prev` sampai ketemu source, lalu urutannya dibalik."

### Bagian 4: Demo Hasil (Anggota 1)
*   **Buka Terminal**.
*   Ketik: `go run main.go`
*   **Analisa Output:**
    *   "Bisa dilihat di sini hasilnya."
    *   "Contoh: Dari Gudang ke 'Outlet X', Waktu Tempuh: 25 menit. Rutenya: Gudang -> Jalan A -> Jalan B -> Outlet X."
    *   "Program berhasil menghitung semua rute tercepat secara otomatis."

### Bagian 5: Kesimpulan (Anggota 2)
*   **Penutup Matkul:**
    *   "Kesimpulannya, mata kuliah Matematika Diskrit khususnya materi Graf dan Pohon (Tree) sangat relevan untuk dunia kerja, terutama bidang logistik dan optimasi sistem."
    *   "Dengan kode ini, efisiensi pengiriman Momoyo bisa meningkat."
*   **Salam Penutup:** "Sekian presentasi dari kelompok kami. Terima kasih."

---

## Checklist Akhir Sebelum Submit (1 Feb)
- [ ] Pastikan durasi video tidak terlalu panjang (target 5-10 menit cukup).
- [ ] Suara kedua anggota masuk.
- [ ] Code terlihat jelas di video (Zoom in font VS Code jika perlu: `Ctrl +`).
- [ ] Upload ke Google Drive/YouTube sesuai instruksi dosen dan kumpulkan Link-nya.
