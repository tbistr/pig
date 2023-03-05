package pig

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Node struct{ *html.Node }

func EmpNode() Node {
	return Node{&html.Node{
		Type: html.DocumentNode,
		Attr: []html.Attribute{},
	}}
}

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
	found := EmpNode()
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
	return MakeNode(found...)
}

func (n Node) Index(i int) (Node, bool) {
	for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
		if i == 0 {
			return c, true
		}
		i--
	}
	return EmpNode(), false
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

func (n Node) Html() (string, error) {
	var b bytes.Buffer
	err := html.Render(&b, n.Node)
	return b.String(), err
}

func (n Node) AttrVal(attr string) (string, bool) {
	for _, a := range n.Node.Attr {
		if a.Key == attr {
			return a.Val, true
		}
	}
	return "", false
}
