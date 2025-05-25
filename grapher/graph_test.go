package grapher

import (
	"testing"
)

func TestAddVertex(t *testing.T) {
	g := InitGraph("Test")
	v1 := g.AddVertex("A")
	v2 := g.AddVertex("B")

	if len(*g.vertexs) != 2 {
		t.Errorf("expected 2 vertices, got %d", len(*g.vertexs))
	}
	if v1.label != "A" || v2.label != "B" {
		t.Errorf("vertex labels not set correctly")
	}
}

func TestAddEdge(t *testing.T) {
	g := InitGraph("Test")
	a := g.AddVertex("A")
	b := g.AddVertex("B")

	a.AddEdge(b)
	if len(*a.edges[1]) != 1 || (*a.edges[1])[0] != b {
		t.Errorf("edge not added correctly to a")
	}
	if len(*b.edges[1]) != 1 || (*b.edges[1])[0] != a {
		t.Errorf("edge not added correctly to b (undirected)")
	}
}

func TestAddDirectedEdge(t *testing.T) {
	g := InitGraph("Test")
	a := g.AddVertex("A")
	b := g.AddVertex("B")

	a.AddDirectedEdge(b)
	if len(*a.edges[1]) != 1 || (*a.edges[1])[0] != b {
		t.Errorf("directed edge not added correctly to a")
	}
	if b.edges[1] != nil && len(*b.edges[1]) != 0 {
		t.Errorf("directed edge should not be added to b")
	}
}

func TestRemoveEdge(t *testing.T) {
	g := InitGraph("Test")
	a := g.AddVertex("A")
	b := g.AddVertex("B")

	a.AddEdge(b)
	a.RemoveEdge(b)
	if len(*a.edges[1]) != 0 {
		t.Errorf("edge not removed from a")
	}
	if len(*b.edges[1]) != 0 {
		t.Errorf("edge not removed from b")
	}
}

func TestRemoveDirectedEdge(t *testing.T) {
	g := InitGraph("Test")
	a := g.AddVertex("A")
	b := g.AddVertex("B")

	a.AddDirectedEdge(b)
	a.RemoveDirectedEdge(b)
	if len(*a.edges[1]) != 0 {
		t.Errorf("directed edge not removed from a")
	}
}
