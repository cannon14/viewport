package api

import (
	"github.com/revel/revel"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"test/app/models"
	"test/app/services"
)

type WizApplicationsApi struct {
	*revel.Controller
}

func (c WizApplicationsApi) GetById(id uint16) revel.Result {
	wizApplication := services.GetWizApplicationById(id)
	return c.RenderJSON(wizApplication)
}

func (c WizApplicationsApi) GetAll() revel.Result {

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
	payload := services.GetAllWizApplicationsAsPayload(start, length)
	payload.Draw = draw
	return c.RenderJSON(payload)
}

func (c WizApplicationsApi) Create() revel.Result {

	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	var form models.Form
	err := c.Params.BindJSON(&form)
	if err != nil {
		logrus.Error(err)
	}

	data := form.Data["0"].(map[string]interface{})

	wizApplication := models.WizApplication{
		Name:           data["name"].(string),
		SubscriptionId: uint16(data["subscriptionId"].(float64)),
		ApplicationId:  uint16(data["applicationId"].(float64)),
		Notes:          data["notes"].(string),
	}

	payload := services.CreateWizApplication(wizApplication)
	payload.Draw = draw

	c.Response.Status = http.StatusCreated
	c.Response.ContentType = "application/json"

	return c.RenderJSON(payload)
}

func (c WizApplicationsApi) Update(id uint16) revel.Result {
	idString := strconv.FormatUint(uint64(id), 10)
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	var form models.Form
	err := c.Params.BindJSON(&form)
	if err != nil {
		logrus.Error(err)
	}

	data := form.Data[idString].(map[string]interface{})

	wizApplication := models.WizApplication{
		Name:           data["name"].(string),
		SubscriptionId: uint16(data["subscriptionId"].(float64)),
		ApplicationId:  uint16(data["applicationId"].(float64)),
		Notes:          data["notes"].(string),
	}

	payload := services.UpdateWizApplication(id, wizApplication)
	payload.Draw = draw

	c.Response.Status = http.StatusCreated
	c.Response.ContentType = "application/json"

	return c.RenderJSON(payload)
}

func (c WizApplicationsApi) Delete(id uint16) revel.Result {
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	services.DeleteWizApplication(id)

	c.Response.Status = http.StatusNoContent
	c.Response.ContentType = "application/json"

	payload := models.Payload[models.WizApplication]{
		Draw: draw,
	}
	return c.RenderJSON(payload)
}
