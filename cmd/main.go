package main

import (
	"fmt"

	"github.com/4d4mszweda/grapher/grapher"
)

func main() {
	fmt.Println("Hello world")
	graph := grapher.InitGraph("Testowy")

	a := graph.AddVertex("A")
	b := graph.AddVertex("B")
	c := graph.AddVertex("C")
	d := graph.AddVertex("D")
	e := graph.AddVertex("E")
	f := graph.AddVertex("F")

	a.AddEdge(b)
	b.AddEdge(c)
	c.AddEdge(a)
	d.AddEdge(e)
	f.AddEdge(e)

	graph.Print()
}
