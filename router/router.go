package router

import (
	"syscall/js"

	"github.com/gismo-stack/gismo/compo"
)

type Router struct {
	Routes map[compo.Component]compo.Component
}

var router Router

func init() {
	router.Routes = make(map[compo.Component]compo.Component)
}

func RegisterRoute(path compo.Component, component compo.Component) {
	router.Routes[path] = component
}

func SetRoot(c compo.IComponent) {
	html := c.Render().JsElement
	js.Global().Get("document").Call("getElementById", "view").Call("appendChild", html)
}
