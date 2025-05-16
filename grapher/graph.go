package grapher

import "fmt"

type Graph struct {
	name    string
	vertexs *[]vertex
}

func InitGraph(name string) *Graph {
	return &Graph{name: name, vertexs: &[]vertex{}}
}

func (g *Graph) Print() {
	fmt.Println("graph name:", g.name)
}

func (g *Graph) AddVertex(vertexName string) *vertex {
	return &vertex{}
}

// vertex
type vertex struct {
	edges map[int32]*[]*vertex
	id    int64
	label string
}

// AddEdgeMannual - add manual spcifed edge
func (v *vertex) AddEdgeManual(vertex *vertex, weight int32, isDirected bool) {}

// AddEdge - add edge beetwen two vertices
func (v *vertex) AddEdge(vertex *vertex) {}

// AddDirectedEdge - add edge from receiver to argument (but not from argument to receiver)
func (v *vertex) AddDirectedEdge(vertex *vertex) {}

// AddEdgeW - add edge with specifed weight beetwen two vertices
func (v *vertex) AddEdgeW(vertex *vertex, weight int32) {}

// AddDirectedEdgeW - add edge with specifed weight from receiver to vertex in argument
func (v *vertex) AddDirectedEdgeW(vertex *vertex, weight int32) {}

// RemoveEdgeManual - remove manual spcifed edge
func (v *vertex) RemoveEdgeManual(vertex *vertex, weight int32, isDirected bool) {}

// RemoveEdge - remove edge beetwen two vertices
func (v *vertex) RemoveEdge(vertex *vertex) {}

// RemoveDirectedEdge - remove edge from receiver to argument (but not from argument to receiver)
func (v *vertex) RemoveDirectedEdge(vertex *vertex) {}

// RemoveEdgeW - remove edge with specifed weight beetwen two vertices
func (v *vertex) RemoveEdgeW(vertex *vertex) {}

// AddDirectedEdgeW - add edge with specifed weight from receiver to vertex in argument
func (v *vertex) RemoveDirectedEdgeW(vertex *vertex) {}
