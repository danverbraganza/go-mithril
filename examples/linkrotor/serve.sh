#!/bin/sh

gopherjs build -w linkrotor.go &
gopherjs serve github.com/danverbraganza/go-mithril/examples/linkrotor.html
