package dijkstra

import (
	"container/heap"
	"math"
	"uasgo/graph"
)

type Item struct {
	node  int
	dist  int
	index int
}

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

// Dijkstra:
// - dist[id]: waktu minimum (menit)
// - prev[id]: predecessor untuk rekonstruksi rute (prev[source] = -1)
// - processedNodes: jumlah node yang benar-benar dipop dari PQ
// - evaluatedEdges: jumlah directed edge yang dievaluasi saat relaksasi
func Dijkstra(g *graph.Graph, source int, n int) (dist []int, prev []int, processedNodes int, evaluatedEdges int) {
	dist = make([]int, n)
	prev = make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
		prev[i] = -1
	}
	dist[source] = 0

	pq := make(PriorityQueue, 0, n)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: source, dist: 0})

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Item)
		if visited[u.node] {
			continue
		}
		visited[u.node] = true
		processedNodes++

		for _, e := range g.Edges[u.node] {
			evaluatedEdges++
			if visited[e.To] {
				continue
			}
			alt := dist[u.node] + e.Weight
			if alt < dist[e.To] {
				dist[e.To] = alt
				prev[e.To] = u.node
				heap.Push(&pq, &Item{node: e.To, dist: alt})
			}
		}
	}
	return dist, prev, processedNodes, evaluatedEdges
}

// ReconstructPath mengembalikan urutan node dari source ke target (termasuk keduanya).
// Jika target tidak terjangkau, mengembalikan nil.
func ReconstructPath(prev []int, source, target int) []int {
	if source == target {
		return []int{source}
	}
	if target < 0 || target >= len(prev) {
		return nil
	}

	path := []int{}
	cur := target
	for cur != -1 {
		path = append(path, cur)
		if cur == source {
			break
		}
		cur = prev[cur]
	}
	if len(path) == 0 || path[len(path)-1] != source {
		return nil
	}

	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

// PathBreakdown: bikin breakdown bobot per edge (buat format "a + b + c").
func PathBreakdown(g *graph.Graph, path []int) (sum int, steps []int) {
	if len(path) < 2 {
		return 0, nil
	}
	for i := 0; i < len(path)-1; i++ {
		u, v := path[i], path[i+1]
		w := edgeWeight(g, u, v)
		steps = append(steps, w)
		sum += w
	}
	return sum, steps
}

func edgeWeight(g *graph.Graph, u, v int) int {
	for _, e := range g.Edges[u] {
		if e.To == v {
			return e.Weight
		}
	}
	return math.MaxInt
}
