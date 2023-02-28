package pig

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Nodes []*html.Node

func Parse(r io.Reader) (Nodes, error) {
	n, err := html.Parse(r)
	return Nodes{n}, err
}

func ParseS(s string) (Nodes, error) {
	return Parse(strings.NewReader(s))
}

func ParseB(b []byte) (Nodes, error) {
	return Parse(bytes.NewReader(b))
}

func (nodes Nodes) Find(m Match) Nodes {
	var inner func(*html.Node)
	var res Nodes
	inner = func(n *html.Node) {
		if m(n) {
			res = append(res, n)
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			inner(c)
		}
	}

	for _, n := range nodes {
		inner(n)
	}
	return res
}

func (nodes Nodes) Text() string {
	var t string
	var inner func(n *html.Node)
	inner = func(n *html.Node) {
		if n.Type == html.TextNode {
			t += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			inner(c)
		}
	}

	for _, n := range nodes {
		inner(n)
	}
	return t
}

func (nodes Nodes) Texts() []string {
	return Map(nodes, func(n *html.Node) string {
		return Nodes{n}.Text()
	})
}

func Map[T any](nodes Nodes, f func(n *html.Node) T) []T {
	res := []T{}
	for _, n := range nodes {
		res = append(res, f(n))
	}
	return res
}
