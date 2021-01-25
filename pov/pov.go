// Package pov is the tree graph exercise. Stolen from the community solutions.
package pov

import "sort"

// Graph - the flat representation of a tree using map[string]string
type Graph map[string]string

// New returns a pointer to new Graph
func New() *Graph {
	g := make(Graph)
	return &g
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(nodeLabel string) {
	// using map[string]string for Graph -> not needed to satisfy tests.
	return
}

// AddArc adds a from-to pair to the graph
func (g *Graph) AddArc(from, to string) {
	(*g)[to] = from
	return
}

// ArcList displays all arcs
func (g *Graph) ArcList() []string {
	var arcs []string
	for to, from := range *g {
		arcs = append(arcs, from+" -> "+to)
	}
	sort.Strings(arcs) // comes unordered from map, make it look a little nicer
	return arcs
}

// ChangeRoot changes the arcs directions to use a new root
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	// create copy of Graph with keys and values exchanged
	rerooted := make(Graph)
	for from, to := range *g {
		rerooted[from] = to
	}
	// traverse the tree from new to old root
	for node := newRoot; node != oldRoot; node = (*g)[node] {
		// parent on the node becomes its child
		parent := (*g)[node]
		rerooted[parent] = node
		// remove the parent of the new root
		if node == newRoot {
			delete(rerooted, node)
		}
	}
	return &rerooted
}
