package tests

import (
	"github.com/revel/revel/testing"
	"github.com/sirupsen/logrus"
	"test/app/models"
	"test/app/services"
)

var app models.Application

type ApplicationSvcTest struct {
	testing.TestSuite
}

func (t *ApplicationSvcTest) Before() {

	app = models.Application{
		Name:       "Test Application",
		Slug:       "test",
		Repository: "http://github",
		TeamId:     1,
	}

	app = services.CreateApplication(app).Data[0]
}

func (t *ApplicationSvcTest) TestGetById() {
	a := services.GetApplicationById(app.Id)

	t.AssertEqual(uint8(0), a.Critical)
	t.AssertEqual(uint8(0), a.High)
	t.AssertEqual(app.Name, a.Name)
}

func (t *ApplicationSvcTest) TestAppIncomingVulnerabilitiesNotExempt() {

	payload, err := services.RecordAndCheckStatus(app.Slug, "Vulnerabilities: CRITICAL: 1, HIGH: 3, MEDIUM: 0, LOW: 0, INFORMATIONAL: 0")

	if err != nil {
		logrus.Error(err)
	}

	updatedApp := services.GetApplicationById(app.Id)

	t.AssertEqual(uint8(1), updatedApp.Critical)
	t.AssertEqual(uint8(3), updatedApp.High)

	t.AssertEqual(false, payload["isExempt"])
}

func (t *ApplicationSvcTest) TestAppNoIncomingVulnerabilitiesExempt() {

	payload, err := services.RecordAndCheckStatus(app.Slug, "Vulnerabilities: CRITICAL: 0, HIGH: 0, MEDIUM: 0, LOW: 0, INFORMATIONAL: 0")

	if err != nil {
		logrus.Error(err)
	}

	updatedApp := services.GetApplicationById(app.Id)

	t.AssertEqual(uint8(0), updatedApp.Critical)
	t.AssertEqual(uint8(0), updatedApp.High)

	t.AssertEqual(true, payload["isExempt"])
}

func (t *ApplicationSvcTest) TestAppIncomingVulnerabilitiesNoChangeExempt() {

	payload, err := services.RecordAndCheckStatus(app.Slug, "Vulnerabilities: CRITICAL: 0, HIGH: 0, MEDIUM: 0, LOW: 0, INFORMATIONAL: 0")

	if err != nil {
		logrus.Error(err)
	}

	updatedApp := services.GetApplicationById(app.Id)

	t.AssertEqual(uint8(0), updatedApp.Critical)
	t.AssertEqual(uint8(0), updatedApp.High)

	t.AssertEqual(true, payload["isExempt"])
}

func (t *ApplicationSvcTest) TestAppIncomingVulnerabilitiesBadInputNoPreviousVuls() {

	payload, err := services.RecordAndCheckStatus(app.Slug, "")

	if err != nil {
		logrus.Error(err)
	}

	updatedApp := services.GetApplicationById(app.Id)

	t.AssertEqual(uint8(0), updatedApp.Critical)
	t.AssertEqual(uint8(0), updatedApp.High)

	t.AssertEqual(true, payload["isExempt"])
}

func (t *ApplicationSvcTest) TestAppIncomingVulnerabilitiesBadInputWithPreviousVuls() {

	app.Critical = 1
	app.High = 1
	app = services.UpdateApplication(app.Id, app).Data[0]

	payload, err := services.RecordAndCheckStatus(app.Slug, "")

	if err != nil {
		logrus.Error(err)
	}

	updatedApp := services.GetApplicationById(app.Id)

	t.AssertEqual(uint8(1), updatedApp.Critical)
	t.AssertEqual(uint8(1), updatedApp.High)

	t.AssertEqual(false, payload["isExempt"])
}

func (t *ApplicationSvcTest) After() {
	services.DeleteApplication(app.Id)
}
