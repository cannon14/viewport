package controllers

import (
	"github.com/revel/revel"
)

type Team struct {
	*revel.Controller
}

func (c Team) Index() revel.Result {
	return c.RenderTemplate("team/index.html")
}
