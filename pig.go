package pig

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Node struct{ *html.Node }

func Parse(r io.Reader) (Node, error) {
	n, err := html.Parse(r)
	return Node{n}, err
}

func ParseS(s string) (Node, error) {
	return Parse(strings.NewReader(s))
}

func ParseB(b []byte) (Node, error) {
	return Parse(bytes.NewReader(b))
}
