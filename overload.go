package pig

import "golang.org/x/net/html"

func (n Node) appendChild(c Node) {
	nc := &html.Node{
		FirstChild: c.Node.FirstChild,
		LastChild:  c.Node.LastChild,

		Type:      c.Type,
		DataAtom:  c.DataAtom,
		Data:      c.Data,
		Namespace: c.Namespace,
		Attr:      c.Attr,
	}
	n.Node.AppendChild(nc)
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
