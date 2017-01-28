package controllers

import "github.com/revel/revel"

type Hello struct {
	*revel.Controller
}

func (c Hello) Index() revel.Result {
	greeting := "Hello World!"
	message := "これはStep2の実装です。(その１)"
   	return c.Render(greeting, message)
}

func (c Hello) Greet(greeting string) revel.Result {
	message := "これはStep2の実装です。(その２)"
	c.RenderArgs["greeting"] = greeting
	c.RenderArgs["message"] = message
	return c.RenderTemplate("Hello/Index.html")
}
