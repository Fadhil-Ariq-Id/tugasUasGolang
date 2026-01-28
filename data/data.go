package data

import (
	"uasgo/graph"
)

// Dataset sesuai tabel "3.2 Data Lokasi dan Koneksi" pada laporan (halaman 11-12).
// Node ID 0..11 (kontigu).

const (
	SumberNode = 0
	JumlahNode = 12
)

func DataSample() *graph.Graph {
	g := graph.NewGraph()

	// Node data (hal. 11)
	g.AddNode(0, "Warehouse Momoyo (Cibubur)")
	g.AddNode(1, "Mall Kelapa Gading")
	g.AddNode(2, "Grand Indonesia (GI)")
	g.AddNode(3, "Pondok Indah Mall (PIM)")
	g.AddNode(4, "BSD City")
	g.AddNode(5, "Bekasi Cyber Park")
	g.AddNode(6, "Depok Town Square")
	g.AddNode(7, "Tangerang City")
	g.AddNode(8, "Bogor Trade Mall")
	g.AddNode(9, "Cibinong City Mall (CCM)")
	g.AddNode(10, "Senayan City")
	g.AddNode(11, "Summarecon Mall Bekasi")

	// Edge data (hal. 11-12) total 20 koneksi undirected

	// Hal. 11
	g.AddEdge(0, 5, 25)
	g.AddEdge(0, 6, 30)
	g.AddEdge(0, 9, 15)
	g.AddEdge(0, 8, 35)

	// Hal. 12
	g.AddEdge(5, 1, 35)
	g.AddEdge(5, 11, 20)
	g.AddEdge(6, 2, 40)
	g.AddEdge(6, 9, 25)
	g.AddEdge(9, 8, 30)
	g.AddEdge(8, 6, 40)
	g.AddEdge(1, 2, 25)
	g.AddEdge(1, 11, 30)
	g.AddEdge(2, 10, 15)
	g.AddEdge(2, 3, 30)
	g.AddEdge(10, 3, 20)
	g.AddEdge(3, 4, 35)
	g.AddEdge(3, 7, 40)
	g.AddEdge(4, 7, 25)
	g.AddEdge(10, 7, 45)
	g.AddEdge(2, 7, 50)

	return g
}
