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

func TestNode_CloneDetach(t *testing.T) {
	//    p
	//  / | \
	// s1 n  s2
	//    | \
	//    c1 c2
	p := makeNode("parent")
	s1 := makeNode("sibling1")
	s2 := makeNode("sibling2")
	c1 := makeNode("child1")
	c2 := makeNode("child2")
	n := makeNode("some data")
	n.Node.Parent = p.Node
	n.Node.PrevSibling = s1.Node
	n.Node.NextSibling = s2.Node
	n.Node.AppendChild(c1.Node)
	n.Node.AppendChild(c2.Node)

	detached := n.CloneDetach()

	if detached.Node == n.Node {
		t.Error("Node.CloneDetach must return a new node but got similar pointer")
	}
	if detached.Parent().Node != nil ||
		detached.PrevSibling().Node != nil ||
		detached.NextSibling().Node != nil {
		t.Error("detached node must have no parent and siblings but got some")
	}
	if detached.FirstChild().Node != c1.Node ||
		detached.LastChild().Node != c2.Node {
		t.Error("detached node must have same children but got different ones")
	}
}

func TestNode_GetE(t *testing.T) {
	c1 := makeNode("c1")
	c2 := makeNode("c2")
	c3 := makeNode("c3")
	tree := NewRoot()
	tree.AppendChild(c1.Node)
	tree.AppendChild(c2.Node)
	tree.AppendChild(c3.Node)

	for name, tt := range map[string]struct {
		node      Node
		i         int
		want      Node
		wantExist bool
	}{
		"empty":        {NewRoot(), 0, Node{}, false},
		"out of range": {tree, 3, Node{}, false},

		"first":    {tree, 0, c1, true},
		"second":   {tree, 1, c2, true},
		"last":     {tree, -1, c3, true},
		"negative": {tree, -2, c2, true},
	} {
		t.Run(name, func(t *testing.T) {

			got, gotExist := tt.node.GetE(tt.i)

			if gotExist != tt.wantExist {
				t.Errorf("GetE(%v) must returns exist == %t but got %t", tt.i, tt.wantExist, gotExist)
			}
			if tt.wantExist && got.Node != tt.want.Node {
				t.Errorf("GetE(%v) must returns node == %v but got %v", tt.i, tt.want, got)
			}
		})
	}
}
