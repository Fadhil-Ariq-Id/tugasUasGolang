# Algoritma Dijkstra (Pseudocode)

Diberikan sebuah Graf $G$ dan titik awal (*source*) $s$.

```text
DIJKSTRA(Graf, source):
    // 1. Inisialisasi
    untuk setiap node v dalam Graf:
        jarak[v] = TAK_HINGGA
        sebelumnya[v] = UNDEFINED
    
    jarak[source] = 0
    Q = Antrean Prioritas yang berisi semua node dengan nilai jarak[v]

    // 2. Loop Utama
    selama Q tidak kosong:
        u = node dalam Q dengan jarak[u] terkecil
        hapus u dari Q

        // 3. Relaksasi Tetangga
        untuk setiap tetangga v dari u:
            // Hitung alternatif jarak melalui u
            alt = jarak[u] + bobot(u, v)
            
            jika alt < jarak[v]:
                jarak[v] = alt
                sebelumnya[v] = u
                // Perbarui posisi v dalam Antrean Prioritas
                Q.perbarui_prioritas(v, alt)

    kembalikan jarak, sebelumnya
```
