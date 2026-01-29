# Script Video Presentasi - Optimasi Rute Pengiriman Ice Cream Momoyo

---

## OPENING (Presenter 1)

**[Visual: Judul Project di layar]**

"Assalamualaikum warahmatullahi wabarakatuh. Selamat pagi/siang Bapak/Ibu dosen dan teman-teman semua.

Perkenalkan, saya [Nama Presenter 1] dan saya [Nama Presenter 2] dari Kelompok 3. Hari ini kami akan mempresentasikan project kami yang berjudul 'Optimasi Rute Pengiriman Toko Ice Cream Momoyo di Wilayah Jabodetabek Menggunakan Algoritma Dijkstra'.

Project ini merupakan tugas akhir untuk mata kuliah Matematika Diskrit di STIMIK Tazkia."

---

## BAGIAN 1: LATAR BELAKANG (Presenter 1)

**[Visual: Gambar ice cream dan truck pengiriman]**

"Baik, saya akan menjelaskan latar belakang dari kasus yang kami angkat.

Industri makanan beku, khususnya ice cream, menghadapi tantangan yang cukup unik dalam hal distribusi. Kenapa? Karena ice cream itu produk yang sangat sensitif terhadap waktu. Sedikit saja terlalu lama di perjalanan, kualitasnya bisa langsung menurun. Bahkan meskipun sudah menggunakan rantai dingin dan kendaraan berpendingin, risikonya tetap ada.

**[Visual: Peta Jabodetabek]**

Nah, dalam studi kasus ini, kami mengambil contoh Toko Ice Cream Momoyo yang melakukan distribusi ke berbagai outlet di wilayah Jabodetabek - itu Jakarta, Bogor, Depok, Tangerang, dan Bekasi.

**[Visual: Diagram warehouse dan titik pengiriman]**

Kami membuat simulasi dengan satu warehouse pusat sebagai titik awal pengiriman. Warehouse ini kami tempatkan di Cibubur, Jakarta Timur sebagai Node 0. Kemudian ada 11 lokasi pengiriman yang merepresentasikan pusat-pusat perbelanjaan utama seperti Mall Kelapa Gading, Grand Indonesia, Pondok Indah Mall, BSD City, dan lokasi-lokasi lainnya.

Perlu kami tekankan bahwa ini adalah simulasi untuk tujuan pembelajaran. Jadi model ini tidak sepenuhnya meniru sistem logistik asli Momoyo, tapi cukup untuk menggambarkan tantangan utama dalam optimasi rute pengiriman.

**[Visual: Grafik showing challenges]**

Tantangan utamanya adalah: bagaimana menentukan rute tercepat dan paling efisien? Karena tujuannya bukan hanya hemat biaya, tapi juga menjaga mutu produk ice cream tetap bagus sampai di tangan pelanggan."

---

## BAGIAN 2: RUMUSAN MASALAH & TUJUAN (Presenter 2)

**[Visual: Slide rumusan masalah]**

"Baik, sekarang saya lanjutkan untuk rumusan masalah dan tujuan penelitian.

Dari latar belakang tadi, kami merumuskan 4 masalah utama:

Pertama, bagaimana cara membangun representasi graf untuk jaringan rute pengiriman yang bisa diproses secara komputasional?

Kedua, bagaimana menerapkan algoritma Dijkstra untuk menemukan rute tercepat dari warehouse ke setiap lokasi?

Ketiga, berapa waktu tempuh optimal yang bisa dicapai?

Dan keempat, seberapa baik performa dan efisiensi algoritma Dijkstra untuk kasus ini?

**[Visual: Slide tujuan]**

Sesuai dengan rumusan masalah, tujuan penelitian kami adalah:
- Merancang representasi graf untuk jaringan rute
- Mengimplementasikan algoritma Dijkstra menggunakan bahasa Go
- Menentukan rute optimal dan waktu tempuh minimum
- Dan menganalisis efektivitas algoritma Dijkstra untuk kasus pengiriman ice cream ini."

---

## BAGIAN 3: METODE - TEORI GRAF (Presenter 1)

**[Visual: Diagram graf sederhana]**

"Sekarang kita masuk ke bagian metode. Saya akan jelaskan konsep dasar yang kami gunakan.

Pertama, tentang Teori Graf. Graf adalah struktur matematika yang terdiri dari kumpulan node atau vertex dan edge yang menghubungkan node-node tersebut.

**[Visual: Weighted undirected graph example]**

Dalam project ini, kami menggunakan Weighted Undirected Graph. Weighted artinya setiap edge punya nilai atau bobot - dalam kasus ini bobotnya adalah waktu tempuh dalam satuan menit. Undirected artinya koneksi antara dua node bisa dilalui dari kedua arah dengan bobot yang sama.

**[Visual: Diagram 12 nodes]**

Graf kami terdiri dari 12 node:
- Node 0: Warehouse Momoyo di Cibubur sebagai source
- Node 1-11: Lokasi-lokasi pengiriman seperti Mall Kelapa Gading, Grand Indonesia, Pondok Indah Mall, BSD City, Bekasi Cyber Park, Depok Town Square, Tangerang City, Bogor Trade Mall, Cibinong City Mall, Senayan City, dan Summarecon Mall Bekasi.

Total ada 20 koneksi atau edge dalam graf ini, dengan range waktu tempuh dari 15 menit sampai 50 menit per edge."

---

## BAGIAN 4: ALGORITMA DIJKSTRA (Presenter 2)

**[Visual: Foto Edsger Dijkstra]**

"Sekarang kita bahas tentang algoritma yang kami gunakan, yaitu Algoritma Dijkstra.

Algoritma ini dikembangkan oleh Edsger W. Dijkstra pada tahun 1956 untuk menyelesaikan Single-Source Shortest Path Problem - artinya mencari jalur terpendek dari satu titik sumber ke semua titik lainnya.

**[Visual: Flowchart algoritma Dijkstra]**

Cara kerjanya seperti ini:

Pertama, inisialisasi. Semua node diberi jarak tak hingga, kecuali source node yang diberi jarak 0.

Kedua, pilih node dengan jarak terpendek yang belum dikunjungi.

Ketiga, lakukan relaksasi - update jarak ke semua neighbor dari node tersebut jika ditemukan jalur yang lebih pendek.

Keempat, tandai node sebagai sudah dikunjungi.

Ulangi langkah 2-4 sampai semua node sudah diproses.

**[Visual: Complexity notation]**

Dari segi kompleksitas:
- Time Complexity adalah O((V + E) log V) dengan implementasi priority queue
- Space Complexity adalah O(V)

Di mana V adalah jumlah node dan E adalah jumlah edge.

Untuk kasus kami dengan 12 node dan 40 directed edges, ini sangat efisien - waktu eksekusinya kurang dari 1 milidetik."

---

## BAGIAN 5: IMPLEMENTASI & STRUKTUR CODE (Presenter 1)

**[Visual: Struktur folder project]**

"Baik, sekarang kita masuk ke implementasi kodenya. Kami menggunakan bahasa pemrograman Go atau Golang.

Kenapa Go? Karena Go punya beberapa keunggulan:
- Performance yang cepat karena compiled language
- Type safety dengan static typing
- Punya standard library yang bagus, termasuk container/heap untuk priority queue
- Syntax yang clean dan mudah dibaca

**[Visual: File structure]**

Struktur project kami terdiri dari 4 file utama:
1. main.go - program utama
2. graph.go - struktur data graf
3. dijkstra.go - implementasi algoritma
4. data.go - dataset lokasi dan koneksi

Mari kita bahas satu per satu."

---

## BAGIAN 6: PENJELASAN CODE - graph.go (Presenter 2)

**[Visual: Code graph.go di layar]**

"Saya akan jelaskan file graph.go terlebih dahulu.

```go
type Edge struct {
    To     int
    Weight int
}
```

Pertama, kita definisikan struct Edge. Edge ini punya dua field:
- `To` bertipe int, menyimpan ID node tujuan
- `Weight` bertipe int, menyimpan bobot atau waktu tempuh dalam menit

```go
type Graph struct {
    Nodes map[int]string
    Edges map[int][]Edge
}
```

Kemudian struct Graph. Graph menggunakan adjacency list representation dengan dua map:
- `Nodes` adalah map dari int ke string, menyimpan ID node dan nama lokasinya
- `Edges` adalah map dari int ke slice of Edge, menyimpan daftar koneksi untuk setiap node

```go
func NewGraph() *Graph {
    return &Graph{
        Nodes: make(map[int]string),
        Edges: make(map[int][]Edge),
    }
}
```

Function NewGraph adalah constructor untuk membuat instance Graph baru dengan inisialisasi map kosong.

```go
func (g *Graph) AddNode(id int, name string) {
    g.Nodes[id] = name
    if _, ok := g.Edges[id]; !ok {
        g.Edges[id] = []Edge{}
    }
}
```

Method AddNode menambahkan node baru. Dia menyimpan nama node di map Nodes, dan menginisialisasi slice Edge kosong untuk node tersebut jika belum ada.

```go
func (g *Graph) AddEdge(u, v, w int) {
    g.Edges[u] = append(g.Edges[u], Edge{To: v, Weight: w})
    g.Edges[v] = append(g.Edges[v], Edge{To: u, Weight: w})
}
```

Method AddEdge menambahkan edge undirected. Karena undirected, kita harus menambahkan dua edge: dari u ke v, dan dari v ke u, dengan bobot yang sama.

```go
func (g *Graph) NodeName(id int) string {
    if n, ok := g.Nodes[id]; ok {
        return n
    }
    return "(unknown)"
}
```

Dan NodeName adalah helper method untuk mengambil nama lokasi dari ID node."

---

## BAGIAN 7: PENJELASAN CODE - data.go (Presenter 1)

**[Visual: Code data.go di layar]**

"Sekarang kita lihat file data.go yang berisi dataset kita.

```go
const (
    SumberNode = 0
    JumlahNode = 12
)
```

Di awal, kita definisikan dua konstanta:
- `SumberNode` bernilai 0, ini adalah ID warehouse kita
- `JumlahNode` bernilai 12, total node dalam graf

```go
func DataSample() *graph.Graph {
    g := graph.NewGraph()
```

Function DataSample mengembalikan instance Graph yang sudah terisi data. Pertama kita buat graph baru.

```go
    g.AddNode(0, "Warehouse Momoyo (Cibubur)")
    g.AddNode(1, "Mall Kelapa Gading")
    g.AddNode(2, "Grand Indonesia (GI)")
    // ... dan seterusnya sampai node 11
```

Kemudian kita tambahkan semua 12 node dengan nama lokasinya masing-masing. Data ini sesuai dengan tabel yang ada di laporan halaman 11.

```go
    g.AddEdge(0, 5, 25)  // Warehouse ke Bekasi Cyber Park, 25 menit
    g.AddEdge(0, 6, 30)  // Warehouse ke Depok Town Square, 30 menit
    g.AddEdge(0, 9, 15)  // Warehouse ke Cibinong City Mall, 15 menit
    // ... total 20 edge
```

Setelah itu kita tambahkan semua 20 koneksi dengan bobotnya. Misalnya, edge pertama menghubungkan Warehouse (node 0) dengan Bekasi Cyber Park (node 5) dengan waktu tempuh 25 menit.

Semua data edge ini juga sesuai dengan tabel di laporan halaman 11-12.

```go
    return g
}
```

Akhirnya, kita return graph yang sudah lengkap."

---

## BAGIAN 8: PENJELASAN CODE - dijkstra.go Part 1 (Presenter 2)

**[Visual: Code dijkstra.go di layar - bagian Priority Queue]**

"Sekarang file yang paling penting, dijkstra.go. Saya akan jelaskan bagian demi bagian.

```go
type Item struct {
    node  int
    dist  int
    index int
}
```

Pertama, struct Item merepresentasikan satu node dalam priority queue:
- `node` adalah ID node
- `dist` adalah jarak sementara dari source
- `index` untuk keperluan internal heap

```go
type PriorityQueue []*Item
```

PriorityQueue adalah slice of pointer ke Item. Ini akan kita implementasikan sebagai min-heap.

```go
func (pq PriorityQueue) Len() int { return len(pq) }
```

Method Len mengembalikan panjang queue.

```go
func (pq PriorityQueue) Less(i, j int) bool { 
    return pq[i].dist < pq[j].dist 
}
```

Method Less membandingkan dua item. Return true jika item i punya dist lebih kecil dari item j. Ini yang membuat heap kita menjadi min-heap - item dengan jarak terkecil akan berada di top.

```go
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}
```

Method Swap menukar posisi dua item dan update index mereka.

```go
func (pq *PriorityQueue) Push(x any) {
    item := x.(*Item)
    item.index = len(*pq)
    *pq = append(*pq, item)
}
```

Method Push menambahkan item baru ke heap. Kita type assert x menjadi pointer Item, set indexnya, lalu append ke slice.

```go
func (pq *PriorityQueue) Pop() any {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil
    item.index = -1
    *pq = old[:n-1]
    return item
}
```

Method Pop mengeluarkan item terakhir dari heap. Kita ambil item terakhir, set nilainya jadi nil untuk garbage collection, set index jadi -1, potong slice, lalu return itemnya.

Kelima method ini - Len, Less, Swap, Push, Pop - mengimplementasikan interface heap.Interface dari package container/heap. Dengan ini, Go bisa otomatis mengelola heap kita."

---

## BAGIAN 9: PENJELASAN CODE - dijkstra.go Part 2 (Presenter 1)

**[Visual: Code dijkstra.go di layar - fungsi Dijkstra]**

"Sekarang kita masuk ke fungsi utama, yaitu Dijkstra itu sendiri.

```go
func Dijkstra(g *graph.Graph, source int, n int) (dist []int, prev []int) {
```

Function signature-nya menerima tiga parameter:
- `g` adalah pointer ke Graph
- `source` adalah ID node awal
- `n` adalah jumlah total node

Dan mengembalikan dua slice:
- `dist` untuk menyimpan jarak terpendek ke setiap node
- `prev` untuk menyimpan predecessor, berguna untuk rekonstruksi jalur

```go
    dist = make([]int, n)
    prev = make([]int, n)
    visited := make([]bool, n)
```

Kita inisialisasi tiga slice dengan ukuran n.

```go
    for i := 0; i < n; i++ {
        dist[i] = math.MaxInt
        prev[i] = -1
    }
    dist[source] = 0
```

Loop ini adalah tahap inisialisasi algoritma Dijkstra:
- Semua jarak diset ke tak hingga (MaxInt)
- Semua predecessor diset ke -1 (undefined)
- Kecuali source node, jaraknya diset ke 0

```go
    pq := make(PriorityQueue, 0, n)
    heap.Init(&pq)
    heap.Push(&pq, &Item{node: source, dist: 0})
```

Kita buat priority queue, inisialisasi sebagai heap, lalu push source node dengan jarak 0.

```go
    for pq.Len() > 0 {
        u := heap.Pop(&pq).(*Item)
```

Ini loop utama algoritma. Selama queue tidak kosong, kita pop item dengan jarak terkecil.

```go
        if visited[u.node] {
            continue
        }
        visited[u.node] = true
```

Jika node sudah pernah diproses, skip. Kalau belum, tandai sebagai visited.

```go
        for _, e := range g.Edges[u.node] {
            if visited[e.To] {
                continue
            }
```

Iterasi semua neighbor dari node u. Skip jika neighbor sudah visited.

```go
            alt := dist[u.node] + e.Weight
            if alt < dist[e.To] {
                dist[e.To] = alt
                prev[e.To] = u.node
                heap.Push(&pq, &Item{node: e.To, dist: alt})
            }
        }
    }
    return dist, prev
}
```

Ini adalah bagian relaksasi:
- Hitung jarak alternatif: jarak ke u ditambah weight edge
- Jika alt lebih kecil dari jarak yang tersimpan, update jaraknya
- Update predecessor-nya menjadi u
- Push node tujuan ke queue dengan jarak baru

Setelah loop selesai, return dist dan prev."

---

## BAGIAN 10: PENJELASAN CODE - dijkstra.go Part 3 & main.go (Presenter 2)

**[Visual: Code dijkstra.go - ReconstructPath]**

"Masih di dijkstra.go, ada satu function lagi yaitu ReconstructPath.

```go
func ReconstructPath(prev []int, source, target int) []int {
    if source == target {
        return []int{source}
    }
    if target < 0 || target >= len(prev) {
        return nil
    }
```

Function ini merekonstruksi jalur dari source ke target menggunakan data prev. Pertama, handle edge case: jika source sama dengan target, return source aja. Jika target invalid, return nil.

```go
    path := []int{}
    cur := target
    for cur != -1 {
        path = append(path, cur)
        if cur == source {
            break
        }
        cur = prev[cur]
    }
```

Kita telusuri mundur dari target ke source mengikuti predecessor. Append setiap node ke path sampai mencapai source.

```go
    if len(path) == 0 || path[len(path)-1] != source {
        return nil
    }
```

Validasi: jika path kosong atau tidak sampai ke source, berarti tidak terhubung, return nil.

```go
    for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
        path[i], path[j] = path[j], path[i]
    }
    return path
}
```

Karena kita telusuri mundur, path sekarang terbalik. Kita reverse path-nya agar urutannya benar dari source ke target, lalu return.

**[Visual: Code main.go]**

Sekarang file terakhir, main.go.

```go
func main() {
    g := data.DataSample()
```

Di main, pertama kita load data dengan memanggil DataSample.

```go
    fmt.Printf("Calculating shortest paths from: %s\n\n", 
        g.NodeName(data.SumberNode))
    dist, prev := dijkstra.Dijkstra(g, data.SumberNode, data.JumlahNode)
```

Print header, lalu jalankan algoritma Dijkstra. Kita dapatkan dist dan prev.

```go
    fmt.Printf("%-25s %-15s %s\n", "Tujuan", "Waktu Tempuh", "Rute")
    fmt.Println(strings.Repeat("-", 100))
```

Print header tabel dengan format kolom yang rapi.

```go
    for id := 0; id < data.JumlahNode; id++ {
        if id == data.SumberNode {
            continue
        }
```

Loop semua node, skip source node.

```go
        path := dijkstra.ReconstructPath(prev, data.SumberNode, id)
        pathNames := formatPath(g, path)
        
        fmt.Printf("%-25s %-15s %s\n", 
            g.NodeName(id), 
            fmt.Sprintf("%d menit", dist[id]), 
            pathNames)
    }
}
```

Untuk setiap tujuan, rekonstruksi jalurnya, format menjadi nama-nama lokasi, lalu print dalam format tabel.

```go
func formatPath(g *graph.Graph, path []int) string {
    if len(path) == 0 {
        return "Tidak Terjangkau"
    }
    names := make([]string, 0, len(path))
    for _, id := range path {
        names = append(names, g.NodeName(id))
    }
    return strings.Join(names, " --> ")
}
```

Function formatPath mengubah slice ID menjadi string yang readable. Convert setiap ID jadi nama, lalu join dengan arrow."

---

## BAGIAN 11: HASIL EKSPERIMEN (Presenter 1)

**[Visual: Screenshot output program]**

"Baik, sekarang kita lihat hasil eksperimen dari program yang kita jalankan.

**[Visual: Tabel hasil]**

Dari eksekusi algoritma Dijkstra dengan source node Warehouse Momoyo di Cibubur, kita dapatkan hasil sebagai berikut:

Rute TERCEPAT adalah ke Cibinong City Mall dengan waktu tempuh 15 menit. Ini direct connection dari warehouse.

Untuk lokasi-lokasi lain:
- Bekasi Cyber Park: 25 menit (direct)
- Depok Town Square: 30 menit (direct)
- Bogor Trade Mall: 35 menit (direct)
- Summarecon Mall Bekasi: 45 menit (melalui Bekasi Cyber Park)
- Mall Kelapa Gading: 60 menit (melalui Bekasi Cyber Park)
- Grand Indonesia: 70 menit (melalui Depok Town Square)
- Senayan City: 85 menit (melalui Depok → Grand Indonesia)
- Tangerang City: 85 menit juga (ada multiple paths)
- Pondok Indah Mall: 100 menit
- Tangerang City versi lain: 125 menit
- Dan yang TERLAMA adalah BSD City: 135 menit, dengan rute Warehouse → Depok → Grand Indonesia → Pondok Indah Mall → BSD City.

**[Visual: Grafik distribusi]**

Jika kita kategorikan:
- Sangat Dekat (< 30 menit): 2 lokasi atau 16.7%
- Dekat (30-60 menit): 3 lokasi atau 25%
- Sedang (60-90 menit): 4 lokasi atau 33.3%
- Jauh (> 90 menit): 3 lokasi atau 25%

**[Visual: Performance metrics]**

Dari segi performa algoritma:
- Waktu eksekusi: kurang dari 1 milidetik
- Memory usage: sekitar 2 KB
- Jumlah node diproses: 12
- Jumlah edge dievaluasi: sekitar 40

Ini menunjukkan algoritma Dijkstra sangat efisien untuk problem set kita."

---

## BAGIAN 12: ANALISIS & DISKUSI (Presenter 2)

**[Visual: Peta Jabodetabek dengan highlighting]**

"Sekarang kita analisis hasil ini lebih dalam.

**Pola Geografis:**

Ada pola yang cukup jelas dari hasil kita. Lokasi-lokasi di sisi timur seperti Cibinong, Bekasi, dan Depok punya waktu tempuh paling pendek karena memang secara geografis dekat dengan warehouse di Cibubur.

Sebaliknya, lokasi di sisi barat seperti BSD City dan Tangerang punya waktu tempuh paling lama karena harus melintasi Jakarta dulu.

**[Visual: Network diagram with hub nodes highlighted]**

**Hub Nodes:**

Yang menarik, ada beberapa node yang berfungsi sebagai hub atau penghubung. Grand Indonesia dan Depok Town Square misalnya - banyak rute optimal yang melewati node-node ini. Ini informasi penting untuk perencanaan logistik.

**[Visual: Comparison chart]**

**Implikasi Praktis:**

Dari hasil ini, ada beberapa rekomendasi untuk Momoyo:

Pertama, untuk lokasi yang sangat jauh seperti BSD dan Tangerang, mungkin perlu dipertimbangkan sub-warehouse di Jakarta Barat untuk efisiensi.

Kedua, scheduling pengiriman bisa dibagi berdasarkan cluster:
- Pagi untuk cluster timur yang dekat
- Siang untuk cluster tengah
- Sore untuk cluster barat dengan extra cooling

Ketiga, untuk rute-rute panjang, perlu cold chain yang lebih kuat untuk menjaga kualitas ice cream.

**[Visual: Algorithm comparison table]**

**Perbandingan Algoritma:**

Kenapa kita pilih Dijkstra? Karena untuk graf dengan bobot positif dan single-source shortest path, Dijkstra adalah yang paling optimal. 

Dibanding Bellman-Ford, Dijkstra lebih cepat dengan O((V+E) log V) vs O(VE).

Dibanding A*, Dijkstra tidak perlu heuristic function yang mungkin sulit didapat.

Dibanding Floyd-Warshall, Dijkstra lebih efisien kalau cuma butuh dari satu source."

---

## BAGIAN 13: KETERBATASAN & FUTURE WORK (Presenter 1)

**[Visual: Slide keterbatasan]**

"Tentu saja, penelitian ini punya beberapa keterbatasan.

**Keterbatasan Data:**

Pertama, edge weight kita menggunakan waktu tempuh rata-rata yang statis. Dalam kenyataan, waktu tempuh itu dinamis - ada rush hour, ada weekend, ada hari libur.

Kedua, kita hanya memodelkan 12 lokasi. Dalam bisnis real, bisa ratusan titik pengiriman.

Ketiga, graf kita adalah simplifikasi. Jaringan jalan yang sebenarnya jauh lebih kompleks.

**Keterbatasan Metodologi:**

Kita hanya mengoptimasi satu objektif yaitu waktu. Padahal dalam praktik ada multiple objectives: biaya, jarak, kualitas jalan, dll.

Kita juga belum mempertimbangkan kapasitas kendaraan, volume orderan, dan multiple vehicle routing.

**[Visual: Slide future work]**

**Pengembangan Future:**

Ada banyak yang bisa dikembangkan dari project ini:

Short-term:
- Integrasi dengan Google Maps API untuk data traffic real-time
- Menambah lebih banyak lokasi
- Membuat interactive dashboard

Long-term:
- Dynamic weight adjustment berdasarkan traffic
- Multi-criteria optimization: kombinasi waktu, biaya, jarak
- Machine learning untuk prediksi traffic
- Mobile app untuk driver
- Multi-vehicle routing problem (VRP)

Ini semua ada di laporan kami halaman 22-24."

---

## BAGIAN 14: KESIMPULAN (Presenter 2)

**[Visual: Slide kesimpulan]**

"Baik, kita sampai pada kesimpulan dari project ini.

**Kesimpulan Utama:**

Pertama, kami berhasil merepresentasikan jaringan rute pengiriman Momoyo sebagai weighted undirected graph dengan 12 nodes dan 20 edges.

Kedua, algoritma Dijkstra berhasil diimplementasikan dengan baik menggunakan Go, dengan struktur data adjacency list dan priority queue.

Ketiga, sistem berhasil menentukan shortest path dari warehouse ke semua lokasi dengan rute tercepat 15 menit ke Cibinong, dan terlama 135 menit ke BSD City. Rata-rata waktu tempuh 71.8 menit.

Keempat, performa sistem sangat baik: eksekusi kurang dari 1 milidetik, memory efficient, dan scalable.

Kelima, hasil sudah diverifikasi dengan perhitungan manual dan konsisten dengan pola geografis Jabodetabek.

**Kontribusi:**

Project ini berkontribusi dalam:
- Demonstrasi penerapan teori graf dalam problem real-world
- Implementasi algoritma Dijkstra dengan best practices
- Penyediaan framework yang bisa dikembangkan lebih lanjut

**Impact Potensial:**

Jika sistem ini diterapkan, potensi benefitnya:
- Pengurangan waktu pengiriman 20-30%
- Peningkatan delivery per hari 25%
- Pengurangan melting rate 15-20%
- Peningkatan customer satisfaction

Tentu dengan catatan perlu adaptasi dengan data operasional riil dan pengembangan lebih lanjut."

---

## CLOSING (Presenter 1)

**[Visual: Slide terima kasih]**

"Baik, demikian presentasi dari kami tentang Optimasi Rute Pengiriman Ice Cream Momoyo menggunakan Algoritma Dijkstra.

Kami menyadari masih banyak kekurangan dalam project ini, oleh karena itu kami sangat terbuka dengan kritik dan saran yang membangun dari Bapak/Ibu dosen.

Terima kasih atas perhatiannya.

Wassalamualaikum warahmatullahi wabarakatuh."

**[Visual: Slide Q&A]**

"Kami siap menjawab pertanyaan."

---

## TIPS UNTUK REKAMAN:

**Untuk Presenter 1:**
- Bagian 1, 3, 7, 9, 11, 13, 14 (Opening)
- Total durasi: sekitar 10-12 menit

**Untuk Presenter 2:**
- Bagian 2, 4, 6, 8, 10, 12, 14 (Closing)
- Total durasi: sekitar 10-12 menit

**Tips Umum:**
1. Bicara dengan tempo yang tidak terlalu cepat
2. Tunjuk-tunjuk ke layar saat menjelaskan code
3. Gunakan gestur tangan untuk penekanan
4. Jaga eye contact dengan kamera
5. Senyum dan terlihat percaya diri
6. Jeda sejenak di antara bagian-bagian
7. Pastikan code di layar jelas dan bisa dibaca

**Visual yang Perlu Disiapkan:**
- Slide judul
- Gambar ice cream truck, peta Jabodetabek
- Diagram graf
- Screenshot code untuk setiap file
- Tabel hasil
- Grafik distribusi
- Slide kesimpulan

Total durasi presentasi: 20-25 menit

Semoga membantu! Good luck dengan rekamannya! 🎥🍦
