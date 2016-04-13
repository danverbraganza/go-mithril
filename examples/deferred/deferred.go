// This package shows the deferred API in action.
package main

import (
	"fmt"
	"time"

	m "github.com/danverbraganza/go-mithril"
)

func main() {
	d := m.Deferred()
	c := time.After(time.Second)
	go func() {
		<-c
		d.Call("resolve", "Hello World")
	}()

	d.Get("promise").Call("then", fmt.Println)
}
