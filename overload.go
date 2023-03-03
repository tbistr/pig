package pig

func (n Node) AppendChild(c Node) {
	n.Node.AppendChild(c.Node)
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
