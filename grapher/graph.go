package grapher

import "fmt"

// Paratygmats:

type Graph struct {
	name  string
	nodes *[]node
}

func InitGraph(name string) *Graph {
	return &Graph{name: name, nodes: &[]node{}}
}

func (g *Graph) Print() {
	fmt.Println("graph name:", g.name)
	fmt.Println(g.nodes)
}

type node struct {
	id    int64
	edges *[]node
}

func (g *Graph) AddNode(n node) *node {
	return &n
}

// func (n *node) AddEdge(e edge) *node {}
