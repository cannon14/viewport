package api

import (
	"github.com/revel/revel"
	"net/http"
	"strconv"
	"test/app/services"
)

type DashboardApi struct {
	*revel.Controller
}

func (c DashboardApi) GhaStats() revel.Result {

	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	application := c.Params.Query.Get("application")
	date := c.Params.Query.Get("date")
	payload := services.GetGhaStats(application, date)
	payload.Draw = draw

	c.Response.Status = http.StatusOK
	c.Response.ContentType = "application/json"

	if len(payload.Error) > 0 {
		c.Response.Status = http.StatusInternalServerError
	}

	return c.RenderJSON(payload)
}

func (c DashboardApi) GetAll() revel.Result {

	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	startString := c.Params.Query.Get("start")
	start := 0
	if len(startString) > 0 {
		result, err := strconv.Atoi(startString)
		if err == nil {
			start = result
		}
	}

	lengthString := c.Params.Query.Get("length")
	length := 10
	if len(lengthString) > 0 {
		result, err := strconv.Atoi(lengthString)
		if err == nil {
			length = result
		}
	}
	payload := services.GetAllApplicationsAsPayload(start, length)
	payload.Draw = draw
	return c.RenderJSON(payload)
}
