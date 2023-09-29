package controllers

import (
	"github.com/revel/revel"
)

type AwsAccount struct {
	*revel.Controller
}

func (c AwsAccount) Index() revel.Result {
	return c.RenderTemplate("awsAccounts/index.html")
}
