package pig

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Node struct{ *html.Node }

func Parse(r io.Reader) (Node, error) {
	n, err := html.Parse(r)
	return Node{n}, err
}

func ParseS(s string) (Node, error) {
	return Parse(strings.NewReader(s))
}

func ParseB(b []byte) (Node, error) {
	return Parse(bytes.NewReader(b))
}

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

func (n Node) Find(m Match) Node {
	return n.FindWithDepth(m, -1)
}

func (n Node) FindChild(ms ...Match) Node {
	found := n
	for _, m := range ms {
		found = found.FindWithDepth(m, 1)
	}
	return found
}

func (n Node) FindDescendant(ms ...Match) Node {
	found := n
	for _, m := range ms {
		found = found.FindWithDepth(m, -1)
	}
	return found
}
