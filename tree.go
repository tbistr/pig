package pig

import (
	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
)

// NewRoot returns a new root node.
func NewRoot() Node {
	return Node{&html.Node{
		Type: html.DocumentNode,
	}}
}

// MakeTree makes a new root node and appends given nodes as its children.
//
// Given nodes are cloned and detached from their original parent.
// Original nodes are not modified.
func MakeTree(ns ...Node) Node {
	p := NewRoot()
	for _, n := range ns {
		p.AppendChild(n.CloneDetach().Node)
	}
	return p
}

// Clone returns a copy of the node.
//
// It does not a deep copy about the connected nodes.
// So child, parent, and sibling are same pointers as the original one.
// Attributes are copied.
func (n Node) Clone() Node {
	clone := *n.Node
	clone.Attr = slices.Clone(n.Attr)
	return Node{&clone}
}

// CloneDetach returns a copy of the node.
//
// The node is detached from its parent and siblings.
// It has same children as the original one.
func (n Node) CloneDetach() Node {
	nc := n.Clone()
	nc.Node.Parent = nil
	nc.Node.PrevSibling = nil
	nc.Node.NextSibling = nil
	return nc
}

// CloneDetachSubtree returns a copy of the subtree that the node is the root.
// The subtree is detached from its parent and siblings.
// func (n Node) CloneDetachSubtree() Node {
// 	var inner func(Node) Node
// 	inner = func(n Node) Node {
// 		for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
// 			subtree := c.CloneDetach()
// 			n.AppendChild(inner(subtree).Node)
// 		}
// 		return n
// 	}
// 	return inner(n.CloneDetach())
// }

// Children returns a slice of all children of the node.
func (n Node) Children() []Node {
	found := []Node{}
	for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
		found = append(found, c)
	}

	return found
}

// Len returns the number of children of the node.
func (n Node) Len() int {
	i := 0
	for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
		i++
	}
	return i
}

// GetE returns the child node at the given index.
//
// You can use negative index same as slices.
// Second return value is true if index is in range.
// If the index is out of range, it returns a new empty node and false.
func (n Node) GetE(index int) (Node, bool) {
	if index < 0 {
		for c := n.LastChild(); c.Node != nil; c = c.PrevSibling() {
			if index == -1 {
				return c, true
			}
			index++
		}
	} else {
		for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
			if index == 0 {
				return c, true
			}
			index--
		}
	}
	return NewRoot(), false
}

// Get returns the child node at the given index.
//
// You can use negative index same as slices.
// If the index is out of range, it returns a new empty node and false.
func (n Node) Get(index int) Node {
	v, _ := n.GetE(index)
	return v
}

// First returns the first child node.
func (n Node) First() Node {
	return n.Get(0)
}

// Last returns the last child node.
func (n Node) Last() Node {
	return n.Get(-1)
}
