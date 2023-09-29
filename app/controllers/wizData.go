package controllers

import (
	"github.com/revel/revel"
)

type WizData struct {
	*revel.Controller
}

func (c WizData) Index() revel.Result {
	return c.RenderTemplate("wiz/index.html")
}
