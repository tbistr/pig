package pig

import (
	"bytes"

	"golang.org/x/net/html"
)

// Text returns the text of the node.
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

// Html returns the HTML representation of the node.
func (n Node) Html() (string, error) {
	var b bytes.Buffer
	err := html.Render(&b, n.Node)
	return b.String(), err
}

// AttrVal returns the value of the attribute.
func (n Node) AttrVal(attr string) (string, bool) {
	for _, a := range n.Node.Attr {
		if a.Key == attr {
			return a.Val, true
		}
	}
	return "", false
}

// Parent returns the parent node.
func (n Node) Parent() Node {
	return Node{n.Node.Parent}
}

// FirstChild returns the first child node.
func (n Node) FirstChild() Node {
	return Node{n.Node.FirstChild}
}

// LastChild returns the last child node.
func (n Node) LastChild() Node {
	return Node{n.Node.LastChild}
}

// PrevSibling returns the previous sibling node.
func (n Node) PrevSibling() Node {
	return Node{n.Node.PrevSibling}
}

// NextSibling returns the next sibling node.
func (n Node) NextSibling() Node {
	return Node{n.Node.NextSibling}
}
