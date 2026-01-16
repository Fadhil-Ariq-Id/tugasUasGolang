package main

// momoyo daatset graph  for doc
func BangunGraph() *Graph {
	g := NewGraph()

	// 12 data lokasi
	g.AddNode(0, "Warehouse Momoyo (Cibubur)")
	g.AddNode(1, "Mall kelapa Gading (Jakarta Pusat)")
	g.AddNode(2, "Grand Indonesia (Jakarta Pusat)")
	g.AddNode(3, "Pondok Indah Mall (Jakarta Pusat)")
	g.AddNode(4, "BSD City (Tangerang)")
	g.AddNode(5, "Bekasi Cyber park (Bekasi Pusat)")
	g.AddNode(6, "Depok Town Square (Depok Pusat)")
	g.AddNode(7, "Tangerang City Square (Tangerang Pusat)")
	g.AddNode(8, "Bogor Trade mall (Bogor Pusat)")
	g.AddNode(9, "Cibinong City Mall (Bogor Pusat)")
	g.AddNode(10, "Senayan City Mall (Jakarta Pusat)")
	g.AddNode(11, "Sumarecon Mall Bekasi (Bekasi Pusat)")

	// hubungkan node dengan edge tiap lokasi berdasarkan jarak yang tak berarha atau undericted graph

	g.AddUndirectedEdge(0, 8, 25)
	g.AddUndirectedEdge(0, 1, 10)
	g.AddUndirectedEdge(0, 2, 15)
	g.AddUndirectedEdge(0, 3, 20)
	g.AddUndirectedEdge(0, 4, 25)
	g.AddUndirectedEdge(0, 5, 30)
	g.AddUndirectedEdge(0, 6, 35)
	g.AddUndirectedEdge(0, 7, 40)
	g.AddUndirectedEdge(0, 9, 45)
	g.AddUndirectedEdge(0, 10, 50)
	g.AddUndirectedEdge(0, 11, 55)
	g.AddUndirectedEdge(1, 2, 10)
	g.AddUndirectedEdge(1, 3, 15)
	g.AddUndirectedEdge(1, 4, 20)
	g.AddUndirectedEdge(1, 5, 25)
	g.AddUndirectedEdge(1, 6, 30)
	g.AddUndirectedEdge(1, 7, 35)
	g.AddUndirectedEdge(1, 9, 45)
	g.AddUndirectedEdge(1, 10, 50)
	g.AddUndirectedEdge(1, 11, 55)
	return g

}
