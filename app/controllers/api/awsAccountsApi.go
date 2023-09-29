package api

import (
	"fmt"
	"github.com/revel/revel"
	"net/http"
	"strconv"
	"test/app/models"
	"test/app/services"
)

type AwsAccountApi struct {
	*revel.Controller
}

func (c AwsAccountApi) GetById(id uint) revel.Result {
	account := services.GetAccountById(id)
	c.ViewArgs["account"] = account
	return c.RenderTemplate("awsAccounts/form.html")
}

func (c AwsAccountApi) GetAll() revel.Result {

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
	payload := services.GetAllAccounts(start, length)
	payload.Draw = draw
	return c.RenderJSON(payload)
}

func (c AwsAccountApi) Create() revel.Result {

	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	account := models.AwsAccount{
		Name:      c.Params.Form.Get("data[0][name]"),
		Alias:     c.Params.Form.Get("data[0][alias]"),
		AccountId: c.Params.Form.Get("data[0][accountId]"),
		Notes:     c.Params.Form.Get("data[0][notes]"),
	}

	payload := services.CreateAccount(account)
	payload.Draw = draw

	c.Response.Status = http.StatusCreated
	c.Response.ContentType = "application/json"

	return c.RenderJSON(payload)
}

func (c AwsAccountApi) Update(id uint) revel.Result {

	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	account := models.AwsAccount{
		Name:      c.Params.Form.Get(fmt.Sprintf("data[%d][name]", id)),
		Alias:     c.Params.Form.Get(fmt.Sprintf("data[%d][alias]", id)),
		AccountId: c.Params.Form.Get(fmt.Sprintf("data[%d][accountId]", id)),
		Notes:     c.Params.Form.Get(fmt.Sprintf("data[%d][notes]", id)),
	}

	payload := services.UpdateAccount(id, account)
	payload.Draw = draw

	c.Response.Status = http.StatusOK
	c.Response.ContentType = "application/json"

	return c.RenderJSON(payload)

}

func (c AwsAccountApi) Delete(id int) revel.Result {
	draw, _ := strconv.Atoi(c.Params.Query.Get("draw"))

	services.DeleteAccount(id)

	c.Response.Status = http.StatusNoContent
	c.Response.ContentType = "application/json"

	payload := models.Payload[models.AwsAccount]{
		Draw: draw,
	}
	return c.RenderJSON(payload)
}
