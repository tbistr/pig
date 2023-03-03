package pig

import (
	"strings"

	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
)

type Match func(Node) bool

func Tag(tag string) Match {
	return func(n Node) bool {
		return n.Type == html.ElementNode && n.Data == tag
	}
}

func Cls(cls string) Match {
	return func(n Node) bool {
		val, ok := attrVal(n, "class")
		return ok && slices.Contains(strings.Split(val, " "), cls)
	}
}

func HasAttr(attr string) Match {
	return func(n Node) bool {
		_, ok := attrVal(n, attr)
		return ok
	}
}

func HasAttrVal(attr, val string) Match {
	return func(n Node) bool {
		mayVal, ok := attrVal(n, attr)
		return ok && mayVal == val
	}
}

func attrVal(n Node, attr string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val, true
		}
	}
	return "", false
}
