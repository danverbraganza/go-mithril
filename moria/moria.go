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

// RequestOptions represents the options for an HTTP request made using the MithrilJS API.
type RequestOptions struct {
	Method          string                                                 // HTTP method to use (GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS). Defaults to GET.
	URL             string                                                 // Path name to send the request to, optionally interpolated with values from Params.
	Params          js.M                                                   // Data to be interpolated into the URL and/or serialized into the query string.
	Body            js.M                                                   // Data to be serialized into the body (for other types of requests).
	Async           *bool                                                  // Whether the request should be asynchronous. Defaults to true.
	User            string                                                 // Username for HTTP authorization. Defaults to undefined.
	Password        string                                                 // Password for HTTP authorization. Defaults to undefined.
	WithCredentials bool                                                   // Whether to send cookies to 3rd party domains. Defaults to false.
	Timeout         int                                                    // Amount of milliseconds a request can take before automatically being terminated. Defaults to undefined.
	ResponseType    string                                                 // Expected type of the response. Defaults to "" if Extract is defined, "json" if missing.
	Config          func(xhr *js.Object) *js.Object                        // Exposes the underlying XMLHttpRequest object for low-level configuration and optional replacement (by returning a new XHR).
	Headers         map[string]string                                      // Headers to append to the request before sending it (applied right before Config).
	Type            func(any js.Object) js.Object                          // A constructor to be applied to each object in the response. Defaults to the identity function.
	Serialize       func(any js.Object) string                             // A serialization method to be applied to Body. Defaults to JSON.stringify, or if Body is an instance of FormData or URLSearchParams, defaults to the identity function (i.e. func(value js.Object) js.Object {return value}).
	Deserialize     func(any js.Object) js.Object                          // A deserialization method to be applied to the xhr.Response or normalized xhr.ResponseText. Defaults to the identity function. If Extract is defined, Deserialize will be skipped.
	Extract         func(xhr *js.Object, options RequestOptions) js.Object // A hook to specify how the XMLHttpRequest response should be read. Useful for processing response data, reading headers and cookies. By default this is a function that returns options.Deserialize(parsedResponse), throwing an exception when the server response status code indicates an error or when the response is syntactically invalid. If a custom Extract callback is provided, the xhr parameter is the XMLHttpRequest instance used for the request, and options is the object that was passed to the m.Request call. Additionally, Deserialize will be skipped and the value returned from the Extract callback will be left as-is when the promise resolves.
	Background      bool                                                   // If false, redraws mounted components upon completion of the request. If true, it does not. Defaults to false.
}
type RequestPromise struct {
	*js.Object
}

func Request(ro RequestOptions) RequestPromise {
	options := js.M{}

	if ro.Method != "" {
		options["method"] = ro.Method
	}
	if ro.URL != "" {
		options["url"] = ro.URL
	}
	if ro.Params != nil {
		options["params"] = ro.Params
	}
	if ro.Body != nil {
		options["body"] = ro.Body
	}
	if ro.Async != nil {
		options["async"] = *ro.Async
	}
	if ro.User != "" {
		options["user"] = ro.User
	}
	if ro.Password != "" {
		options["password"] = ro.Password
	}
	options["withCredentials"] = ro.WithCredentials
	if ro.Timeout != 0 {
		options["timeout"] = ro.Timeout
	}
	if ro.ResponseType != "" {
		options["responseType"] = ro.ResponseType
	}
	if ro.Config != nil {
		options["config"] = ro.Config
	}
	if ro.Headers != nil {
		options["headers"] = ro.Headers
	}
	if ro.Type != nil {
		options["type"] = ro.Type
	}
	if ro.Serialize != nil {
		options["serialize"] = ro.Serialize
	}
	if ro.Deserialize != nil {
		options["deserialize"] = ro.Deserialize
	}
	if ro.Extract != nil {
		options["extract"] = ro.Extract
	}
	options["background"] = ro.Background

	return RequestPromise{m.Request(options)}
}

// Then adds a callback function to be executed when the RequestPromise resolves.
func (r RequestPromise) Then(callable func(data *js.Object)) {
	r.Call("then", js.MakeFunc(
		func(this *js.Object, args []*js.Object) interface{} {
			callable(args[0])
			return nil
		}))
}
