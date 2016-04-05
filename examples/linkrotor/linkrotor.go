package main

import (
	m "github.com/danverbraganza/go-mithril"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type Page struct{}

func ListPages() *js.Object {
	return m.Request(js.M{
		"method": "GET",
		"url":    "pages.json",
	})
}

type Demo struct {
	p *js.Object
}

func (d *Demo) Controller() map[string]interface{} {
	d.p = ListPages()
	return js.M{
		"pages": d.p,
		"rotate": func() {
			d.p.Invoke().Call("push", d.p.Invoke().Call("shift"))
		},
	}
}

func (d Demo) View() m.VirtualElement {
	var log = js.Global.Get("console").Get("log")

	children := js.S{}
	log.Invoke(d.p.Invoke())
	for _, rawPage := range d.p.Invoke().Interface().(js.S) {
		page := rawPage.(js.Object)
		children = append(
			children,
			m.M("a", js.M{
				"href": page.Get("url")}, page.Get("title")))
	}

	return m.M("div", nil, children...)
}

func main() {
	demo := Demo{}
	d.Controller()
	js.Global.Get("console").Get("dir").Invoke(demo)

	m.Render(
		dom.GetWindow().Document().GetElementByID("example"),
		m.M("input", nil))
}
