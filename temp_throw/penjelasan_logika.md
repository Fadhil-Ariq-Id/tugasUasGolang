# Panduan Logika & Hafalan Coding (Edisi Simple Teacher)

File ini dibuat khusus untuk melatih otak memahami **TUJUAN** di balik setiap baris code. Kita tidak hanya membaca kode, tapi mengerti "kenapa kode ini ada di sini".

---

## 1. Kamus Hafalan Syntax (Wajib Ingat!)

Ini adalah "kosakata" bahasa Go yang akan sering kamu temui. Hafalkan tujuannya!

| Syntax | Nama | Tujuan / Fungsi Utama (Hafalan) | Contoh Analogi |
| :--- | :--- | :--- | :--- |
| **`make(...)`** | **Alokasi Memori** | "Tolong siapkan tempat kosong di memori komputer untuk data saya." Tanpa ini, data tipe map/slice akan *nil* (kosong melompong) dan error kalau diisi. | Seperti menyewa ruko kosong sebelum mulai jualan. |
| **`{}`** | **Kurung Kurawal** | 1. **Data**: "Ini isi paket datanya". <br> 2. **Scope**: "Ini wilayah kekuasaan fungsi/loop ini". | 1. Kardus paket. <br> 2. Pagar rumah. |
| **`*`** | **Pointer** | "Jangan copy datanya, kasih tau ALAMAT aslinya aja". Biar hemat memori dan perubahan tersimpan di data asli. | Mengasih "Google Maps location" alih-alih memindahkan gedung rumahnya ke kamu. |
| **`time`** | **Waktu** | Untuk mengukur seberapa cepat kodemu berlari (performance benchmark). | Stopwatch pelatih lari. |
| **`struct`** | **Struktur Data** | Wadah custom untuk mengelompokkan data yang berbeda tipe jadi satu kesatuan. | Formulir KTP (ada nama string, umur int, tgl lahir date). |

---

## 2. Konsep Dasar Konseptual: Graph & Node

Bayangkan aplikasi ini adalah **Aplikasi Ojek Online**.

1.  **Node (Titik)**: Adalah **LOKASI**.
    *   Contoh: "Mall Kelapa Gading", "Grand Indonesia".
    *   Di kode: Direpresentasikan dengan `int` (ID angka, misal 0, 1, 2) biar komputer cepat memprosesnya.
2.  **Edge (Garis Hubung)**: Adalah **JALAN RAYA**.
    *   Menghubungkan dua Node.
    *   Punya **Weight (Bobot)**: Di sini bobotnya adalah **Waktu Tempuh (menit)**. Semakin kecil angkanya, semakin cepat (macetnya sedikit).
3.  **Graph**: Adalah **PETA KESELURUHAN** yang berisi semua Lokasi (Nodes) dan Jalan (Edges).

---

## 3. Bedah Logic Per File (Interaksi Efisien)

### A. `graph/graph.go` -> "Si Pembuat Peta"
**Tujuan**: Menyiapkan struktur data. Belum ada isinya, cuma definisinya.

*   **Logic Penting**:
    ```go
    type Graph struct {
        Nodes map[int]string // Kamus: ID 1 artinya "Mall A"
        Edges map[int][]Edge // Navigasi: Dari ID 1 bisa ke mana aja?
    }
    ```
*   **Kenapa pakai `*Graph` di receiver function `func (g *Graph)`?**
    *   Agar saat kita `AddNode` (tambah lokasi), kita menambahkannya ke Peta Asli (`g`), bukan ke fotokopian peta. Efisiensi memori tingkat tinggi!

### B. `data/data.go` -> "Si Pengisi Data"
**Tujuan**: Mengisi Peta kosong tadi dengan data nyata dari Laporan/Dokumen.

*   **Logic Penting**:
    1.  `g := graph.NewGraph()` -> Panggil tukang peta pro (bikin map kosong pakai `make`).
    2.  `g.AddNode(...)` -> Tancapkan paku lokasi di peta.
    3.  `g.AddEdge(...)` -> Gambar garis jalan antar paku.
*   **Efisiensi**: Memisahkan data dari logika. Kalau ada jalan baru dibangun, kita cuma ubah file ini. File `dijkstra` dan `main` tidak perlu disentuh.

### C. `dijkstra/dijsktra.go` -> "Si Otak Pintar (GPS)"
**Tujuan**: Mencari rute tercepat. Ini bagian paling rumit tapi paling penting.

*   **Logic Priority Queue (Antrian Prioritas)**:
    *   Bayangkan antrian masuk bioskop, tapi yang tiketnya paling murah (jarak terpendek) boleh nyerobot ke depan.
    *   **Tujuan**: Agar komputer SELALU memproses jalan yang paling menjanjikan duluan. Ini kunci efisiensi algoritma Dijkstra.
*   **Logic `Dijkstra(...)`**:
    *   `dist` array: Papan skor jarak tercepat ke setiap lokasi. Awalnya infinity (tak terhingga), pelan-pelan diupdate jadi angka real.
    *   `prev` array: "Jejak kaki". Untuk tahu "saya sampai ke Grand Indonesia tadi lewat mana ya yang paling cepat?". Ini dipakai buat `ReconstructPath`.

### D. `main.go` -> "Si Mandor / Controller"
**Tujuan**: Mengatur alur kerja program dari awal sampai akhir.

*   **Alur Interaksi (Flow Hafalan)**:
    1.  **Start Stopwatch**: `start := time.Now()` (Ayo mulai hitung!).
    2.  **Bikin Peta**: `g := data.BuildGraphFromDoc()` (Ambil data).
    3.  **Jalankan GPS**: `Dijkstra(g, ...)` (Mikir keras cari jalan).
    4.  **Stop Stopwatch**: `time.Since(start)` (Selesai! Berapa lama mikirnya?).
    5.  **Lapor Hasil**: `fmt.Printf` (Cetak struk laporan ke layar user).

---

## Rangkuman Interaksi Node (Simulasi Otak)

1.  **`main`** teriak: "Woi `data`, mana petanya?"
2.  **`data`** jawab: "Nih, peta graph udah gw isi lengkap node & edge-nya pakai `make` biar siap pakai."
3.  **`main`** teriak lagi: "Oke, sekarang `dijkstra`, cariin jalan dari Warehouse ke semua mall!"
4.  **`dijkstra`** mikir:
    *   "Gw mulai dari Warehouse (Node 0)."
    *   "Cek tetangga.. oh bisa ke Cibinong (waktu 15). Catat."
    *   "Masukin ke Priority Queue. Yang paling cepat gw proses lagi."
    *   "Ulang terus sampai semua ketemu."
    *   "Nih boss `main`, hasilnya (dist & prev)."
5.  **`main`** tutup: "Oke, gw print laporannya sekarang. Thanks semua."

***

**Tips Hafalan Terakhir:**
*   Lihat `{}` -> Ingat **Wadah/Isi**.
*   Lihat `*` -> Ingat **Asli/Hemat**.
*   Lihat `make` -> Ingat **Sewa Tempat**.
