
## 1. Struktur project dan file

Project Go yang dideskripsikan:

```text
project/
├── main.go        // Program utama
├── graph.go       // Struktur data graph
├── dijkstra.go    // Implementasi algoritma Dijkstra
└── data.go        // Dataset lokasi dan koneksi
```

Fungsi tiap file:

- **main.go**  
  - Menginisialisasi graph dengan memanggil fungsi/variabel dari `data.go`.  
  - Memanggil fungsi Dijkstra (di `dijkstra.go`) dengan source node = 0 (warehouse).  
  - Mencetak jarak shortest path ke setiap node, rute lengkap (path reconstruction), serta metrik performa (waktu eksekusi).  
  - *Output terminal yang diharapkan*:  
    - Daftar “Shortest distance dari warehouse ke tiap lokasi (node 0–11) dalam menit”.  
    - Rincian rute untuk beberapa destinasi (misalnya Cibinong, Senayan, BSD, Tangerang).  
    - Waktu eksekusi total (mis. `< 1 ms`) dan jumlah node/edge yang diproses.

- **graph.go**  
  - Mendefinisikan struktur data untuk merepresentasikan graf.  
  - Menyediakan operasi dasar: menambah node, menambah edge (undirected → menambah 2 directed edge di adjacency list).  
  - Dipakai oleh `data.go` untuk membangun graph dan oleh `dijkstra.go` untuk traversal.

- **dijkstra.go**  
  - Berisi implementasi fungsi `Dijkstra(graph, source)` menggunakan `container/heap` sebagai priority queue.
  - Menghasilkan:  
    - `distance[]`: jarak/waktu tempuh minimum dari source ke tiap node.  
    - `previous[]`: predecessor setiap node, untuk rekonstruksi path.  
  - Biasanya ada fungsi tambahan untuk `reconstructPath(target)` yang membaca `previous[]` dan membalik urutan node.

- **data.go**  
  - Menyimpan *dataset* statis:  
    - Data node (12 lokasi) dengan ID dan nama.  
    - Data edge (20 koneksi) dengan waktu tempuh (15–50 menit).  
    - ID source node = 0 (Warehouse Momoyo Cibubur).
  - Mengisi struct `Graph` di `graph.go` dengan memanggil fungsi `AddEdge(u, v, weight)` untuk setiap baris koneksi.

## 2. Struktur data utama di Go

Dari bagian “Struktur Data” di laporan:

### 2.1 Struct Edge

```go
type Edge struct {
    to     int  // node tujuan
    weight int  // bobot edge (waktu dalam menit)
}
```

- Dipakai sebagai elemen adjacency list: setiap node punya slice `[]Edge`.  
- Data yang dibutuhkan: ID node tujuan (0–11) dan waktu tempuh (menit) sesuai tabel edge.

### 2.2 Struct Graph

```go
type Graph struct {
    nodes map[int]string   // mapping node ID -> nama lokasi
    edges map[int][]Edge   // adjacency list: node ID -> daftar Edge
}
```

- `nodes` diisi dengan 12 entri (0–11) sesuai tabel “Node Data”.
- `edges` diisi dengan 20 *undirected* koneksi → 40 `Edge` di adjacency list.

Biasanya akan ada fungsi:

```go
func NewGraph() *Graph
func (g *Graph) AddNode(id int, name string)
func (g *Graph) AddEdge(u, v, w int)  // menambah u→v dan v→u
```

### 2.3 Priority queue

Laporan menyebut penggunaan `container/heap` untuk priority queue.
Strukturnya biasanya:

```go
type Item struct {
    node int
    dist int
    index int // untuk heap.Interface
}

type PriorityQueue []*Item
```

Dan diimplementasikan sesuai `heap.Interface` (Len, Less, Swap, Push, Pop).
Fungsi: menyimpan kandidat node dengan jarak sementara; selalu mengeluarkan node dengan `dist` terkecil.

## 3. Data yang diperlukan (node & edge)

### 3.1 Node data (12 node)

- ID 0: Warehouse Momoyo (Cibubur) – Jakarta Timur  
- ID 1: Mall Kelapa Gading – Jakarta Utara  
- ID 2: Grand Indonesia – Jakarta Pusat  
- ID 3: Pondok Indah Mall – Jakarta Selatan  
- ID 4: BSD City – Tangerang Selatan  
- ID 5: Bekasi Cyber Park – Bekasi  
- ID 6: Depok Town Square – Depok  
- ID 7: Tangerang City – Tangerang  
- ID 8: Bogor Trade Mall – Bogor  
- ID 9: Cibinong City Mall – Cibinong  
- ID10: Senayan City – Jakarta Selatan  
- ID11: Summarecon Mall Bekasi – Bekasi  

Ini harus muncul di `nodes map[int]string` dan di output terminal saat menampilkan hasil.

### 3.2 Edge data (20 undirected edges)

Contoh (tidak lengkap):  

- 0–8: 25 menit  
- 0–5: 30 menit  
- 0–9: 15 menit  
- 0–6: 35 menit  
- 5–1: 35 menit  
- 6–2: 40 menit  
- 4–10: 20 menit  
- 10–2: 15 menit  
- 2–3: 30 menit  
- 3–7: 40 menit  
- 7–11: 45–50 menit, dst.  

Semua 20 baris harus dimasukkan ke fungsi `AddEdge` di `data.go`.

## 4. Fungsi Dijkstra dan outputnya

### 4.1 Fungsi inti

Pseudocode di laporan diterjemahkan ke Go kurang lebih:

```go
func Dijkstra(g *Graph, source int) (dist map[int]int, prev map[int]int) {
    // inisialisasi dist dan prev
    // inisialisasi priority queue dengan semua node, dist[source] = 0
    // loop: pop node dengan dist terkecil, relaksasi semua tetangga
    // return dist, prev
}
```

**Input fungsi:**

- `g`: pointer ke `Graph` yang sudah berisi nodes dan edges.  
- `source`: 0 (warehouse Cibubur).  

**Output fungsi:**

- `dist[id]`: waktu tempuh minimum dari warehouse ke node `id` (dalam menit).  
- `prev[id]`: node sebelumnya di shortest path menuju `id`.  

### 4.2 Path reconstruction

Biasanya ada helper:

```go
func ReconstructPath(prev map[int]int, target int) []int {
    // mulai dari target, mundur lewat prev[target], prev[prev[target]], ...
    // sampai source, lalu dibalik urutannya
}
```

**Output di terminal:**

Untuk tiap node penting, misalnya Senayan City:

- “Warehouse → Depok Town Square → Grand Indonesia → Senayan City”  
- Waktu = `30 + 40 + 15 = 85 menit`.  

Laporan memberi contoh rute dan breakdown seperti itu; program sebaiknya menghasilkan format yang mirip.

## 5. Perilaku program di terminal (expected)

Mengacu ke bab 3.3–3.5:

Saat `go run .` dijalankan, idealnya terminal menampilkan:

1. **Daftar shortest distance dari warehouse**  

   Contoh:

   ```text
   Shortest distance dari Warehouse (0) ke setiap lokasi:
   0 -> 9 (Cibinong City Mall): 15 menit
   0 -> 5 (Bekasi Cyber Park): 25 menit
   0 -> 6 (Depok Town Square): 30 menit
   ...
   0 -> 4 (BSD City): 135 menit
   ```

2. **Rute lengkap untuk beberapa node**  

   ```text
   Rute ke Senayan City (10): 0 -> 6 -> 2 -> 10 (Warehouse → Depok → GI → Senayan), total 85 menit
   Rute ke BSD City (4): 0 -> 6 -> 2 -> 3 -> 4, total 135 menit
   ```

3. **Ringkasan metrik**  

   ```text
   Waktu eksekusi: 0.8 ms
   Node diproses: 12
   Edge dievaluasi: ~40
   ```

4. (Opsional) daftar top 3 tercepat dan terlama, atau kategori “sangat dekat/dekat/sedang/jauh” seperti di tabel analisis.

***

Jika ingin, langkah berikutnya bisa dibuatkan skeleton `main.go`, `graph.go`, `dijkstra.go`, dan `data.go` yang persis mengikuti desain di laporan sehingga tinggal kamu lengkapi atau langsung run.Berikut ringkasan seluruh struktur kode Go yang dijelaskan di laporan, termasuk file, struct/fungsi utama, data yang dibutuhkan, dan output yang diharapkan di terminal.

## 1. Struktur project dan file

Project Go yang dideskripsikan:

```text
project/
├── main.go        // Program utama
├── graph.go       // Struktur data graph
├── dijkstra.go    // Implementasi algoritma Dijkstra
└── data.go        // Dataset lokasi dan koneksi
```

Fungsi tiap file:

- **main.go**  
  - Menginisialisasi graph dengan memanggil fungsi/variabel dari `data.go`.  
  - Memanggil fungsi Dijkstra (di `dijkstra.go`) dengan source node = 0 (warehouse).  
  - Mencetak jarak shortest path ke setiap node, rute lengkap (path reconstruction), serta metrik performa (waktu eksekusi).  
  - *Output terminal yang diharapkan*:  
    - Daftar “Shortest distance dari warehouse ke tiap lokasi (node 0–11) dalam menit”.  
    - Rincian rute untuk beberapa destinasi (misalnya Cibinong, Senayan, BSD, Tangerang).  
    - Waktu eksekusi total (mis. `< 1 ms`) dan jumlah node/edge yang diproses.

- **graph.go**  
  - Mendefinisikan struktur data untuk merepresentasikan graf.  
  - Menyediakan operasi dasar: menambah node, menambah edge (undirected → menambah 2 directed edge di adjacency list).  
  - Dipakai oleh `data.go` untuk membangun graph dan oleh `dijkstra.go` untuk traversal.

- **dijkstra.go**  
  - Berisi implementasi fungsi `Dijkstra(graph, source)` menggunakan `container/heap` sebagai priority queue.
  - Menghasilkan:  
    - `distance[]`: jarak/waktu tempuh minimum dari source ke tiap node.  
    - `previous[]`: predecessor setiap node, untuk rekonstruksi path.  
  - Biasanya ada fungsi tambahan untuk `reconstructPath(target)` yang membaca `previous[]` dan membalik urutan node.

- **data.go**  
  - Menyimpan *dataset* statis:  
    - Data node (12 lokasi) dengan ID dan nama.  
    - Data edge (20 koneksi) dengan waktu tempuh (15–50 menit).  
    - ID source node = 0 (Warehouse Momoyo Cibubur).
  - Mengisi struct `Graph` di `graph.go` dengan memanggil fungsi `AddEdge(u, v, weight)` untuk setiap baris koneksi.

## 2. Struktur data utama di Go

Dari bagian “Struktur Data” di laporan:

### 2.1 Struct Edge

```go
type Edge struct {
    to     int  // node tujuan
    weight int  // bobot edge (waktu dalam menit)
}
```

- Dipakai sebagai elemen adjacency list: setiap node punya slice `[]Edge`.  
- Data yang dibutuhkan: ID node tujuan (0–11) dan waktu tempuh (menit) sesuai tabel edge.

### 2.2 Struct Graph

```go
type Graph struct {
    nodes map[int]string   // mapping node ID -> nama lokasi
    edges map[int][]Edge   // adjacency list: node ID -> daftar Edge
}
```

- `nodes` diisi dengan 12 entri (0–11) sesuai tabel “Node Data”.
- `edges` diisi dengan 20 *undirected* koneksi → 40 `Edge` di adjacency list.

Biasanya akan ada fungsi:

```go
func NewGraph() *Graph
func (g *Graph) AddNode(id int, name string)
func (g *Graph) AddEdge(u, v, w int)  // menambah u→v dan v→u
```

### 2.3 Priority queue

Laporan menyebut penggunaan `container/heap` untuk priority queue.
Strukturnya biasanya:

```go
type Item struct {
    node int
    dist int
    index int // untuk heap.Interface
}

type PriorityQueue []*Item
```

Dan diimplementasikan sesuai `heap.Interface` (Len, Less, Swap, Push, Pop).
Fungsi: menyimpan kandidat node dengan jarak sementara; selalu mengeluarkan node dengan `dist` terkecil.

## 3. Data yang diperlukan (node & edge)

### 3.1 Node data (12 node)

- ID 0: Warehouse Momoyo (Cibubur) – Jakarta Timur  
- ID 1: Mall Kelapa Gading – Jakarta Utara  
- ID 2: Grand Indonesia – Jakarta Pusat  
- ID 3: Pondok Indah Mall – Jakarta Selatan  
- ID 4: BSD City – Tangerang Selatan  
- ID 5: Bekasi Cyber Park – Bekasi  
- ID 6: Depok Town Square – Depok  
- ID 7: Tangerang City – Tangerang  
- ID 8: Bogor Trade Mall – Bogor  
- ID 9: Cibinong City Mall – Cibinong  
- ID10: Senayan City – Jakarta Selatan  
- ID11: Summarecon Mall Bekasi – Bekasi  

Ini harus muncul di `nodes map[int]string` dan di output terminal saat menampilkan hasil.

### 3.2 Edge data (20 undirected edges)

Contoh (tidak lengkap):  

- 0–8: 25 menit  
- 0–5: 30 menit  
- 0–9: 15 menit  
- 0–6: 35 menit  
- 5–1: 35 menit  
- 6–2: 40 menit  
- 4–10: 20 menit  
- 10–2: 15 menit  
- 2–3: 30 menit  
- 3–7: 40 menit  
- 7–11: 45–50 menit, dst.  

Semua 20 baris harus dimasukkan ke fungsi `AddEdge` di `data.go`.

## 4. Fungsi Dijkstra dan outputnya

### 4.1 Fungsi inti

Pseudocode di laporan diterjemahkan ke Go kurang lebih:

```go
func Dijkstra(g *Graph, source int) (dist map[int]int, prev map[int]int) {
    // inisialisasi dist dan prev
    // inisialisasi priority queue dengan semua node, dist[source] = 0
    // loop: pop node dengan dist terkecil, relaksasi semua tetangga
    // return dist, prev
}
```

**Input fungsi:**

- `g`: pointer ke `Graph` yang sudah berisi nodes dan edges.  
- `source`: 0 (warehouse Cibubur).  

**Output fungsi:**

- `dist[id]`: waktu tempuh minimum dari warehouse ke node `id` (dalam menit).  
- `prev[id]`: node sebelumnya di shortest path menuju `id`.  

### 4.2 Path reconstruction

Biasanya ada helper:

```go
func ReconstructPath(prev map[int]int, target int) []int {
    // mulai dari target, mundur lewat prev[target], prev[prev[target]], ...
    // sampai source, lalu dibalik urutannya
}
```

**Output di terminal:**

Untuk tiap node penting, misalnya Senayan City:

- “Warehouse → Depok Town Square → Grand Indonesia → Senayan City”  
- Waktu = `30 + 40 + 15 = 85 menit`.  

Laporan memberi contoh rute dan breakdown seperti itu; program sebaiknya menghasilkan format yang mirip.

## 5. Perilaku program di terminal (expected)

Mengacu ke bab 3.3–3.5:

Saat `go run .` dijalankan, idealnya terminal menampilkan:

1. **Daftar shortest distance dari warehouse**  

   Contoh:

   ```text
   Shortest distance dari Warehouse (0) ke setiap lokasi:
   0 -> 9 (Cibinong City Mall): 15 menit
   0 -> 5 (Bekasi Cyber Park): 25 menit
   0 -> 6 (Depok Town Square): 30 menit
   ...
   0 -> 4 (BSD City): 135 menit
   ```

2. **Rute lengkap untuk beberapa node**  

   ```text
   Rute ke Senayan City (10): 0 -> 6 -> 2 -> 10 (Warehouse → Depok → GI → Senayan), total 85 menit
   Rute ke BSD City (4): 0 -> 6 -> 2 -> 3 -> 4, total 135 menit
   ```

3. **Ringkasan metrik**  

   ```text
   Waktu eksekusi: 0.8 ms
   Node diproses: 12
   Edge dievaluasi: ~40
   ```

4. (Opsional) daftar top 3 tercepat dan terlama, atau kategori “sangat dekat/dekat/sedang/jauh” seperti di tabel analisis.

***

langkah berikutnya bisa dibuatkan skeleton `main.go`, `graph.go`, `dijkstra.go`, dan `data.go` yang persis mengikuti desain di laporan 