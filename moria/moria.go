// Package moria contains idiomatic Go data types and helper functions for
// defining Mithril components within Go.
package moria

import (
	m "github.com/danverbraganza/go-mithril"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

// A View is anything that can be rendered.
type View interface {
	View() interface{}
}

type VirtualElement struct {
	*js.Object
}

// Render renders things
func Render(root dom.Node, cell View, force bool) {
	m.Render(root, cell.View(), force)
}

func M(selector string, attrs js.M, children ...View) VirtualElement {
	renderableChildren := make([]interface{}, len(children))
	for i, child := range children {
		renderableChildren[i] = child.View()
	}
	return VirtualElement{m.M(selector, attrs, renderableChildren)}
}

type S string

func (s S) View() interface{} {
	return s
}

func Trust(trusted string) VirtualElement {
	return VirtualElement{m.Trust(trusted)}
}

// VirtualElement implements interface View.
// However, there is a bug here.
func (v VirtualElement) View() interface{} {
	return v.Object
}

func Version() string {
	return m.Version()
}

// Mount takes a Component, converts (or creates) the appropriate mithril
// representation, and then calls it.
func Mount(root dom.Node, component Component) {
	fauxComponent := js.M{
		"init":       component.Init,
		"view":       component.View,
		"controller": component.Controller,
	}

	m.Mount(root, fauxComponent)
}
