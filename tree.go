package pig

import (
	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
)

func NewRoot() Node {
	return Node{&html.Node{
		Type: html.DocumentNode,
	}}
}

func MakeTree(ns ...Node) Node {
	p := NewRoot()
	for _, n := range ns {
		n.CloneDetach()
		p.AppendChild(n.Node)
	}
	return p
}

func (n Node) Clone() Node {
	clone := *n.Node
	clone.Attr = slices.Clone(n.Attr)
	return Node{&clone}
}

func (n Node) CloneDetach() Node {
	nc := n.Clone()
	nc.Node.Parent = nil
	nc.Node.PrevSibling = nil
	nc.Node.NextSibling = nil
	return nc
}

func (n Node) CloneTree() Node {
	var inner func(Node) Node
	inner = func(n Node) Node {
		for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
			subtree := c.CloneDetach()
			n.AppendChild(inner(subtree).Node)
		}
		return n
	}
	return inner(n.Clone())
}

func (n Node) Children() []Node {
	found := []Node{}
	for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
		found = append(found, c)
	}

	return found
}

func (n Node) Len() int {
	i := 0
	for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
		i++
	}
	return i
}

func (n Node) GetE(index int) (Node, bool) {
	if index < 0 {
		for c := n.LastChild(); c.Node != nil; c = c.PrevSibling() {
			if index == -1 {
				return c, true
			}
			index++
		}
	} else {
		for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
			if index == 0 {
				return c, true
			}
			index--
		}
	}
	return NewRoot(), false
}

func (n Node) Get(index int) Node {
	v, _ := n.GetE(index)
	return v
}

func (n Node) First() Node {
	return n.Get(0)
}

func (n Node) Last() Node {
	return n.Get(-1)
}
