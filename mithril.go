// Package mithril exports explicit bindings for the Mithril Javascript Library.
// This version targets v 0.2.3 of the Mithril API.
// See http://mithril.js.org/mithril.html
// These bindings assume that the correct Mithril script has already been loaded into
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

// Version returns the version of the underlying Mithril library.
func Version() string {
	return m.Call("version").String()
}

// M composes virtual elements that can be rendered via Render().
// The variable parameter children must be either strings or other virtual
// elements.
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

// Render renders a given virtual element cell to a DOM Node. Iff force is true,
// this will force the recreation of elements.
func Render(root dom.Node, cell *js.Object, force bool) {
	m.Call("render", root, cell, force)
}

// Trust annotates a string as trusted, so that HTML entities will not be escaped.
func Trust(trusted string) *js.Object {
	return m.Call("trust", trusted)
}

// Prop creates a getter/setter function.
// store is the initial value.
func Prop(store *js.Object) *js.Object {
	return m.Call("prop", store)
}

// Component initializes a Component object by parametrizing it with args.
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
	return m.Call("mount", root, element)
}

// Redraw requests an aggressive redraw of the view.
// This redraw happens even if there are pending AJAX requests or other async
// requests, so make sure templates have null checks in place to account for
// this. Redraws will not occur if a redraw is currently in progress.
func Redraw(force bool) {
	m.Call("redraw", force)
}

// RedrawStrategy sets the strategy of redrawing, to either "all", "diff", or
// "none". Call with one or no arguments.
func RedrawStrategy(strategy ...string) string {
	strategyFunc := m.Get("redraw").Get("strategy").Invoke
	if len(strategy) > 0 {
		return strategyFunc(strategy[0]).String()
	}
	return strategyFunc().String()
}

// WithAttr is an event handler factory. It returns a method that can be bound
// to a DOM element's event listener, to implement databinding from the view to
// the model.
func WithAttr(prop, withAttrCallback, callbackThis *js.Object) *js.Object {
	return m.Call("withAttr", prop, withAttrCallback, callbackThis)
}

// RouteDefine allows you to define the routes for a Single-Page Application.
func RouteDefine(rootElement dom.Node, defaultRoute string, routes js.M) *js.Object {
	return m.Call("route", rootElement, defaultRoute, routes)
}

// Property for getting and setting the route mode. Call with 0 or 1 arguments
func RouteMode(mode ...string) string {
	modeFunc := m.Get("route").Get("mode").Invoke
	if len(mode) > 0 {
		return modeFunc(mode[0]).String()
	}
	return modeFunc().String()
}

// RouteRedirect automatically and programmatically allows you to redirect to a
// given route.
func RouteRedirect(path string, params js.M, replaceHistory bool) {
	m.Call("route", path, params, replaceHistory)
}

// Route returns the current route. If you're looking for bindings to the other
// Mithril "route" methods, look at RouteDefine, RouteMode and RouteRedirect.
// RouteAbstraction has not been implemented so far.
func Route() string {
	return m.Call("route").String()
}

// BuildQueryString serializes an object into its URI encoded querystring
// representation.
func BuildQueryString(object *js.M) string {
	return m.Get("route").Call("buildQueryString", object).String()
}

// ParseQueryString deserializes an object from its URI encoded querystring
// representation.
func ParseQueryString(queryString string) *js.Object {
	return m.Get("route").Call("parseQueryString", queryString)
}

// Deferred constructs a modified deferred object.
func Deferred() *js.Object {
	return m.Call("deferred")
}

// Sync composes an array of promises into one promise.
func Sync(args ...*js.Object) *js.Object {
	callArgs := js.S{}
	for _, elem := range args {
		callArgs = append(callArgs, elem)
	}
	return m.Call("sync", callArgs...)
}

// StartComputation is used in conjuction with EndComputation to signal to
// Mithril when asynchronous work has been completed, and a redraw should
// happen.
func StartComputation() {
	m.Call("startComputation")
}

// EndComputation is used in conjuction with StartComputation to signal to
// Mithril when asynchronous work has been completed, and a redraw should
// happen.
func EndComputation() {
	m.Call("endComputation")
}

// Create an asynchronous HTTP request to some url. Returns a deferred prop that
// can be invoked to get the correct value.
func Request(options js.M) *js.Object {
	if options == nil {
		options = js.M{}
	}

	return m.Call("request", options)
}

// Deps is used in testing to replace the default window object on which Mithril
// depends.
func Deps(window *js.Object) *js.Object {
	return m.Call("deps", window)
}
