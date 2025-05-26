package grapher

import (
	"fmt"
	"sort"
)

type Graph struct {
	name      string
	vertexs   *[]vertex
	graphType GraphType
}

type GraphType struct {
	GraphType []string
	Actual    bool
}

func InitGraph(name string) *Graph {
	return &Graph{name: name, vertexs: &[]vertex{}}
}

func (g *Graph) GetGraphType() []string {
	if g.graphType.Actual {
		return g.graphType.GraphType
	}

	// TODO: finish calssification algorithms
	return []string{}
}

func (g *Graph) Print() {
	fmt.Println("graph name:", g.name)
	fmt.Println("it's a ", g.graphType)
	if g.vertexs == nil || len(*g.vertexs) == 0 {
		fmt.Println("  (empty graph)")
		return
	}
	for _, v := range *g.vertexs {
		fmt.Printf("  Vertex %d [%s]:", v.id, v.label)
		edgesOut := []string{}
		for weight, neighbors := range v.edges {
			for _, n := range *neighbors {
				edgesOut = append(edgesOut, fmt.Sprintf(" --%d--> %d[%s]", weight, n.id, n.label))
			}
		}
		if len(edgesOut) == 0 {
			fmt.Print(" (no edges)")
		} else {
			for _, e := range edgesOut {
				fmt.Print(e)
			}
		}
		fmt.Println()
	}
}

// AddVertex - add vertex to graph
func (g *Graph) AddVertex(label string) *vertex {
	if g.vertexs == nil {
		g.vertexs = &[]vertex{}
	}
	v := &vertex{
		id:    int64(len(*g.vertexs)),
		label: label,
		edges: make(map[int32]*[]*vertex),
	}
	*g.vertexs = append(*g.vertexs, *v)
	return v
}

// vertex
type vertex struct {
	edges map[int32]*[]*vertex
	id    int64
	label string
}

func (v *vertex) GetId() int64 {
	return v.id
}

// AddEdgeMannual - add manual spcifed edge
func (v *vertex) AddEdgeManual(u *vertex, weight int32, isDirected bool) {
	if v.edges[weight] == nil {
		v.edges[weight] = &[]*vertex{}
	}
	*v.edges[weight] = append(*v.edges[weight], u)
	sort.Slice(*v.edges[weight], func(i, j int) bool {
		return (*v.edges[weight])[i].id < (*v.edges[weight])[j].id
	})

	if !isDirected {
		if u.edges[weight] == nil {
			u.edges[weight] = &[]*vertex{}
		}
		*u.edges[weight] = append(*u.edges[weight], v)
		sort.Slice(*u.edges[weight], func(i, j int) bool {
			return (*u.edges[weight])[i].id < (*u.edges[weight])[j].id
		})
	}
}

// AddEdge - add edge beetwen two vertices
func (v *vertex) AddEdge(u *vertex) {
	v.AddEdgeManual(u, 1, false)
}

// AddDirectedEdge - add edge from receiver to argument (but not from argument to receiver)
func (v *vertex) AddDirectedEdge(u *vertex) {
	v.AddEdgeManual(u, 1, true)
}

// AddEdgeW - add edge with specifed weight beetwen two vertices
func (v *vertex) AddEdgeW(u *vertex, weight int32) {
	v.AddEdgeManual(u, weight, false)
}

// AddDirectedEdgeW - add edge with specifed weight from receiver to vertex in argument
func (v *vertex) AddDirectedEdgeW(u *vertex, weight int32) {
	v.AddEdgeManual(u, weight, true)
}

// RemoveEdgeManual - remove manual spcifed edge
func (v *vertex) RemoveEdgeManual(u *vertex, weight int32, isDirected bool) {
	if neighbors, ok := v.edges[weight]; ok {
		newNeighbors := make([]*vertex, 0, len(*neighbors))
		for _, n := range *neighbors {
			if n.id != u.id {
				newNeighbors = append(newNeighbors, n)
			}
		}
		*neighbors = newNeighbors
	}

	if !isDirected {
		if neighbors, ok := u.edges[weight]; ok {
			newNeighbors := make([]*vertex, 0, len(*neighbors))
			for _, n := range *neighbors {
				if n.id != v.id {
					newNeighbors = append(newNeighbors, n)
				}
			}
			*neighbors = newNeighbors
		}
	}
}

// RemoveEdge - remove edge beetwen two vertices
func (v *vertex) RemoveEdge(u *vertex) {
	v.RemoveEdgeManual(u, 1, false)
}

// RemoveDirectedEdge - remove edge from receiver to argument (but not from argument to receiver)
func (v *vertex) RemoveDirectedEdge(u *vertex) {
	v.RemoveEdgeManual(u, 1, true)
}

// RemoveEdgeW - remove edge with specifed weight beetwen two vertices
func (v *vertex) RemoveEdgeW(u *vertex, weight int32) {
	v.RemoveEdgeManual(u, weight, false)
}

// RemoveDirectedEdgeW - remove edge with specifed weight from receiver to vertex in argument
func (v *vertex) RemoveDirectedEdgeW(u *vertex, weight int32) {
	v.RemoveEdgeManual(u, weight, true)
}
