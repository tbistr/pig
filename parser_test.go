package pig

import (
	"testing"

	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
)

const A = `
<div>
    <p>最初のp</p>
    <p>next p</p>
</div>`

func TestNode_Find(t *testing.T) {
	case1, err := ParseS(A)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	for _, tt := range []struct {
		name string
		n    Nodes
		m    Match
		want []string
	}{
		{"normal", case1,
			func(node *html.Node) bool { return node.Data == "p" },
			[]string{"最初のp", "next p"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.n.Find(tt.m).Texts()
			if !slices.Equal(got, tt.want) {
				t.Errorf("Node.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
