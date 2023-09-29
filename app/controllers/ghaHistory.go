package controllers

import (
	"github.com/revel/revel"
)

type GhaHistory struct {
	*revel.Controller
}

func (c GhaHistory) Index() revel.Result {
	return c.RenderTemplate("ghaHistory/index.html")
}
