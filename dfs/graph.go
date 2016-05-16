package dfs

type Digraph struct {
	V   int
	E   int
	adj [][]int
}

func (dg *Digraph) Init(v int) {
	dg.V = v
	dg.E = 0
	dg.adj = make([][]int, v)
}

func (dg *Digraph) GetV(v int) []int {
	return dg.adj[v]
}

func (dg *Digraph) GetAdj() [][]int {
	return dg.adj
}

func (dg *Digraph) AddEdge(v int, w int) {
	for _, data := range dg.adj[v] {
		if data == w {
			return
		}
	}
	dg.adj[v] = append(dg.adj[v], w)
}

func (dg *Digraph) Reverse() Digraph {
	newdg := Digraph{}
	newdg.Init(dg.V)
	for i := range dg.adj {
		for _, data := range dg.adj[i] {
			newdg.AddEdge(data, i)
		}
	}
	return newdg
}
