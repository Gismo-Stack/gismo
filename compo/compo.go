package compo

import (
	"syscall/js"
)

// TODO: Sort functions by logical application ("containers", "text styles", etc)

type IComponent interface {
	Render(children ...HtmlElement) HtmlElement
}

type Component struct {
	This IComponent
}

func (c Component) Render(children ...HtmlElement) HtmlElement {
	return H2(Attr{}, "Hit default Gismo Render path!")
}

func A(attr Attr, text string, href string) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "a")
	html.Set("innerHTML", text)
	html.Set("href", href)

	setAttribute(html, attr)

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func B(attr Attr, text string, children ...HtmlElement) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "b")
	html.Set("innerHTML", text)

	setAttribute(html, attr)

	for _, child := range children {
		html.Call("appendChild", child.JsElement)
	}

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func I(attr Attr, text string, children ...HtmlElement) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "i")
	html.Set("innerHTML", text)

	setAttribute(html, attr)

	for _, child := range children {
		html.Call("appendChild", child.JsElement)
	}

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func H1(attr Attr, text string) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "h1")
	html.Set("innerHTML", text)

	setAttribute(html, attr)

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func H2(attr Attr, text string) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "h2")
	html.Set("innerHTML", text)

	setAttribute(html, attr)

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func H3(attr Attr, text string) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "h3")
	html.Set("innerHTML", text)

	setAttribute(html, attr)

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func H4(attr Attr, text string) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "h4")
	html.Set("innerHTML", text)

	setAttribute(html, attr)

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func H5(attr Attr, text string) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "h5")
	html.Set("innerHTML", text)

	setAttribute(html, attr)

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func H6(attr Attr, text string) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "h6")
	html.Set("innerHTML", text)

	setAttribute(html, attr)

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func Div(attr Attr, children ...HtmlElement) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "div")

	setAttribute(html, attr)

	for _, child := range children {
		html.Call("appendChild", child.JsElement)
	}

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func P(attr Attr, text string, children ...HtmlElement) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "p")
	html.Set("innerHTML", text)

	setAttribute(html, attr)

	for _, child := range children {
		html.Call("appendChild", child.JsElement)
	}

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func Span(attr Attr, text string, children ...HtmlElement) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "span")
	html.Set("innerHTML", text)

	setAttribute(html, attr)

	for _, child := range children {
		html.Call("appendChild", child.JsElement)
	}

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func Br(attr Attr) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "br")

	setAttribute(html, attr)

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}

func Img(attr Attr, src string) HtmlElement {
	document := js.Global().Get("document")
	html := document.Call("createElement", "img")
	html.Set("src", src)

	setAttribute(html, attr)

	return HtmlElement{
		Node:      html.String(),
		Attr:      Attr{Class: attr.Class},
		JsElement: html,
	}
}
