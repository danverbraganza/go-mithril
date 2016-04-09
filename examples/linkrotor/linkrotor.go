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
			p.Invoke().Call("push", p.Invoke().Call("shift"))
		},
	}
}

func View(this *js.Object, controller []*js.Object) interface{} {
	pages := controller[0].Get("pages")
	children := js.S{}
	p := pages.Invoke()
	for i := 0; i < p.Length(); i++ {
		page := p.Index(i)
		children = append(
			children,
			m.M(
				"a",
				js.M{
					"href": page.Get("url").String()},
				page.Get("title").String()))
	}

	children = append(children,
		m.M(
			"button",
			js.M{"onclick": controller[0].Get("rotate")},
			"Rotate links"))

	return m.M("div", js.M{}, children...)
}

func main() {
	x := js.Global.Get("Object").New()
	x.Set("view", js.MakeFunc(View))
	x.Set("controller", js.MakeFunc(Controller))
	m.Mount(
		dom.GetWindow().Document().GetElementByID("example"),
		x,
	)
}
