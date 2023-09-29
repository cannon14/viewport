package api

import (
	"github.com/revel/revel"
	"test/app/services"
)

type WizVulnerabilitiesApi struct {
	*revel.Controller
}

func (c WizVulnerabilitiesApi) UpdateVulnerabilities() revel.Result {

	services.UpdateAppVulnerabilityCount()

	return c.RenderText("done")
}
