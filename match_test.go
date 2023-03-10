package pig

import (
	"testing"

	"golang.org/x/net/html"
)

var (
	div_2attrs = Node{&html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{
			{Key: "k1", Val: "foo bar   baz		hoge\n\n fuga"},
			{Key: "k2", Val: "bar"},
		},
	}}

	p_cls = Node{&html.Node{
		Type: html.ElementNode,
		Data: "p",
		Attr: []html.Attribute{
			{Key: "class", Val: "cls1 cls2 cls3"},
		},
	}}

	div_id = Node{&html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{
			{Key: "id", Val: "id1 id2"},
		},
	}}

	T = func(n Node) bool { return true }
	F = func(n Node) bool { return false }
)

func TestMatch_And(t *testing.T) {
	for name, tt := range map[string]struct {
		m    Match
		ms   []Match
		want bool
	}{
		"T":           {T, []Match{}, true},
		"F":           {F, []Match{}, false},
		"T && T":      {T, []Match{T}, true},
		"T && F":      {T, []Match{F}, false},
		"F && T":      {F, []Match{T}, false},
		"F && F":      {F, []Match{F}, false},
		"T && T && T": {T, []Match{T, T}, true},
	} {
		t.Run(name, func(t *testing.T) {
			if got := tt.m.And(tt.ms...)(div_2attrs); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatch_Or(t *testing.T) {
	for name, tt := range map[string]struct {
		m    Match
		ms   []Match
		want bool
	}{
		"T":           {T, []Match{}, true},
		"F":           {F, []Match{}, false},
		"T || T":      {T, []Match{T}, true},
		"T || F":      {T, []Match{F}, true},
		"F || T":      {F, []Match{T}, true},
		"F || F":      {F, []Match{F}, false},
		"T || T || T": {T, []Match{T, T}, true},
	} {
		t.Run(name, func(t *testing.T) {
			if got := tt.m.Or(tt.ms...)(div_2attrs); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatch_Not(t *testing.T) {
	for name, tt := range map[string]struct {
		m    Match
		want bool
	}{
		"T": {T, false},
		"F": {F, true},
	} {
		t.Run(name, func(t *testing.T) {
			if got := tt.m.Not()(div_2attrs); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	for name, tt := range map[string]struct {
		ms   []Match
		want bool
	}{
		"no args":     {[]Match{}, true},
		"T":           {[]Match{T}, true},
		"F":           {[]Match{F}, false},
		"T && T":      {[]Match{T, T}, true},
		"T && F":      {[]Match{T, F}, false},
		"F && T":      {[]Match{F, T}, false},
		"F && F":      {[]Match{F, F}, false},
		"T && T && T": {[]Match{T, T, T}, true},
	} {
		t.Run(name, func(t *testing.T) {
			if got := And(tt.ms...)(div_2attrs); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOr(t *testing.T) {
	for name, tt := range map[string]struct {
		ms   []Match
		want bool
	}{
		"no args":     {[]Match{}, false},
		"T":           {[]Match{T}, true},
		"F":           {[]Match{F}, false},
		"T || T":      {[]Match{T, T}, true},
		"T || F":      {[]Match{T, F}, true},
		"F || T":      {[]Match{F, T}, true},
		"F || F":      {[]Match{F, F}, false},
		"T || T || T": {[]Match{T, T, T}, true},
	} {
		t.Run(name, func(t *testing.T) {
			if got := Or(tt.ms...)(div_2attrs); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTag(t *testing.T) {
	for name, tt := range map[string]struct {
		node Node
		tag  string
		want bool
	}{
		"match1":     {div_id, "div", true},
		"match2":     {p_cls, "p", true},
		"not match1": {div_2attrs, "p", false},
		"not match2": {p_cls, "div", false},
		"no arg":     {div_2attrs, "", false},
	} {
		t.Run(name, func(t *testing.T) {
			if got := Tag(tt.tag)(tt.node); got != tt.want {
				html, _ := tt.node.Html()
				t.Errorf("\n"+
					"node = %s\n"+
					"Tag(%s)(node) = %v, want %v",
					html, tt.tag, got, tt.want)
			}
		})
	}
}

func TestCls(t *testing.T) {
	for name, tt := range map[string]struct {
		node Node
		cls  string
		want bool
	}{
		"match1":          {p_cls, "cls1", true},
		"match2":          {p_cls, "cls2", true},
		"not match1":      {p_cls, "cls4", false},
		"not match2":      {div_2attrs, "cls1", false},
		"cant match 2cls": {p_cls, "cls1 cls2", false},
		"no arg":          {div_2attrs, "", false},
	} {
		t.Run(name, func(t *testing.T) {
			if got := Cls(tt.cls)(tt.node); got != tt.want {
				html, _ := tt.node.Html()
				t.Errorf("\n"+
					"node = %s\n"+
					"Cls(%s)(node) = %v, want %v",
					html, tt.cls, got, tt.want)
			}
		})
	}
}

func TestID(t *testing.T) {
	for name, tt := range map[string]struct {
		node Node
		id   string
		want bool
	}{
		"match1":         {div_id, "id1", true},
		"match2":         {div_id, "id2", true},
		"not match1":     {div_id, "id3", false},
		"not match2":     {p_cls, "id1", false},
		"cant match ids": {div_id, "id1 id2", false},
		"no arg":         {div_2attrs, "", false},
	} {
		t.Run(name, func(t *testing.T) {
			if got := ID(tt.id)(tt.node); got != tt.want {
				html, _ := tt.node.Html()
				t.Errorf("\n"+
					"node = %s\n"+
					"ID(%s)(node) = %v, want %v",
					html, tt.id, got, tt.want)
			}
		})
	}
}

func TestHasAttr(t *testing.T) {
	for name, tt := range map[string]struct {
		node Node
		attr string
		want bool
	}{
		"match1":    {div_2attrs, "k1", true},
		"match2":    {div_2attrs, "k2", true},
		"not match": {div_2attrs, "k1 k2", false},
		"no arg":    {div_2attrs, "", false},
	} {
		t.Run(name, func(t *testing.T) {
			if got := HasAttr(tt.attr)(tt.node); got != tt.want {
				html, _ := tt.node.Html()
				t.Errorf("\n"+
					"node = %s\n"+
					"HasAttr(%s)(node) = %v, want %v",
					html, tt.attr, got, tt.want)
			}
		})
	}
}

func TestHasAttrVal(t *testing.T) {
	for name, tt := range map[string]struct {
		node      Node
		attr, val string
		want      bool
	}{
		"match1":     {div_2attrs, "k1", "foo bar   baz		hoge\n\n fuga", true},
		"match2":     {div_2attrs, "k2", "bar", true},
		"not match1": {div_2attrs, "k1", "foo bar", false},
		"no arg":     {div_2attrs, "", "", false},
	} {
		t.Run(name, func(t *testing.T) {
			if got := HasAttrVal(tt.attr, tt.val)(tt.node); got != tt.want {
				html, _ := tt.node.Html()
				t.Errorf("\n"+
					"node = %s\n"+
					"HasAttrVal(%s, %s)(node) = %v, want %v",
					html, tt.attr, tt.val, got, tt.want)
			}
		})
	}
}
