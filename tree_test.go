package pig

import (
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func makeNode(data string) Node {
	return Node{&html.Node{Type: html.ElementNode, Data: data}}
}

func TestMakeTree(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		p := MakeTree()

		if p.FirstChild().Node != nil {
			t.Errorf(
				"MakeTree(nil) must return a root node with no children but got %v",
				*p.FirstChild().Node)
		}
	})
	t.Run("Children must be cloned and detached", func(t *testing.T) {
		someParent := makeNode("original parent")
		someParent.AppendChild(makeNode("child who has parent").Node)
		children := []Node{
			makeNode("c1"),
			makeNode("c2"),
			someParent.FirstChild(),
		}

		p := MakeTree(children...)

		for i, c := 0, p.FirstChild(); c.Node != nil; i, c = i+1, c.NextSibling() {
			if c.Node == children[i].Node {
				t.Errorf("MakeTree must clone children but got %v", *c.Node)
			}
			if c.Parent().Node != p.Node {
				t.Errorf("MakeTree must detach children but (%v)'s parent is %v", c.Data, *c.Parent().Node)
			}
		}
	})
}

func TestNode_Clone(t *testing.T) {
	p := makeNode("parent")
	s := makeNode("sibling")
	n := Node{&html.Node{
		Type: html.ElementNode,
		Data: "some data",
		Attr: []html.Attribute{
			{Key: "a", Val: "b"},
			{Key: "c", Val: "d"},
		},
		Parent:      p.Node,
		NextSibling: s.Node,
	}}
	n.Node.AppendChild(makeNode("child").Node)

	clone := n.Clone()

	if clone.Node == n.Node {
		t.Error("Node.Clone must return a new node but got similar pointer")
	}
	if clone.Parent().Node != n.Parent().Node ||
		clone.NextSibling().Node != n.NextSibling().Node ||
		clone.FirstChild().Node != n.FirstChild().Node {
		t.Error("cloned copy must have same parents, siblings, and children but got different ones")
	}
	if &clone.Attr == &n.Attr {
		t.Error("cloned copy must have different attributes but got similar pointer")
	}
	// Each attribute is a struct value, so no need to check pointer.
	if !reflect.DeepEqual(clone.Attr, n.Attr) {
		t.Error("cloned copy must have same attributes but got different ones")
	}
}
