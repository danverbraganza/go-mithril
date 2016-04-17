package moria

// Component is a reusable Mithril component.
type Component interface {
	Init()
	Controller() Controller
	View(Controller) View
}

// Controller is an empty interface right now.
type Controller interface {
}
