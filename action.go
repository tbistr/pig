package pig

import "golang.org/x/net/html"

func (n Node) Find(m Match) Node {
	var inner func(Node)
	var found Node
	inner = func(n Node) {
		if m(n) {
			found.AppendChild(n)
			return
		}
		for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
			inner(c)
		}
	}

	inner(n)
	return found
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

func Map[T any](node Node, f func(n Node) T) []T {
	res := []T{}
	for c := node.FirstChild(); c.Node != nil; c = c.NextSibling() {
		res = append(res, f(c))
	}
	return res
}

func (node Node) Each(f func(Node) Node) []Node {
	return Map(node, f)
}
