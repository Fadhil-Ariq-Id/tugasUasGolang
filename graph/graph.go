
package main

// array menggunakan slice , stuct untuyk edge atau edge list 
type Graph struct {
nodes map[int]string
edges map[int][]int
}

// new graph => data.go => bangun graph 
func NewGraph() *Graph {
return &Graph{
	nodes: make(map[int]string),
	edges: make(map[int][]Edge),

}
}

// add node => data.go => new node for graf 

func (g * Graph) AddNode(id, name string) {
 g.nodes[id] = name
 if _, ok := g.edges[id]; !ok {
	g.edges[id] = []Edge{}

} }

//  menambahakan edge bebas 2 arah atau undreited node )
func (g, *Graph) AddUndirectedEdge(u, v, w int)
{ 
g.edges[u] = append(g.edges[u], Edge{to: v, weight: w})
g.edges[v] = append(g.edges[v], Edge{to: u, weight: w})
}
