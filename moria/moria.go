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

type F func(*[]View)

func (f F) View() interface{} {
	children := []View{}
	f(&children)
	return children
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

// Version returns the version of the underlying Mithril library.
func Version() string {
	return m.Version()
}

// Mount takes a Component, converts (or creates) the appropriate mithril
// representation, and then calls it.
func Mount(root dom.Node, component Component) {
	m.Mount(root, wrapComponent(component))
}

// Route takes a mapping of routes to components, creates the appropriate
// mithril representation, and then calls it.
func Route(root dom.Node, initial string, routes map[string]Component) {
	fauxComponents := js.M{}
	for k, component := range routes {
		fauxComponents[k] = wrapComponent(component)
	}
	m.RouteDefine(root, initial, fauxComponents)
}
