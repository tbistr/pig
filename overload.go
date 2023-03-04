package pig

func (n Node) appendChild(c Node) {
	nc := EmpNode()
	nc.Node.FirstChild = c.Node.FirstChild
	nc.Node.LastChild = c.Node.LastChild

	nc.Node.Type = c.Type
	nc.Node.DataAtom = c.DataAtom
	nc.Node.Data = c.Data
	nc.Node.Namespace = c.Namespace
	copy(nc.Node.Attr, c.Attr)

	n.Node.AppendChild(nc.Node)
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
