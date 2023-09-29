package api

import (
	"github.com/revel/revel"
	"net/http"
	"strconv"
	"test/app/models"
	"test/app/services"
)

type GhaHistoryApi struct {
	*revel.Controller
}

func addFilter(c GhaHistoryApi) map[string]string {
	filters := make(map[string]string)
	filters["application_name"] = c.Params.Query.Get("applicationName")
	filters["environment"] = c.Params.Query.Get("environment")
	filters["event"] = c.Params.Query.Get("event")
	filters["reference"] = c.Params.Query.Get("reference")
	filters["status"] = c.Params.Query.Get("status")
	filters["actor"] = c.Params.Query.Get("actor")
	filters["from"] = c.Params.Query.Get("from")
	filters["to"] = c.Params.Query.Get("to")

	for k, v := range filters {
		if len(v) == 0 {
			delete(filters, k)
		}
	}

	return filters
}

func (c GhaHistoryApi) GetAll() revel.Result {

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
	length := 0
	if len(lengthString) > 0 {
		result, err := strconv.Atoi(lengthString)
		if err == nil {
			length = result
		}
	}
	payload := services.GetAllGhaHistoryAsPayload(start, length, addFilter(c))
	payload.Draw = draw
	return c.RenderJSON(payload)
}

func (c GhaHistoryApi) Create() revel.Result {
	var history models.GhaHistory
	err := c.Params.BindJSON(&history)

	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderError(err)

	}

	payload := services.CreateGhaHistory(history)

	c.Response.Status = http.StatusCreated
	c.Response.ContentType = "application/json"

	return c.RenderJSON(payload)
}

func (c GhaHistoryApi) Delete(id int) revel.Result {
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	services.DeleteGhaHistory(id)

	c.Response.Status = http.StatusNoContent
	c.Response.ContentType = "application/json"

	payload := models.Payload[models.AwsAccount]{
		Draw: draw,
	}
	return c.RenderJSON(payload)
}
