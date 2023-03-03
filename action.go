package pig

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
