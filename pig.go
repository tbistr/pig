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
	var found Node
	inner = func(n Node) {
		if m(n) {
			found.appendChild(n)
			return
		}
		for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
			inner(c)
		}
	}

	inner(n)
	return found
}

func (n Node) Index(i int) (Node, bool) {
	for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
		if i == 0 {
			return c, true
		}
		i--
	}
	return Node{}, false
}

func (n Node) Text() string {
	var t string
	var inner func(n Node)
	inner = func(n Node) {
		if n.Type == html.TextNode {
			t += n.Data
		}
		for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
			inner(c)
		}
	}

	inner(n)
	return t
}

func (n Node) AttrVal(attr string) (string, bool) {
	for _, a := range n.Node.Attr {
		if a.Key == attr {
			return a.Val, true
		}
	}
	return "", false
}
