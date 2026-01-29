# Audit Ringkasan Script Video Presentasi

## Pendahuluan
Studi kasus ini mengangkat permasalahan distribusi Ice Cream Momoyo di wilayah Jabodetabek yang sensitif terhadap waktu. Distribusi dilakukan dari satu warehouse pusat di Cibubur ke 11 outlet di berbagai lokasi (Mall Kelapa Gading, Grand Indonesia, dll). Tantangan utama adalah menentukan rute tercepat dan paling efisien untuk menjaga mutu produk. Tujuan penelitian ini adalah merancang representasi graf, mengimplementasikan algoritma Dijkstra, serta menganalisis efektivitasnya dalam menentukan rute optimal.

## Metode
Penelitian menggunakan konsep **Teori Graf** jenis *Weighted Undirected Graph* dengan 12 node (lokasi) dan 20 edge (jalan). Bobot (*weight*) pada edge merepresentasikan waktu tempuh dalam menit. Algoritma yang digunakan adalah **Algoritma Dijkstra** untuk menyelesaikan *Single-Source Shortest Path Problem*. Algoritma ini bekerja dengan menginisialisasi jarak, memilih node terpendek, melakukan relaksasi (update jarak tetangga), dan menandai node yang sudah dikunjungi.

## Experimental dan Results
Implementasi dilakukan menggunakan bahasa pemrograman **Go (Golang)** karena performa dan dukungan library yang baik. Struktur kode terdiri dari `main.go`, `graph.go`, `dijkstra.go` (menggunakan Priority Queue), dan `data.go`.
**Hasil Eksperimen:**
- Rute tercepat: Cibinong City Mall (15 menit).
- Rute terlama: BSD City (135 menit).
- Waktu eksekusi algoritma: < 1 milidetik.
- Memory usage: ~2 KB.
- Algoritma berhasil memetakan rute terpendek dari Warehouse ke 11 lokasi lainnya.

## Discussion
Analisis menunjukkan pola geografis dimana lokasi di sisi timur (Cibinong, Bekasi, Depok) memiliki waktu tempuh lebih singkat dibanding sisi barat (Tangerang, BSD). Teridentifikasi *Hub Nodes* penting seperti Grand Indonesia dan Depok Town Square yang sering dilewati rute optimal. Algoritma Dijkstra terbukti lebih efisien untuk kasus *single-source* dengan bobot positif dibandingkan Bellman-Ford (karena kompleksitas waktu) atau Floyd-Warshall.

## Kesimpulan
Sistem berhasil merepresentasikan jaringan distribusi sebagai graf dan mengimplementasikan algoritma Dijkstra dengan sukses. Sistem mampu menentukan rute optimal dengan performa tinggi (cepat dan hemat memori). Penerapan sistem ini berpotensi mengurangi waktu pengiriman hingga 20-30%, mengurangi tingkat pelelehan produk, dan meningkatkan efisiensi operasional harian.

## Future Work
**Jangka Pendek:** Integrasi dengan Google Maps API untuk data trafik *real-time*, penambahan lokasi, dan pembuatan dashboard interaktif.
**Jangka Panjang:** Penyesuaian bobot dinamis (*dynamic weight*) berdasarkan kondisi lalu lintas, optimasi multi-kriteria (waktu, biaya, jarak), penerapan Machine Learning untuk prediksi trafik, serta pengembangan ke arah *Multi-Vehicle Routing Problem* (VRP) dan aplikasi mobile.
