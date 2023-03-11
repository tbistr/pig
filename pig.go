package pig

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Node is a wrapper of html.Node.
type Node struct{ *html.Node }

// Parse parses HTML from io.Reader.
func Parse(r io.Reader) (Node, error) {
	n, err := html.Parse(r)
	return Node{n}, err
}

// ParseS parses HTML from string.
func ParseS(s string) (Node, error) {
	return Parse(strings.NewReader(s))
}

// ParseB parses HTML from []byte.
func ParseB(b []byte) (Node, error) {
	return Parse(bytes.NewReader(b))
}

// FindWithDepth finds nodes that match the given condition.
//
// If depth is more than 0, it searches descendants up to the depth.
// If depth is -1, it searches all descendants.
//
// It returns first node that matches the condition on each branch.
// If no node is found, it returns an empty node.
// Found nodes are bundled in a new root node.
func (n Node) FindWithDepth(m Match, depth int) Node {
	var find func(Node, int)
	found := NewRoot()
	find = func(subject Node, depth int) {
		if m(subject) {
			found.AppendChild(subject.CloneDetach().Node)
			return
		}
		if depth == 0 {
			return
		}
		for c := subject.FirstChild(); c.Node != nil; c = c.NextSibling() {
			find(c, depth-1)
		}
	}

	find(n, depth)
	return found
}

// Find finds nodes that match the given condition.
//
// It returns first node that matches the condition on each branch.
// If no node is found, it returns an empty node.
// Found nodes are bundled in a new root node.
func (n Node) Find(m Match) Node {
	return n.FindWithDepth(m, -1)
}

// FindChild finds a node that matches the given conditions.
//
// It searches for a node that matches each condition in ms, but only among direct children of the previous matching node.
// For example, if len(ms) == 2, it searches a node that matches ms[1] that is a direct child of a node that matches ms[0].
//
// It returns first node that matches the condition on each branch.
// If no node is found, it returns an empty node.
// Found nodes are bundled in a new root node.
func (n Node) FindChild(ms ...Match) Node {
	found := n
	for _, m := range ms {
		found = found.FindWithDepth(m, 1)
	}
	return found
}

// FindDescendant finds a node that matches the given conditions.
//
// It searches for a node that matches each condition in ms, but only among descendants of the previous matching node.
// For example, if len(ms) == 2, it searches a node that matches ms[1] that is a descendant of a node that matches ms[0].
//
// It returns first node that matches the condition on each branch.
// If no node is found, it returns an empty node.
// Found nodes are bundled in a new root node.
func (n Node) FindDescendant(ms ...Match) Node {
	found := n
	for _, m := range ms {
		found = found.FindWithDepth(m, -1)
	}
	return found
}

// FindAllWithDepth finds nodes that match the given condition.
//
// If depth is more than 0, it searches descendants up to the depth.
// If depth is -1, it searches all descendants.
//
// It returns all node that matches the condition.
// If no node is found, it returns an empty node.
// Found nodes are bundled in a new root node.
func (n Node) FindAllWithDepth(m Match, depth int) Node {
	var find func(Node, int)
	found := NewRoot()
	find = func(subject Node, depth int) {
		if m(subject) {
			found.AppendChild(subject.CloneDetach().Node)
		}
		if depth == 0 {
			return
		}
		for c := subject.FirstChild(); c.Node != nil; c = c.NextSibling() {
			find(c, depth-1)
		}
	}

	find(n, depth)
	return found
}

// FindAll finds nodes that match the given condition.
//
// It returns all node that matches the condition.
// If no node is found, it returns an empty node.
// Found nodes are bundled in a new root node.
func (n Node) FindAll(m Match) Node {
	return n.FindAllWithDepth(m, -1)
}

// TODO: Add other FindAll~ methods?
// TODO: Make Find~ methods functional option pattern?
