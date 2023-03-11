package pig

import (
	"strings"

	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Match func(Node) bool

// And returns a Match that returns AND of the given Matches.
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

// Or returns a Match that returns OR of the given Matches.
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

// And returns a Match that returns AND of the given Matches.
func (m Match) And(ms ...Match) Match {
	return func(n Node) bool {
		if !m(n) {
			return false
		}
		return And(ms...)(n)
	}
}

// Or returns a Match that returns OR of the given Matches.
func (m Match) Or(ms ...Match) Match {
	return func(n Node) bool {
		if m(n) {
			return true
		}
		return Or(ms...)(n)
	}
}

// Not returns a Match that returns negation of the given Match.
func (m Match) Not() Match {
	return func(n Node) bool {
		return !m(n)
	}
}

// Tag returns a Match that returns true if the given node's tag is the given tag.
func Tag(tag string) Match {
	return func(n Node) bool {
		return n.Type == html.ElementNode && n.Data == tag
	}
}

// Cls returns a Match that returns true if the given node has the given class.
func Cls(cls string) Match {
	return func(n Node) bool {
		val, ok := n.AttrVal("class")
		return ok && slices.Contains(strings.Split(val, " "), cls)
	}
}

// ID returns a Match that returns true if the given node has the given id.
func ID(id string) Match {
	return func(n Node) bool {
		val, ok := n.AttrVal("id")
		return ok && slices.Contains(strings.Split(val, " "), id)
	}
}

// HasAttr returns a Match that returns true if the given node has the given attribute.
func HasAttr(attr string) Match {
	return func(n Node) bool {
		_, ok := n.AttrVal(attr)
		return ok
	}
}

// HasAttrVal returns a Match that returns true if the given node has the given attribute and the attribute's value is the given value.
func HasAttrVal(attr, val string) Match {
	return func(n Node) bool {
		mayVal, ok := n.AttrVal(attr)
		return ok && mayVal == val
	}
}

// Atom returns a Match that returns true if the given node's atom is the given atom.
func Atom(a atom.Atom) Match {
	return func(n Node) bool {
		return n.DataAtom == a
	}
}
