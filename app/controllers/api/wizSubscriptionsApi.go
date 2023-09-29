package api

import (
	"github.com/revel/revel"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"test/app/models"
	"test/app/services"
)

type WizSubscriptionsApi struct {
	*revel.Controller
}

func (c WizSubscriptionsApi) GetById(id uint16) revel.Result {
	wizProject := services.GetWizSubscriptionById(id)
	return c.RenderJSON(wizProject)
}

func (c WizSubscriptionsApi) GetAll() revel.Result {

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
	payload := services.GetAllWizSubscriptionAsPayload(start, length)
	payload.Draw = draw
	return c.RenderJSON(payload)
}

func (c WizSubscriptionsApi) Create() revel.Result {

	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	var form models.Form
	err := c.Params.BindJSON(&form)
	if err != nil {
		logrus.Error(err)
	}

	data := form.Data["0"].(map[string]interface{})

	wizSubscription := models.WizSubscription{
		Name:           data["name"].(string),
		SubscriptionId: data["subscriptionId"].(string),
		ProjectId:      uint16(data["projectId"].(float64)),
		Environment:    data["environment"].(string),
		Notes:          data["notes"].(string),
	}

	payload := services.CreateWizSubscription(wizSubscription)
	payload.Draw = draw

	c.Response.Status = http.StatusCreated
	c.Response.ContentType = "application/json"

	return c.RenderJSON(payload)
}

func (c WizSubscriptionsApi) Update(id uint16) revel.Result {
	idString := strconv.FormatUint(uint64(id), 10)
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	var form models.Form
	err := c.Params.BindJSON(&form)
	if err != nil {
		logrus.Error(err)
	}

	data := form.Data[idString].(map[string]interface{})

	wizSubscription := models.WizSubscription{
		Name:           data["name"].(string),
		SubscriptionId: data["subscriptionId"].(string),
		ProjectId:      uint16(data["projectId"].(float64)),
		Environment:    data["environment"].(string),
		Notes:          data["notes"].(string),
	}

	payload := services.UpdateWizSubscription(id, wizSubscription)
	payload.Draw = draw

	c.Response.Status = http.StatusCreated
	c.Response.ContentType = "application/json"

	return c.RenderJSON(payload)
}

func (c WizSubscriptionsApi) Delete(id uint16) revel.Result {
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	services.DeleteWizSubscription(id)

	c.Response.Status = http.StatusNoContent
	c.Response.ContentType = "application/json"

	payload := models.Payload[models.WizSubscription]{
		Draw: draw,
	}
	return c.RenderJSON(payload)
}
