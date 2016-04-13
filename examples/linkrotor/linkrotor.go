package main

import (
	m "github.com/danverbraganza/go-mithril"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func ListPages() *js.Object {
	return m.Request(js.M{
		"method": "GET",
		"url":    "pages.json",
	})
}

func Controller(this *js.Object, args []*js.Object) interface{} {
	var p = ListPages()
	return js.M{
		"pages": p,
		"rotate": func() {
			actual := p.Invoke().Interface().([]interface{})
			rest, head := actual[:1], actual[1:]
			p.Invoke(append(head, rest...))
		},
	}
}

// View is a function that takes a controller as its first argument, and returns
// a view. Unfortunately, due to the need for compatibility with MakeFunc, the
// signature is poor.
func View(this *js.Object, args []*js.Object) interface{} {
	controller := args[0]
	pages := controller.Get("pages")
	children := js.S{}
	p := pages.Invoke()
	for i := 0; i < p.Length(); i++ {
		page := p.Index(i)
		children = append(
			children,
			m.M("a", js.M{
				"href": page.Get("url").String()},
				page.Get("title").String()))
	}

	children = append(children,
		m.M("button", js.M{
			"onclick": controller.Get("rotate")},
			"Rotate links"))

	return m.M("div", js.M{}, children...)
}

func main() {
	example := dom.GetWindow().Document().GetElementByID("example")
	m.Mount(
		example,
		js.M{
			"view":       js.MakeFunc(View),
			"controller": js.MakeFunc(Controller)},
	)
}
