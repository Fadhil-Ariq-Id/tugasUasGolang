package dijkstra

import (
	"container/heap"
	"math"
	"uasgo/graph"
)

// Item merepresentasikan satu node dalam priority queue.
type Item struct {
	node  int // ID dari node
	dist  int // Jarak sementara dari source
	index int // Keperluan internal package container/heap
}

// PriorityQueue mengimplementasikan heap.Interface untuk menyimpan Item.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

// Dijkstra mengeksekusi algoritma jalur terpendek dari satu titik awal (source).
// Mengembalikan slice jarak (dist) dan slice predecessor (prev) untuk rekonstruksi jalur.
func Dijkstra(g *graph.Graph, source int, n int) (dist []int, prev []int) {
	// 1. Inisialisasi: Semua jarak dianggap tak hingga, predecessor tidak terdefinisi (-1).
	dist = make([]int, n)
	prev = make([]int, n)
	visited := make([]bool, n) // Untuk optimasi agar tidak memproses node yang sudah stabil.

	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
		prev[i] = -1
	}
	dist[source] = 0

	// Inisialisasi Priority Queue.
	pq := make(PriorityQueue, 0, n)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: source, dist: 0})

	// 2. Loop Utama: Proses node dengan jarak terkecil yang tersedia.
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Item)

		// Jika node sudah pernah diproses, lewati.
		if visited[u.node] {
			continue
		}
		visited[u.node] = true

		// 3. Relaksasi: Periksa semua tetangga dari node u.
		for _, e := range g.Edges[u.node] {
			if visited[e.To] {
				continue
			}

			// Hitung alternatif jarak baru.
			alt := dist[u.node] + e.Weight
			if alt < dist[e.To] {
				dist[e.To] = alt
				prev[e.To] = u.node
				// Masukkan ke priority queue untuk diproses di iterasi berikutnya.
				heap.Push(&pq, &Item{node: e.To, dist: alt})
			}
		}
	}
	return dist, prev
}

// ReconstructPath menyusun urutan node dari source ke target berdasarkan data predecessor (prev).
func ReconstructPath(prev []int, source, target int) []int {
	if source == target {
		return []int{source}
	}
	if target < 0 || target >= len(prev) {
		return nil
	}

	path := []int{}
	cur := target
	// Telusuri mundur dari target ke source.
	for cur != -1 {
		path = append(path, cur)
		if cur == source {
			break
		}
		cur = prev[cur]
	}

	// Jika target tidak terhubung ke source.
	if len(path) == 0 || path[len(path)-1] != source {
		return nil
	}

	// Balikkan path agar urutannya menjadi source -> ... -> target.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
