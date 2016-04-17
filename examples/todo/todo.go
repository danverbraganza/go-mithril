// A simple MVC TODO app implemented using Moria.
package main

import (
	"fmt"

	m "github.com/danverbraganza/go-mithril"
	"github.com/danverbraganza/go-mithril/moria"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

// Todo is a struct that will hold each item of our Todo app.
type Todo struct {
	Description string
	Done        bool
}

// TodoVM is the ViewModel.
type TodoVM struct {
	// Description is the current description typed into the input box.
	Description string
	// list is a slice of Todo items.
	list []Todo
}

// Add creates a new Todo list from the current description, and adds it to the
// list.
func (vm *TodoVM) Add() {
	if vm.Description != "" {
		vm.list = append(vm.list, Todo{Description: vm.Description})
		vm.Description = ""
	}
}

// TodoComponent is the component that holds the state of the object.
// TodoComponent implements moria.Controller.
type TodoComponent struct {
	ViewModel TodoVM
}

// Init is a Noop.
func (TodoComponent) Init() {}

// For this component, the controller is the same.
// This is probably bad.
func (t *TodoComponent) Controller() moria.Controller {
	return t
}

// View returns a renderable view for this Component.
func (t *TodoComponent) View(c moria.Controller) moria.View {

	// The head of the table.
	tableContents := []moria.View{
		moria.M("tr", nil,
			moria.M("th", nil,
				moria.S("Done")),
			moria.M("th", nil, moria.S("Task Description")),
		),
	}

	// Create each individual todo element.
	for i, task := range t.ViewModel.list {
		textDecor := "none"
		if task.Done {
			textDecor = "line-through"
		}
		tableContents = append(
			tableContents,
			moria.M("tr", nil,
				moria.M("td", nil,
					moria.M("input[type=checkbox]", js.M{
						"checked": (t.ViewModel.list[i]).Done,
						"onchange": m.WithAttr(
							"checked",
							moria.BoolProp(
								&(t.ViewModel.list[i].Done))),
					})),
				moria.M("td", js.M{"style": js.M{
					"textDecoration": textDecor,
				}},
					moria.S(task.Description)),
			),
		)
	}

	// Create the main view, and include the individual elements.
	return moria.M("html", nil,
		moria.M("body", nil,
			moria.M("input", js.M{
				"value": &t.ViewModel.Description,
				"onchange": m.WithAttr(
					"value",
					moria.StringProp(&t.ViewModel.Description)),
			}),
			moria.M("button", js.M{
				"onclick": (&t.ViewModel).Add}, moria.S("Add")),
			moria.M("table", nil, tableContents...,
			),
		),
	)
}

func main() {
	// Here's an example of the struct in action.
	myTask := Todo{Description: "Write Code"}

	fmt.Println(myTask.Description) // Prints "Write Code"
	fmt.Println(myTask.Done)        // Prints false

	// Here's an example of how the list works.
	var list = []Todo{}

	fmt.Println(len(list)) // Prints 0.

	// Initialize a view model to play around with.
	vm := TodoVM{}
	vm.Description = ""
	// Try adding a Todo
	vm.Add()
	fmt.Println(len(vm.list)) // Prints 0, because the task is rejected.

	// Example of string properties.
	stringy := "Fail"
	// Create a property by using pointers!
	stringyProp := moria.StringProp(&stringy)
	// Mithril can call the property how it expects.
	stringyProp.Invoke("Success")
	fmt.Println(stringy) // Prints "Success"

	// Now let's put the entire application together.
	myComponent := &TodoComponent{}
	// And mount it.
	moria.Mount(dom.GetWindow().Document(), myComponent)
}
