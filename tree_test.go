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
