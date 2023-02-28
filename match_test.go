package pig

import (
	"testing"
)

const HTML1 = `
<div>
    <a atr1="hogeA">a1</a>
    <a atr2="fugaA">a2</a>
    <b atr2="fugaB">b1</b>
    <b atr2="fugaB">b2</b>
</div>`

func TestTag(t *testing.T) {
	ns, _ := ParseS(HTML1)

	for _, tt := range []struct {
		caseName, tag, want string
	}{
		{"case1", "a", "a1a2"},
		{"case2", "c", ""},
	} {
		t.Run(tt.caseName, func(t *testing.T) {
			if got := ns.Find(Tag(tt.tag)).Text(); got != tt.want {
				t.Errorf("Find(Tag(%s)) = %v, want %v", tt.tag, got, tt.want)
			}
		})
	}
}
