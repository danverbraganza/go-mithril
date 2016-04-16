// Package moria contains idiomatic Go data types and helper functions for
// defining Mithril components within Go.
package moria

import (
	m "github.com/danverbraganza/go-mithril"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

// Renderable is anything that can be rendered.
type Renderable interface {
	Renderable() interface{}
}

type VirtualElement struct {
	*js.Object
}

// Render renders things
func Render(root dom.Node, cell Renderable, force bool) {
	m.Render(root, cell.Renderable(), force)
}

func M(selector string, attrs js.M, children ...Renderable) VirtualElement {
	renderableChildren := make([]interface{}, len(children))
	for i, child := range children {
		renderableChildren[i] = child.Renderable()
	}
	return VirtualElement{m.M(selector, attrs, renderableChildren)}
}

type S string

func (s S) Renderable() interface{} {
	return s
}

// VirtualElement implements interface Renderable.
// However, there is a bug here.
func (v VirtualElement) Renderable() interface{} {
	return v.Object
}

func Version() string {
	return m.Version()
}
