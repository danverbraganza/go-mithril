# go-mithril
GopherJS bindings to MithrilJS

This code is super pre-alpha. Things will break, and the interface will be
changed with no notice.

### YOU HAVE BEEN WARNED

## Why do this?

Mithril.js is a very small and expressive client-side MVC framework. These
bindings are intended for Developers like me who welcome the strong typing and
semantics of Go, and want to use it to easily build great front-end experiences.

Mithril's small API size makes it a great target for defining such a binding.
However, it's expressiveness lead to some very JS-centric idioms that are harder
to employ in Go.

My goal is to come up with a translation layer to facilitate app creation in
idiomatic Go, which then calls out to Mithril for all the underlying work.

## Examples

Honestly, watch this space for better examples.

But
[here](https://github.com/danverbraganza/go-mithril/blob/master/examples/linkrotor/linkrotor.go)
is some code that is equivalent to the first example on http://mithril.js.org/index.html

## Contributing

Sorry, at the moment I'm enjoying lone-wolfing this project. Also, it's in a
very inchoate state, and so there's not a substantive way someone could
contribute right now. Finally, I'm enjoying the ability to clobber the git
history with force pushes to master while I'm the only one here :P

Of course, as this project matures beyond simple hackery, I'll add more
structure to make it possible for others to contribute. Until then, feel free to
play around with your own forks, and/or send me an email and I can let you know
when the project will be ready.

## Changelog
2016-04-09: With a lot of creakiness and careful conversion between *js.Object
and interface{}, the first working example of this library has been launched. It
is nowhere near being production ready, however.
