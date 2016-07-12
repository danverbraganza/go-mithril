package moria

import "github.com/gopherjs/gopherjs/js"

// Component is a reusable Mithril component.
type Component interface {
	Controller() Controller
	View(Controller) View
}

// Controller is an empty interface right now.
type Controller interface {
}

// wrapComponent is used internally to wrap Go components to mithril.
// Unfortunately, roundtripping the result of Controller() through Mithril and
// GopherJS replaces the Go pointer with a copy. This is not what we want, so
// this function provides a way to sneak the correct result of Controller across
// to View.
func wrapComponent(c Component) js.M {
	// Keep this to memoize controller.
	var controller Controller

	return js.M{
		"controller": func() Controller {
			controller = c.Controller()
			return controller
		},

		"view": func(Controller) View {
			return c.View(controller)
		},
	}

}
