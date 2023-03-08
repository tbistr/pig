package pig

import "golang.org/x/net/html"

func NewRoot() Node {
	return Node{&html.Node{
		Type: html.DocumentNode,
	}}
}

func MakeTree(ns ...Node) Node {
	p := NewRoot()
	for _, n := range ns {
		p.appendChild(n)
	}
	return p
}

func (n Node) appendChild(c Node) {
	nc := NewRoot()
	nc.Node.FirstChild = c.Node.FirstChild
	nc.Node.LastChild = c.Node.LastChild

	nc.Node.Type = c.Type
	nc.Node.DataAtom = c.DataAtom
	nc.Node.Data = c.Data
	nc.Node.Namespace = c.Namespace
	nc.Node.Attr = c.Attr

	n.Node.AppendChild(nc.Node)
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
