package graph

// Edge merepresentasikan koneksi ke node tujuan dengan bobot waktu (menit).
type Edge struct {
	To     int
	Weight int
}

// Graph direpresentasikan sebagai adjacency list.
type Graph struct {
	Nodes map[int]string // nodeID -> nama lokasi
	Edges map[int][]Edge // nodeID -> daftar edge
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]string),
		Edges: make(map[int][]Edge),
	}
}

func (g *Graph) AddNode(id int, name string) {
	g.Nodes[id] = name
	if _, ok := g.Edges[id]; !ok {
		g.Edges[id] = []Edge{}
	}
}

// AddEdge menambahkan edge undirected: u->v dan v->u.
func (g *Graph) AddEdge(u, v, w int) {
	g.Edges[u] = append(g.Edges[u], Edge{To: v, Weight: w})
	g.Edges[v] = append(g.Edges[v], Edge{To: u, Weight: w})
}

func (g *Graph) NodeName(id int) string {
	if n, ok := g.Nodes[id]; ok {
		return n
	}
	return "(unknown)"
}
