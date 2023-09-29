package controllers

import (
	"net/http"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Health() revel.Result {
	c.Response.Status = http.StatusOK
	c.Response.ContentType = "application/json"

	return c.RenderJSON("UP!")
}
