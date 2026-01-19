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

// Dijkstra calculates the shortest path from source to all other nodes.
func Dijkstra(g *graph.Graph, source int, n int) (dist []int, prev []int) {
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

		for _, e := range g.Edges[u.node] {
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
	return dist, prev
}

// ReconstructPath builds the path from source to target.
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

	// Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
