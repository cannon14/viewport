package api

import (
	"github.com/revel/revel"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"test/app/models"
	"test/app/services"
	"test/app/utils"
)

type UsersApi struct {
	*revel.Controller
}

func (c UsersApi) GetById(id uint16) revel.Result {
	user := services.GetUserById(id)
	return c.RenderJSON(user)
}

func (c UsersApi) GetAll() revel.Result {

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
	payload := services.GetAllUsers(start, length)
	payload.Draw = draw
	return c.RenderJSON(payload)
}

func (c UsersApi) Create() revel.Result {

	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	var form models.Form
	err := c.Params.BindJSON(&form)
	if err != nil {
		logrus.Error(err)
	}

	data := form.Data["0"].(map[string]interface{})
	teamData := data["teams"].([]interface{})

	teamIds := utils.ConvertSlice[uint16](teamData)

	hash, err := utils.HashPassword(data["password"].(string))

	if err != nil {
		logrus.Error(err)
		c.Response.Status = http.StatusInternalServerError
		c.Response.ContentType = "application/json"
		return c.RenderJSON("Internal Server Error")
	}

	user := models.User{
		Name:     data["name"].(string),
		Title:    data["title"].(string),
		Email:    data["email"].(string),
		Password: hash,
	}

	payload := services.CreateUser(user, teamIds)
	payload.Draw = draw

	c.Response.Status = http.StatusCreated
	c.Response.ContentType = "application/json"

	return c.RenderJSON(payload)
}

func (c UsersApi) Update(id uint16) revel.Result {

	idString := strconv.FormatUint(uint64(id), 10)
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	var form models.Form
	err := c.Params.BindJSON(&form)

	if err != nil {
		logrus.Error(err)
	}

	data := form.Data[idString].(map[string]interface{})
	teamData := data["teams"].([]interface{})

	teamIds := utils.ConvertSlice[uint16](teamData)

	hash := data["password"].(string)

	if len(hash) < 59 {
		hash, err = utils.HashPassword(data["password"].(string))
		if err != nil {
			logrus.Error(err)
			c.Response.Status = http.StatusInternalServerError
			c.Response.ContentType = "application/json"
			return c.RenderJSON("Internal Server Error")
		}
	}
	user := models.User{
		Name:     data["name"].(string),
		Title:    data["title"].(string),
		Email:    data["email"].(string),
		Password: hash,
	}

	payload, err := services.UpdateUserForPayload(id, user)

	if err != nil {
		services.UpdateUserTeamMap(id, teamIds)
		c.Response.Status = http.StatusNotFound
		c.Response.ContentType = "application/json"
		return c.RenderJSON("Not Found")
	}

	payload.Draw = draw

	c.Response.Status = http.StatusOK
	c.Response.ContentType = "application/json"

	return c.RenderJSON(payload)

}

func (c UsersApi) Delete(id uint16) revel.Result {
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	services.DeleteUser(id)

	c.Response.Status = http.StatusNoContent
	c.Response.ContentType = "application/json"

	payload := models.Payload[models.User]{
		Draw: draw,
	}
	return c.RenderJSON(payload)
}
