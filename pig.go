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

func (n Node) Find(m Match) Node {
	var inner func(Node)
	found := NewRoot()
	inner = func(in Node) {
		if m(in) {
			found.appendChild(in)
			return
		}
		for c := in.FirstChild(); c.Node != nil; c = c.NextSibling() {
			inner(c)
		}
	}

	inner(n)
	return found
}

func (n Node) FindDescendant(ms ...Match) Node {
	found := n
	for _, m := range ms {
		found = found.Find(m)
	}
	return found
}

func (n Node) FindChild(ms ...Match) Node {
	found := n.Children()
	inner := func(ins []Node, m Match) []Node {
		iFound := []Node{}
		for _, in := range ins {
			for c := in.FirstChild(); c.Node != nil; c = c.NextSibling() {
				if m(c) {
					iFound = append(iFound, c)
				}
			}
		}
		return iFound
	}

	for _, m := range ms {
		found = inner(found, m)
	}
	return MakeTree(found...)
}
