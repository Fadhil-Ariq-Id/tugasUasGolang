# Script Video Presentasi (Versi Ringkas 5-7 Menit)
**Fokus: Efisiensi & Inti Kode**

---

## 0:00 - 0:15 | Pendahuluan (15 Detik)
**[Visual: Cover Slide -> Peta Distribusi]**

"Assalamualaikum. Kami dari Kelompok 3 akan mempresentasikan optimasi rute pengiriman Ice Cream Momoyo di Jabodetabek.
Tantangan utamanya adalah produk yang sensitif suhu, sehingga rute pengiriman harus secepat mungkin. Solusi kami adalah membangun sistem pencarian rute terpendek dari Warehouse Cibubur ke 11 outlet menggunakan **Algoritma Dijkstra** berbasis **Golang** untuk performa maksimal."

---

## 0:15 - 5:15 | Method, Experiment & Result (5 Menit)
**[Visual: Diagram Graf -> Code Editor (graph.go)]**

"(Presenter Teknis)
Mari bedah inti teknisnya. Kami memodelkan peta ini sebagai **Weighted Undirected Graph** dimana Node adalah lokasi dan Bobot adalah waktu tempuh.

Masuk ke **Implementasi Code**.
Pertama, efisiensi struktur data. Di `graph.go`, kami tidak menggunakan *Adjacency Matrix* yang boros memori `O(V^2)`, melainkan **Adjacency List** menggunakan map:
`Edges map[int][]Edge`.
Ini membuat iterasi tetangga saat relaksasi sangat efisien hanya pada node yang terhubung saja.

**[Visual: Code dijkstra.go (Struct Item & PriorityQueue)]**

Kunci kecepatan algoritma Dijkstra ada di antrian prioritas. Jika pakai array biasa, pencarian node minimum butuh `O(V)`.
Kami mengimplementasikan **Priority Queue** menggunakan **Min-Heap** di `dijkstra.go`.
Lihat struct `PriorityQueue` ini. Kami mengimplementasikan interface `heap.Interface` bawaan Go.
Method `Less` kami set `pq[i].dist < pq[j].dist`, memastikan node dengan jarak terpendek selalu ada di root heap (index 0).
Ini menjamin operasi pengambilan node terdekat berjalan dalam kompleksitas Logaritmik `O(log V)`.

**[Visual: Code dijkstra.go (Fungsi Dijkstra)]**

Ini adalah *Core Logic* fungsi `Dijkstra`:
1.  **Inisialisasi**: Kami siapkan array `dist` dengan nilai `MaxInt` (Infinity), dan set source ke 0.
2.  **Main Loop**:
    `for pq.Len() > 0`
    Kami lakukan `heap.Pop`. Ini operasi yang sangat murah secara komputasi.
    Kami tambahkan check `if visited[u.node] continue` untuk menghindari pemrosesan ulang node yang sudah optimal. Ini *cutting* proses yang tidak perlu.
3.  **Relaxation Process** (Inti Algoritma):
    Kami loop semua tetangga (`range g.Edges`).
    Logikanya: `if alt < dist[e.To]`.
    Jika jalur via node sekarang lebih cepat dari jalur lama, kami update nilai `dist` dan simpan *path*-nya di array `prev`.
    Lalu **Push** ke heap. Dengan pendekatan *Greedy* ini, kita tidak perlu mengecek semua kombinasi jalan, cukup yang menjanjikan saja.

**[Visual: Terminal Output / Demo Program]**

**Result & Eksperimen**:
Kami menguji sistem dengan 12 Node tersebar (Jakarta-Depok-Bogor-Bekasi-Tangerang).
Saat dijalankan di `main.go`, algoritma memberi hasil instan.
Output menunjukkan:
- **Tercepat**: 15 menit ke Cibinong (Direct connection).
- **Terjauh**: 135 menit ke BSD City. Perhatikan grafiknya, algoritma cerdas memilih rute memutar lewat Depok dan Grand Indonesia karena total waktunya lebih kecil dibanding jalur alternatif yang macet.

**Performance Audit**:
- Waktu Eksekusi: **< 1 milidetik** (mikrodetik).
- Space Complexity: `O(V)` karena Adjacency List.
- Time Complexity: Stabil di `O((V+E) log V)`.
Ini membuktikan implementasi Go dan Min-Heap kami sangat efisien untuk skala produksi."

---

## 5:15 - 5:25 | Discussion (10 Detik)
**[Visual: Grafik Analisis]**

"Hasil ini mengkonfirmasi pola geografis: lokasi Timur (Cibinong/Bekasi) jauh lebih cepat dijangkau dibanding Barat. Secara algoritma, Dijkstra menang telak dibanding Bellman-Ford untuk graf bobot positif karena tidak perlu melakukan relaksasi berulang pada semua edge (V-1 kali)."

---

## 5:25 - 5:35 | Kesimpulan (10 Detik)
**[Visual: Summary Point]**

"Kesimpulannya, sistem simulasi rute Momoyo ini **berhasil 100%**. Kami memvalidasi bahwa kombinasi Golang dan Priority Queue menghasilkan pencarian rute yang akurat, scalable, dan sangat ringan dijalankan di server dengan resource terbatas sekalipun."

---

## 5:35 - 5:50 | Solutions & Future Work (15 Detik)
**[Visual: Future Roadmap]**

"Sebagai solusi jangka panjang, kami akan mengintegrasikan **Google Maps API** agar bobot waktu bersifat *real-time* mengikuti kemacetan, bukan statis. Kami juga akan mengembangkan ini menjadi **Multi-Vehicle Routing** untuk manajemen armada truk skala besar. Sekian, terima kasih."
