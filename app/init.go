package app

import (
	"fmt"
	"log"
	"test/app/services"
	"test/app/services/auth"

	_ "github.com/revel/modules"
	"github.com/revel/revel"
	"github.com/sirupsen/logrus"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {

	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	//revel.InterceptFunc(doNothing, revel.BEFORE, &controllers.Auth{})
	//revel.InterceptFunc(checkUser, revel.BEFORE, &controllers.AwsAccount{})
	//revel.InterceptFunc(checkUser, revel.BEFORE, &controllers.Team{})
	//revel.InterceptFunc(checkUser, revel.BEFORE, &controllers.User{})
	//revel.InterceptFunc(checkUser, revel.BEFORE, &controllers.Application{})

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	//revel.OnAppStart(initProvider)
	revel.OnAppStart(DataChecks)
	revel.OnAppStart(services.InitDB)
	revel.OnAppStart(initAuth0)
	// revel.OnAppStart(FillCache)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

// simple example or user auth
//func checkUser(c *revel.Controller) revel.Result {
//	un, _ := c.Session.Get("profile")
//
//	if un == nil {
//		c.Flash.Error("Please log in first")
//		return c.Redirect(controllers.Auth.Login)
//	}
//	return nil
//}

func doNothing(c *revel.Controller) revel.Result { return nil }

// DataChecks Outputs what environment we are in
func DataChecks() {
	logrus.Info(fmt.Sprintf("Running in %s mode", revel.RunMode))
}

func initAuth0() {
	_, err := auth.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}
}
