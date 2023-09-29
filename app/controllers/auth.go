package controllers

import (
	"github.com/revel/revel"
	"github.com/sirupsen/logrus"
	"net/url"
	auth2 "test/app/services/auth"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Index() revel.Result {
	return c.RenderTemplate("login/index.html")
}

func (c Auth) Login() revel.Result {

	state, _ := auth2.GenerateRandomState()
	//if err != nil {
	//	return c.RenderError(err)
	//}

	// Save the state inside the session.
	_ = c.Session.Set("state", state)
	//if err != nil {
	//	return c.RenderError(err)
	//}
	auth := auth2.Auth0

	//if err != nil {
	//	return c.RenderError(err)
	//}

	auth0CodeUrl := auth.AuthCodeURL(state)
	return c.Redirect(auth0CodeUrl)
}

func (c Auth) Logout() revel.Result {
	logoutUrl, err := url.Parse("https://" + revel.Config.StringDefault("auth0.domain", "") + "/v2/logout")
	if err != nil {
		c.Flash.Error(err.Error())
		logrus.Error(err.Error())
		return c.Redirect(Dashboard.Index)
	}

	returnTo, err := url.Parse("http://" + c.Request.Host)
	if err != nil {
		c.Flash.Error(err.Error())
		logrus.Error(err.Error())
		return c.Redirect(Dashboard.Index)
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", revel.Config.StringDefault("auth0.client.id", ""))
	logoutUrl.RawQuery = parameters.Encode()

	c.Session.Del("access_token")
	c.Session.Del("profile")
	c.Session.Del("email")

	return c.Redirect(logoutUrl.String())
}

func (c Auth) Callback() revel.Result {

	auth := auth2.Auth0

	if c.Params.Query.Get("state") != c.Session["state"] {
		c.Flash.Error("Invalid state parameter.")
		logrus.Error("Invalid state parameter.")
		return c.Redirect(Dashboard.Index)
	}

	// Exchange an authorization code for a token.
	token, err := auth.Exchange(c.Request.Context(), c.Params.Query.Get("code"))
	if err != nil {
		c.Flash.Error("Failed to exchange an authorization code for a token.")
		logrus.Error("Failed to exchange an authorization code for a token.")
		return c.Redirect(Dashboard.Index)
	}

	idToken, err := auth.VerifyIDToken(c.Request.Context(), token)
	if err != nil {
		c.Flash.Error("Failed to verify ID Token.")
		logrus.Error("Failed to verify ID Token.")
		return c.Redirect(Dashboard.Index)
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		c.Flash.Error(err.Error())
		logrus.Error(err.Error())
		return c.Redirect(Dashboard.Index)
	}

	c.Session.Set("access_token", token.AccessToken)
	c.Session.Set("profile", profile)
	c.Session.Set("email", profile["email"])

	logrus.Info(c.Session.Get("profile"))

	return c.Redirect(Dashboard.Index)
}
