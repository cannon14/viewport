package api

import (
	"fmt"
	"github.com/revel/revel"
	"net/http"
	"strconv"
	"test/app/models"
	"test/app/services"
)

type TeamsApi struct {
	*revel.Controller
}

func (c TeamsApi) GetById(id uint) revel.Result {
	team := services.GetTeamById(id)
	return c.RenderJSON(team)
}

func (c TeamsApi) GetAll() revel.Result {

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
	payload := services.GetAllTeams(start, length)
	payload.Draw = draw
	return c.RenderJSON(payload)
}

func (c TeamsApi) Create() revel.Result {

	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	team := models.Team{
		Name: c.Params.Form.Get("data[0][name]"),
	}

	payload := services.CreateTeam(team)
	payload.Draw = draw

	c.Response.Status = http.StatusCreated
	c.Response.ContentType = "application/json"

	return c.RenderJSON(payload)
}

func (c TeamsApi) Update(id uint) revel.Result {

	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	team := models.Team{
		Name: c.Params.Form.Get(fmt.Sprintf("data[%d][name]", id)),
	}

	payload := services.UpdateTeam(id, team)
	payload.Draw = draw

	c.Response.Status = http.StatusOK
	c.Response.ContentType = "application/json"

	return c.RenderJSON(payload)

}

func (c TeamsApi) Delete(id uint) revel.Result {
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	services.DeleteTeam(id)

	c.Response.Status = http.StatusNoContent
	c.Response.ContentType = "application/json"

	payload := models.Payload[models.Team]{
		Draw: draw,
	}
	return c.RenderJSON(payload)
}
