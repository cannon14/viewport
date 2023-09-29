package api

import (
	"github.com/revel/revel"
	"github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strconv"
	"test/app/models"
	"test/app/services"
	"time"
)

type ApplicationsApi struct {
	*revel.Controller
}

func (c ApplicationsApi) GetById(id uint16) revel.Result {
	application := services.GetApplicationById(id)
	return c.RenderJSON(application)
}

func (c ApplicationsApi) GetAll() revel.Result {

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

func (c ApplicationsApi) Create() revel.Result {
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	var form models.Form
	err := c.Params.BindJSON(&form)
	if err != nil {
		logrus.Error(err)
	}

	data := form.Data["0"].(map[string]interface{})
	application := models.Application{
		TeamId:     uint16(data["teamId"].(float64)),
		Name:       data["name"].(string),
		Slug:       data["slug"].(string),
		Repository: data["repository"].(string),
	}

	payload := services.CreateApplication(application)
	payload.Draw = draw

	c.Response.Status = http.StatusCreated

	return c.RenderJSON(payload)
}

func (c ApplicationsApi) Update(id uint16) revel.Result {

	idString := strconv.FormatUint(uint64(id), 10)
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	var form models.Form
	err := c.Params.BindJSON(&form)
	if err != nil {
		logrus.Error(err)
	}

	data := form.Data[idString].(map[string]interface{})

	teamId := uint(data["teamId"].(float64))

	application := models.Application{
		TeamId:       uint16(teamId),
		Name:         data["name"].(string),
		Slug:         data["slug"].(string),
		Repository:   data["repository"].(string),
		ExemptReason: data["exemptReason"].(string),
	}

	exemptTo := data["exemptTo"].(string)

	if len(exemptTo) > 0 {
		//I have to remove T00:00:00Z if it comes in because the front end
		//can send yyyy-MM-dd or yyyy-MM-ddT00:00:00Z
		m1 := regexp.MustCompile("T(.+)Z")
		exemptTo = m1.ReplaceAllString(exemptTo, "")
		val, err := time.Parse("2006-01-02", exemptTo)
		if err != nil {
			logrus.Error(err)
		}

		application.ExemptTo = &val
	}

	requestedBy := data["requestedBy"].(float64)
	if requestedBy > 0 {

		val := uint16(requestedBy)
		application.RequestedBy = &val
	}

	approvedBy := data["approvedBy"].(float64)
	if approvedBy > 0 {

		val := uint16(approvedBy)
		application.ApprovedBy = &val
	}

	payload := services.UpdateApplication(id, application)
	payload.Draw = draw

	c.Response.Status = http.StatusOK

	return c.RenderJSON(payload)

}

func (c ApplicationsApi) Delete(id uint16) revel.Result {
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	services.DeleteApplication(id)

	c.Response.Status = http.StatusNoContent

	payload := models.Payload[models.Application]{
		Draw: draw,
	}
	return c.RenderJSON(payload)
}

func (c ApplicationsApi) Status() revel.Result {

	slug := c.Params.Query.Get("slug")
	payload, err := services.ApplicationStatus(slug)

	if err != nil {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON("Not Found")
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON(payload)
}

func (c ApplicationsApi) GetVulnerabilityStatus() revel.Result {

	slug := c.Params.Query.Get("slug")
	vuls := c.Params.Query.Get("vulnerability_counts")

	if len(slug) == 0 || len(vuls) == 0 {
		c.Response.Status = http.StatusUnprocessableEntity
		return c.RenderJSON("Bad Request")
	}

	payload, err := services.RecordAndCheckStatus(slug, vuls)

	if err != nil {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON("Not Found")
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON(payload)
}
