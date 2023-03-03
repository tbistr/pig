package pig

func MakeNode(ns ...Node) Node {
	p := Node{}
	for _, n := range ns {
		p.AppendChild(n)
	}
	return p
}

func (n Node) Children() []Node {
	return nil
}
