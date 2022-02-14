# go-mithril
GopherJS bindings to MithrilJS

### Why do this?
> *Because I can*

Mithril.js is a very small and expressive client-side MVC framework. These Go
bindings are intended for Developers like me who welcome the strong typing and
semantics of Go, and want to use it to easily build great front-end experiences.

Mithril's small API size makes it a great target for defining such a binding.
However, it's expressiveness lead to some very JS-centric idioms that are harder
to employ in Go.

### Who is the target audience?

> *Me, essentially.*

To determine if go-mithril is a good fit for your project, ask if you have:

* Developers with a strong preference for writing Go over Javascript.
* Tolerance for slightly lower performance introduced by Gopherjs and having to
  go through a translation.
* Developers with at least a passing knowledge of front-end technologies, and
  finally,
* Developers with at least a passing knowledge of how mithril itself works.

Perhaps as this project matures, sufficient documentation and abstraction solidity
will mean that knowledge of mithril itself will not be necessary.

### How mature is this project?
> *Not very.*

This is a very new project. A lot of the methods exposed in `mithril.go` have
not yet been tested, and those that have had not been tested extensively. If you
elect to use this project, do so with the understanding that the API is still
unstable, and breaking changes have not yet been ruled out.

### How do I use it?
> *Glad you asked.*

The most basic bindings are available in the package
`github.com/danverbraganza/go-mithril`.

For example:

```
m.Render(
    dom.GetWindow().Document().GetElementByID("container"),
    m.M("a[href='/index.html']", nil, "Home"))
)
```

produces the following output

```
<div id="container">
     <a href="index.html">Home</a>
</div>
```

For a look at a slightly larger example,
[here](https://github.com/danverbraganza/go-mithril/blob/master/examples/linkrotor/linkrotor.go)
is some code that is equivalent to the first example involving rotating links at
[http://mithril.js.org/index.html](http://mithril.js.org/index.html)

This example showcases simple creation of components, views and controllers,
making an asynchronous request, and binding a view to a model.

### Moria

The bindings provided in package `mithril` do not provide for an idiomatic Go
approach to crafting front-end code. For a more pleasant approach, use *moria*,
a set of strongly typed functions, types and interfaces you implement to get the
same behaviour.

At present, Moria is only partially implemented.

### Can I contribute?
> *Yes, but...*

At the moment, I'm really enjoying lone-wolfing on the project in my spare time.
Due to the constraints of professional committments I fear I will not be able to
effectively manage contributions in a timely manner. However, contributions will
be gratefully accepted as I able.

### Changelog

2022-02-14: Migrated to using Go modules for dependency management

2016-04-16: Scrapped RenderWithForce in favour of adding the field to render.
Loosened some types in Render. Started work on moria, and brought it to enough
completeness of features so that the Todo app tutorial could be built.

2016-04-12: The basic bindings to Mithril have been completed. However, with no
tests, it's very hard to determine if it is correct.

2016-04-09: With a lot of creakiness and careful conversion between *js.Object
and interface{}, the first working example of this library has been launched. It
is nowhere near being production ready, however.
