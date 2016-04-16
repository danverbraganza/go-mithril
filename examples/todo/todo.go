// A simple MVC TODO app. Made with moria.
package main

import (
	"github.com/danverbraganza/go-mithril/moria"
	"honnef.co/go/js/dom"
)

func main() {
	example := dom.GetWindow().Document().GetElementByID("example")

	moria.Render(
		example,
		moria.M("div[contenteditable]", nil,
			moria.M("b", nil,
				moria.S("TODO APP GOES HERE"))),
		false)
}
