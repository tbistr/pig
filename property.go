package pig

import (
	"bytes"

	"golang.org/x/net/html"
)

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

func (n Node) Parent() Node {
	return Node{n.Node.Parent}
}

func (n Node) FirstChild() Node {
	return Node{n.Node.FirstChild}
}

func (n Node) LastChild() Node {
	return Node{n.Node.LastChild}
}

func (n Node) PrevSibling() Node {
	return Node{n.Node.PrevSibling}
}

func (n Node) NextSibling() Node {
	return Node{n.Node.NextSibling}
}
