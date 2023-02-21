package pig

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Node html.Node
type Nodes []*Node

type Match func(*Node) bool

func Parse(r io.Reader) (*Node, error) {
	n, err := html.Parse(r)
	nn := Node(*n)
	return &nn, err
}

func ParseS(s string) (*Node, error) {
	return Parse(strings.NewReader(s))
}

func ParseB(b []byte) (*Node, error) {
	return Parse(bytes.NewReader(b))
}

func (node *Node) Find(m Match) Nodes {
	var inner func(*Node)
	var res Nodes
	inner = func(n *Node) {
		if m(n) {
			res = append(res, n)
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			tmp := Node(*c)
			inner(&tmp)
		}
	}

	inner(node)
	return res
}

func (n *Node) Text() string {
	var inner func(n *Node)
	var t string
	inner = func(n *Node) {
		if n.Type == html.TextNode {
			t += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			tmp := Node(*c)
			inner(&tmp)
		}
	}

	inner(n)
	return t
}
