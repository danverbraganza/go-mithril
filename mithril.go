// Package mithril exports explicit bindings for the Mithril Javascript Library.
// This version targets v 0.2.3 of the Mithril API.
// See http://mithril.js.org/mithril.html
// These bindings assume that the Mithril script has already been loaded into
// the Global namespace before init() is called.
package mithril

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var m *js.Object

func init() {
	m = js.Global.Get("m")
}

func Version() string {
	return m.Call("version").String()
}

func M(selector string, attrs js.M, children ...interface{}) *js.Object {
	if attrs == nil {
		attrs = js.M{}
	}
	x := js.Global.Get("Array").New()
	for _, c := range children {
		x.Call("push", c)
	}
	return m.Invoke(selector, attrs, x)
}

// Render renders a given virtual element cell to a DOM Node. Will not force
// recreation of elements by default.
func Render(root dom.Node, cell *js.Object) {
	m.Call("render", root.Underlying(), js.InternalObject(cell))
}

// RenderWithForce is the same as Render, but accepts an additional parameter,
// forceRecreation.
func RenderWithForce(root dom.Node, cell *js.Object, forceRecreation bool) {
	m.Call("render", root.Underlying(), js.InternalObject(cell), forceRecreation)
}

// Trust annotates a string as trusted, so that HTML entities will not be escaped.
func Trust(trusted string) *js.Object {
	return m.Call("trust", trusted)
}

func Prop(store *js.Object) *js.Object {
	return m.Call("prop", store)
}

func Component(component *js.Object, args ...*js.Object) *js.Object {
	callArgs := []interface{}{component}
	for _, elem := range args {
		callArgs = append(callArgs, elem)
	}
	return m.Call("component", callArgs...)
}

// Mount "hooks up" an element for continuous rendering at a target node.
// This interface is pretty clunky right now until I figure out what is needed.
// Should be an interface, with view and controller.
func Mount(root dom.Node, element interface{}) *js.Object {
	return m.Call("mount", root.Underlying(), element)
}

func Redraw() {
	m.Call("redraw")
}

func WithAttr(prop, withAttrCallback, callbackThis *js.Object) *js.Object {
	return m.Call("withAttr", prop, withAttrCallback, callbackThis)
}

// Route has not been extracted with care for types.
func Route(root, arg1, arg2, vdom *js.Object) *js.Object {
	return m.Call("route", root, arg1, arg2, vdom)
}

// Sync
func Sync(args ...*js.Object) *js.Object {
	callArgs := js.S{}
	for _, elem := range args {
		callArgs = append(callArgs, elem)
	}
	return m.Call("sync", callArgs...)
}

type Options struct {
	Method,
	Url string
}

func Request(options js.M) *js.Object {
	if options == nil {
		options = js.M{}
	}

	return m.Call("request", options)
}

// Controller is an interface.
type Controller interface {
}
