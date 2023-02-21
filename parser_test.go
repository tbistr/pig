package pig

import (
	"testing"

	"golang.org/x/exp/slices"
)

const A = `<div>
    <p>最初のp</p>
    <p>next p</p>
</div>`

func TestNode_Find(t *testing.T) {
	case1, err := ParseS(A)
	if err != nil {
		t.Errorf("failed to parse: %v", err)
	}

	tests := []struct {
		n    *Node
		m    Match
		want []string
	}{
		{case1, func(node *Node) bool { return node.Data == "p" }, []string{"最初のp", "next p"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			p := tt.n.Find(tt.m)
			got := []string{p[0].Text(), p[1].Text()}
			if !slices.Equal(got, tt.want) {
				t.Errorf("Node.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
