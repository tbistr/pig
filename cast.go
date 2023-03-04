package pig

func MakeNode(ns ...Node) Node {
	p := Node{}
	for _, n := range ns {
		p.appendChild(n)
	}
	return p
}

func (n Node) Children() []Node {
	found := []Node{}
	for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
		found = append(found, c)
	}

	return found
}
