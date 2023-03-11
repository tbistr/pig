package pig

// Map returns a slice of the results of applying the function f to each node in the node's children.
func Map[T any](node Node, f func(n Node) T) []T {
	res := []T{}
	for c := node.FirstChild(); c.Node != nil; c = c.NextSibling() {
		res = append(res, f(c))
	}
	return res
}

// Each calls the function f for each node in the node's children.
func (node Node) Each(f func(int, Node)) {
	for i, c := range node.Children() {
		f(i, c)
	}
}

// EachBreak calls the function f for each node in the node's children.
// If f returns true, the iteration is stopped.
func (node Node) EachBreak(f func(int, Node) bool) {
	for i, c := range node.Children() {
		if f(i, c) {
			break
		}
	}
}
