package pig

func (n Node) Children() []Node {
	found := []Node{}
	for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
		found = append(found, c)
	}

	return found
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
	return EmpNode(), false
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

func (n Node) ChildrenNum() int {
	i := 0
	for c := n.FirstChild(); c.Node != nil; c = c.NextSibling() {
		i++
	}
	return i
}
