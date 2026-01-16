package main

// item berisi queue = antrian
type item struct {
	node int // node tujuan
	dist int // jarak dari node sumber ke node tujuan
	idx  int // index untuk min heap
}

// antrian prioritas berdasarkan dist/ jarak min -heap = efeisien program (yang terkecil sellau diatas (binary tree))
// Priority Queue digunakan dalam algoritma Dijkstra untuk selalu mengambil node dengan jarak terkecil

type PriorityQueue []*item

// Len: mengembalikan panjang/ukuran dari priority queue
// interface sort.Interface untuk implementasi heap
func (pq PriorityQueue) Len() int { return len(pq) }

// Less: membandingkan dua item di index i dan j (nilai terkecil diatas )
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

// Swap: menukar posisi dua item di index i dan j
// TUJUAN SWAP:
//  1. Menukar elemen pq[i] dengan pq[j] dalam array
//  2. Memperbarui field idx pada kedua item agar sesuai dengan posisi baru mereka
//     (idx penting untuk operasi heap seperti heap.Fix yang perlu tahu posisi item)
//
// Swap diperlukan oleh heap untuk mempertahankan struktur min-heap
// ketika ada perubahan nilai atau penambahan/penghapusan elemen
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i] // Tukar posisi elemen di array
	pq[i].idx = i               // Update index item yang sekarang di posisi i
	pq[j].idx = j               // Update index item yang sekarang di posisi j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.idx = n
	*pq = append(*pq, item)
}
