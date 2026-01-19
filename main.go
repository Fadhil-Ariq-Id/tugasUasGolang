package main

import (
	"fmt"
	"strings"
	"uasgo/data"
	"uasgo/dijkstra"
	"uasgo/graph"
)

func main() {
	// 1. Build Graph
	g := data.BuildGraphFromDoc()

	// 2. Run Dijkstra
	fmt.Printf("Calculating shortest paths from: %s\n\n", g.NodeName(data.SourceNode))
	dist, prev := dijkstra.Dijkstra(g, data.SourceNode, data.NodeCount)

	// 3. Print Results
	fmt.Printf("%-25s %-15s %s\n", "Tujuan", "Waktu Tempuh", "Rute")
	fmt.Println(strings.Repeat("-", 60))

	for id := 0; id < data.NodeCount; id++ {
		if id == data.SourceNode {
			continue
		}

		path := dijkstra.ReconstructPath(prev, data.SourceNode, id)
		pathNames := formatPath(g, path)

		fmt.Printf("%-25s %-15s %s\n", g.NodeName(id), fmt.Sprintf("%d menit", dist[id]), pathNames)
	}
}

func formatPath(g *graph.Graph, path []int) string {
	if len(path) == 0 {
		return "Tidak Terjangkau"
	}
	names := make([]string, 0, len(path))
	for _, id := range path {
		names = append(names, g.NodeName(id))
	}
	return strings.Join(names, " -> ")
}
