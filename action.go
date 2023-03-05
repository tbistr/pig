package pig

func Map[T any](node Node, f func(n Node) T) []T {
	res := []T{}
	for c := node.FirstChild(); c.Node != nil; c = c.NextSibling() {
		res = append(res, f(c))
	}
	return res
}

func (node Node) Each(f func(int, Node)) {
	for i, c := range node.Children() {
		f(i, c)
	}
}

func (node Node) EachBreak(f func(int, Node) bool) {
	for i, c := range node.Children() {
		if f(i, c) {
			break
		}
	}
}
