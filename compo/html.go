package compo

import "syscall/js"

type Attr struct {
	Class string
}

type HtmlElement struct {
	Node      string
	Attr      Attr
	JsElement js.Value
}

func (a Attr) isEmpty() bool {
	return a.Class != ""
}

func setAttribute(html js.Value, attr Attr) {
	if attr.isEmpty(){
		html.Call("setAttribute", "class", attr.Class)
	}
}
