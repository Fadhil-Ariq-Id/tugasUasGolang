# Rencana dan Script Presentasi: Optimasi Rute Pengiriman Ice Cream Momoyo

Ini adalah draf materi dan skrip untuk presentasi video Anda. Anda bisa membaginya dengan rekan Anda. Bagian yang ditandai `[NAMA 1]` dan `[NAMA 2]` adalah saran pembagian bicara.

---

## 1. Pembukaan dan Latar Belakang

**[NAMA 1]**

**(Slide 1: Judul)**
"Selamat pagi/siang/sore, Bapak/Ibu dosen. Kami dari kelompok [Sebutkan nama kelompok] yang beranggotakan saya, [Nama 1], dan rekan saya..."

**[NAMA 2]**

"...saya [Nama 2]. Pada kesempatan ini, kami akan mempresentasikan hasil proyek Ujian Akhir Semester kami untuk mata kuliah Matematika Diskrit dengan judul 'Optimasi Rute Pengiriman Ice Cream Momoyo Menggunakan Algoritma Dijkstra'."

**(Slide 2: Latar Belakang Masalah)**

**[NAMA 1]**

"Latar belakang dari studi kasus yang kami angkat adalah sebuah perusahaan es krim fiktif bernama 'Momoyo'. Perusahaan ini memiliki satu gudang utama di Cibubur dan harus mendistribusikan produknya ke 11 cabang yang tersebar di berbagai mal di area Jabodetabek."

"Masalah utamanya adalah efisiensi. Setiap pengiriman memiliki biaya operasional, terutama dari segi waktu tempuh. Semakin lama waktu perjalanan, semakin tinggi biayanya dan semakin besar risiko produk mencair. Oleh karena itu, menemukan rute tercepat dari gudang ke setiap cabang adalah hal yang krusial."

"Tujuan dari proyek ini adalah untuk membangun sebuah model sistem yang dapat secara otomatis menghitung dan menentukan rute pengiriman terpendek (dalam hal ini, tercepat berdasarkan waktu tempuh) dari gudang pusat ke semua cabang mal tujuan."

---

## 2. Metode yang Digunakan

**(Slide 3: Konsep Graf)**

**[NAMA 2]**

"Untuk memodelkan masalah ini, kami menggunakan struktur data Graf dari teori Matematika Diskrit. Graf ini memetakan jaringan jalan antara lokasi-lokasi yang ada."

*   **Node (Simpul)**: Setiap lokasi, baik itu gudang pusat Momoyo maupun mal cabang, kami representasikan sebagai sebuah 'node' atau 'simpul' dalam graf. Total ada 12 simpul dalam model kami.
*   **Edge (Sisi)**: Setiap rute jalan yang menghubungkan dua lokasi secara langsung kami representasikan sebagai 'edge' atau 'sisi'.
*   **Weight (Bobot)**: Setiap 'edge' memiliki 'weight' atau 'bobot', yang dalam kasus ini adalah estimasi waktu tempuh dalam satuan menit antara dua lokasi yang terhubung."

"Graf yang kami gunakan adalah **graf tidak berarah (undirected graph)**, yang berarti jika ada jalan dari Lokasi A ke B dengan waktu tempuh X menit, maka kami asumsikan jalan dari B ke A juga memakan waktu X menit."

**(Slide 4: Algoritma Dijkstra)**

**[NAMA 1]**

"Setelah jaringan jalan berhasil dimodelkan ke dalam bentuk graf, langkah selanjutnya adalah menemukan rute terpendek. Untuk ini, kami mengimplementasikan **Algoritma Dijkstra**."

"Secara singkat, cara kerja Algoritma Dijkstra adalah sebagai berikut:"
1.  "Algoritma dimulai dari satu titik awal, yang dalam kasus kami adalah **Gudang Momoyo (Source Node)**."
2.  "Jarak dari titik awal ke dirinya sendiri diatur menjadi 0, sedangkan jarak ke semua titik lain diatur menjadi tak terhingga."
3.  "Algoritma kemudian secara berulang memilih simpul yang belum dikunjungi dengan jarak terpendek dari titik awal."
4.  "Dari simpul terpilih, algoritma akan memeriksa semua tetangganya dan memperbarui (merelaksasi) jarak mereka jika ditemukan rute yang lebih pendek melalui simpul saat ini."
5.  "Proses ini terus berlanjut hingga semua simpul yang dapat dijangkau telah dikunjungi. Hasil akhirnya adalah sebuah set data yang berisi jarak terpendek dari *source* ke semua simpul lain, serta 'peta' untuk merekonstruksi jalur tersebut."

"Algoritma ini sangat cocok untuk kasus kita karena dapat secara efisien menemukan rute tercepat dari satu sumber ke banyak tujuan."

---

## 3. Hasil dan Pembahasan

**(Slide 5: Tabel Hasil Rute Terpendek)**

**[NAMA 2]**

"Setelah kami mengimplementasikan model graf dan algoritma Dijkstra ke dalam program Go, kami menjalankan program tersebut. Berikut adalah hasil output yang menunjukkan rute pengiriman tercepat dari Gudang Momoyo di Cibubur ke setiap mal cabang:"

"Program kami menghasilkan output berupa tabel yang berisi **Tujuan**, **Total Waktu Tempuh**, dan **Rute** yang harus dilalui. Berikut adalah hasilnya:"

| Tujuan                   | Waktu Tempuh | Rute                                                                                             |
| ------------------------ | ------------ | ------------------------------------------------------------------------------------------------ |
| Mall Kelapa Gading       | 60 menit     | `Warehouse Momoyo (Cibubur) -> Bekasi Cyber Park -> Mall Kelapa Gading`                          |
| Grand Indonesia (GI)     | 70 menit     | `Warehouse Momoyo (Cibubur) -> Depok Town Square -> Grand Indonesia (GI)`                        |
| Pondok Indah Mall (PIM)  | 100 menit    | `Warehouse Momoyo (Cibubur) -> Depok Town Square -> Grand Indonesia (GI) -> Pondok Indah Mall (PIM)` |
| BSD City                 | 135 menit    | `... -> Grand Indonesia (GI) -> Pondok Indah Mall (PIM) -> BSD City`                             |
| Bekasi Cyber Park        | 25 menit     | `Warehouse Momoyo (Cibubur) -> Bekasi Cyber Park`                                                |
| Depok Town Square        | 30 menit     | `Warehouse Momoyo (Cibubur) -> Depok Town Square`                                                |
| Tangerang City           | 120 menit    | `Warehouse Momoyo (Cibubur) -> Depok Town Square -> Grand Indonesia (GI) -> Tangerang City`        |
| Bogor Trade Mall         | 35 menit     | `Warehouse Momoyo (Cibubur) -> Bogor Trade Mall`                                                 |
| Cibinong City Mall (CCM) | 15 menit     | `Warehouse Momoyo (Cibubur) -> Cibinong City Mall (CCM)`                                         |
| Senayan City             | 85 menit     | `Warehouse Momoyo (Cibubur) -> Depok Town Square -> Grand Indonesia (GI) -> Senayan City`        |
| Summarecon Mall Bekasi   | 45 menit     | `Warehouse Momoyo (Cibubur) -> Bekasi Cyber Park -> Summarecon Mall Bekasi`                      |

"Dari hasil ini, kita bisa melihat dengan jelas rute tercepat ke setiap tujuan. Sebagai contoh, untuk mengirim ke Summarecon Mall Bekasi, rute tercepat adalah melalui Bekasi Cyber Park dengan total waktu 45 menit, meskipun ada jalan lain yang mungkin terlihat lebih dekat di peta."

---

## 4. Penjelasan Kode

**[NAMA 1]**

"Sekarang kami akan menjelaskan implementasi kode program kami yang ditulis dalam bahasa Go. Kami membagi kode ke dalam beberapa file untuk menjaga keterbacaan dan modularitas."

**(Slide 6: `graph/graph.go`)**

"File pertama adalah `graph.go`. File ini mendefinisikan struktur dasar dari graf."

*   **`type Edge struct`**: Ini adalah struct untuk merepresentasikan sebuah 'sisi'. `To` adalah ID dari simpul tujuan, dan `Weight` adalah bobotnya (waktu tempuh).
*   **`type Graph struct`**: Ini adalah struct utama graf. `Nodes` adalah sebuah map untuk menyimpan nama lokasi berdasarkan ID-nya. `Edges` adalah adjacency list, yaitu sebuah map yang memetakan ID simpul ke sebuah array (slice) dari `Edge`, yang berisi semua koneksi dari simpul tersebut.
*   **`NewGraph()`**: Fungsi ini hanya untuk inisialisasi, membuat objek graf baru yang masih kosong.
*   **`AddNode()`**: Metode untuk menambahkan simpul baru ke dalam graf.
*   **`AddEdge()`**: Metode ini menambahkan koneksi antara dua simpul `u` dan `v` dengan bobot `w`. Karena graf kita tidak berarah, kita menambahkan sisi dari `u` ke `v` dan juga dari `v` ke `u`.

**(Slide 7: `data/data.go`)**

**[NAMA 2]**

"Selanjutnya adalah file `data.go`. File ini bertanggung jawab untuk menyediakan data spesifik untuk kasus Momoyo."

*   **`const`**: Kami mendefinisikan konstanta untuk ID node sumber (`SumberNode = 0`) dan jumlah total node (`JumlahNode = 12`) agar kode lebih mudah dibaca.
*   **`DataSample()`**: Ini adalah fungsi utama di file ini.
    *   Pertama, ia membuat objek graf baru menggunakan `graph.NewGraph()`.
    *   Kemudian, ia memanggil `g.AddNode()` sebanyak 12 kali untuk menambahkan semua lokasi, dari Gudang Momoyo hingga Summarecon Mall Bekasi, sesuai dengan data yang kami kumpulkan.
    *   Terakhir, ia memanggil `g.AddEdge()` untuk menambahkan 20 koneksi jalan antar lokasi beserta estimasi waktu tempuhnya. Data ini merepresentasikan jaringan jalan yang ada.

**(Slide 8: `dijkstra/dijkstra.go` - Bagian 1: Priority Queue)**

**[NAMA 1]**

"File `dijkstra.go` adalah inti dari logika pencarian rute. Ini adalah implementasi dari Algoritma Dijkstra. Untuk efisiensi, algoritma ini memerlukan struktur data **Priority Queue**."

*   **`type Item struct`**: Mendefinisikan item yang akan disimpan di dalam Priority Queue. Isinya adalah `node` (ID simpul), `dist` (jarak dari sumber), dan `index` (untuk keperluan internal library `heap`).
*   **`type PriorityQueue []*Item`**: Mendefinisikan tipe Priority Queue sebagai slice dari pointer ke `Item`.
*   **`Len()`, `Less()`, `Swap()`, `Push()`, `Pop()`**: Ini adalah metode-metode yang wajib diimplementasikan agar tipe `PriorityQueue` kita memenuhi `heap.Interface` dari library standar Go. `Less` adalah bagian terpenting, karena ia menentukan bahwa antrian akan diurutkan berdasarkan jarak (`dist`) terkecil.

**(Slide 9: `dijkstra/dijkstra.go` - Bagian 2: Algoritma & Rekonstruksi)**

**[NAMA 2]**

"Masih di file `dijkstra.go`, berikut adalah fungsi utamanya."

*   **`Dijkstra(...)`**:
    1.  **Inisialisasi**: Membuat tiga slice: `dist` untuk menyimpan jarak terpendek (diisi dengan nilai 'tak hingga'), `prev` untuk melacak dari mana kita mencapai sebuah node (diisi -1), dan `visited` untuk menandai node yang sudah selesai diproses. Jarak ke node sumber diatur menjadi 0.
    2.  **Setup Priority Queue**: Priority Queue diinisialisasi dan node sumber langsung dimasukkan sebagai item pertama.
    3.  **Loop Utama**: Loop berjalan selama Priority Queue tidak kosong. Di setiap iterasi, ia mengambil (`Pop`) item dengan jarak terkecil.
    4.  **Relaksasi**: Untuk node yang baru saja diambil, ia memeriksa semua tetangganya (`g.Edges`). Jika ditemukan jalur alternatif yang lebih pendek ke tetangga tersebut, jaraknya diperbarui (`dist[e.To] = alt`), `prev` di-update, dan tetangga tersebut dimasukkan (`Push`) ke dalam Priority Queue.
    5.  **Return**: Setelah loop selesai, fungsi mengembalikan `dist` dan `prev`.

*   **`ReconstructPath(...)`**: Fungsi ini menggunakan array `prev` untuk membangun kembali rute. Ia bekerja dengan menelusuri dari `target` ke belakang melalui `prev` hingga mencapai `source`, lalu membalik urutannya untuk mendapatkan jalur yang benar dari `source` ke `target`.

**(Slide 10: `main.go`)**

**[NAMA 1]**

"Terakhir, `main.go` adalah file yang menyatukan semuanya."

*   **`main()`**:
    1.  **`g := data.DataSample()`**: Langkah pertama adalah memanggil fungsi `DataSample` untuk mendapatkan graf yang sudah terisi data lokasi dan rute Momoyo.
    2.  **`dist, prev := dijkstra.Dijkstra(...)`**: Selanjutnya, memanggil fungsi `Dijkstra` dengan graf `g`, titik awal `data.SumberNode`, dan jumlah total node. Hasilnya disimpan dalam variabel `dist` dan `prev`.
    3.  **Print Results**: Bagian terakhir dari `main` adalah loop untuk mencetak hasil. Ia akan mengiterasi semua node. Untuk setiap node tujuan, ia memanggil `dijkstra.ReconstructPath` untuk membangun jalur, lalu memformat dan mencetak nama tujuan, total waktu tempuh dari `dist`, dan rute yang sudah diformat.

---

## 5. Kesimpulan

**(Slide 11: Kesimpulan Proyek)**

**[NAMA 2]**

"Dari proyek yang telah kami lakukan, dapat ditarik beberapa kesimpulan:"
1.  "Masalah optimasi rute pengiriman dapat dimodelkan secara efektif menggunakan struktur data graf, di mana lokasi adalah simpul dan jalan adalah sisi berbobot."
2.  "Algoritma Dijkstra terbukti menjadi solusi yang tepat dan efisien untuk menemukan jalur terpendek (tercepat) dari satu titik asal (gudang) ke berbagai titik tujuan (cabang mal)."
3.  "Dengan implementasi ini, 'Momoyo' dapat menghemat waktu dan biaya operasional dengan selalu menggunakan rute pengiriman yang paling optimal."

**(Slide 12: Kesimpulan Perkuliahan & Penutup)**

**[NAMA 1]**

"Selama mengikuti perkuliahan Matematika Diskrit, kami telah mempelajari banyak konsep fundamental yang ternyata sangat aplikatif di dunia nyata dan ilmu komputer. Mata kuliah ini membuka wawasan kami tentang bagaimana logika, teori himpunan, graf, dan berbagai konsep lainnya menjadi fondasi dari solusi-solusi komputasi yang kompleks, seperti masalah optimasi rute yang kami selesaikan dalam proyek ini."

"Kami menyadari bahwa pemahaman yang kuat tentang matematika diskrit adalah aset yang sangat berharga bagi seorang insinyur perangkat lunak."

**[NAMA 2]**

"Sekian presentasi dari kami. Terima kasih atas perhatian Bapak/Ibu dosen. Selanjutnya kami membuka sesi tanya jawab jika ada."

---

Semoga draf ini membantu! Saya akan menjalankan kode sekarang untuk mengisi bagian hasil yang kosong.
