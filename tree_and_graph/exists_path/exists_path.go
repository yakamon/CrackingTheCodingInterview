package ep

import (
	"container/list"
)

// Graph is data structure of graph.
type Graph struct {
	Nodes []*Node
}

// New returns a new graph.
func New() *Graph {
	var nodes []*Node
	return &Graph{nodes}
}

// AppendNode appends one or more node to graph.
func (g *Graph) AppendNode(nodes ...*Node) {
	g.Nodes = append(g.Nodes, nodes...)
}

// Node is data structure of graph node.
type Node struct {
	// Value    interface{}
	Children []*Node
	Visited  bool
}

// NewNode returns a new graph node.
func NewNode() *Node {
	var children []*Node
	return &Node{children, false}
}

// AppendChildren appends children nodes.
func (n *Node) AppendChildren(children ...*Node) {
	n.Children = append(n.Children, children...)
}

// ExistsPath searches graph and checks if path from start to end is exists.
func (g *Graph) ExistsPath(start, end *Node) bool {
	if start == end {
		return true
	}

	for _, n := range g.Nodes {
		n.Visited = false
	}

	q := list.New()
	q.PushBack(start)
	for q.Len() > 0 {
		parent := q.Remove(q.Front()).(*Node)
		if parent == nil {
			continue
		}

		for _, n := range parent.Children {
			if n.Visited != false {
				continue
			}
			if n == end {
				return true
			}
			q.PushBack(n)
		}
		parent.Visited = true
	}
	return false
}
