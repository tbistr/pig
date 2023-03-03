package pig

import (
	"strings"

	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
)

type Match func(Node) bool

func And(ms ...Match) Match {
	return func(n Node) bool {
		for _, m := range ms {
			if !m(n) {
				return false
			}
		}
		return true
	}
}

func Or(ms ...Match) Match {
	return func(n Node) bool {
		for _, m := range ms {
			if m(n) {
				return true
			}
		}
		return false
	}
}

func Tag(tag string) Match {
	return func(n Node) bool {
		return n.Type == html.ElementNode && n.Data == tag
	}
}

func Cls(cls string) Match {
	return func(n Node) bool {
		val, ok := n.AttrVal("class")
		return ok && slices.Contains(strings.Split(val, " "), cls)
	}
}

func ID(id string) Match {
	return func(n Node) bool {
		val, ok := n.AttrVal("id")
		return ok && slices.Contains(strings.Split(val, " "), id)
	}
}

func HasAttr(attr string) Match {
	return func(n Node) bool {
		_, ok := n.AttrVal(attr)
		return ok
	}
}

func HasAttrVal(attr, val string) Match {
	return func(n Node) bool {
		mayVal, ok := n.AttrVal(attr)
		return ok && mayVal == val
	}
}
