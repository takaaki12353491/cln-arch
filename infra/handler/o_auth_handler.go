package handler

import (
	"cln-arch/interface/controller"
	inputdata "cln-arch/usecase/input/data"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type OAuthHandler struct {
	controller controller.OAuthController
}

func NewOAuthHandler() *OAuthHandler {
	return &OAuthHandler{}
}

// Auth ...
// @summary
// @description
// @tags OAuth
// @accept json
// @produce json
// @success 307 {object} outputdata.Auth ""
// @failure 400 {string} string ""
// @router /auth/github/auth [get]
func (h *OAuthHandler) Auth(c *Context) error {
	oAuth, err := h.controller.Auth()
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
		return err
	}
	c.Response().Header().Set("Location", oAuth.URL)
	return c.JSON(http.StatusTemporaryRedirect, oAuth)
}

// Callback ...
// @summary
// @description
// @tags OAuth
// @accept json
// @produce json
// @success 200 {object} outputdata.Callback ""
// @failure 400 {string} string ""
// @router /auth/github/callback [get]
func (h *OAuthHandler) Callback(c *Context) error {
	request := &inputdata.CallbackRequest{}
	c.Bind(request)
	oCallback, err := h.controller.Callback(request)
	if err != nil {
		c.String(statusCode(err), err.Error())
		return err
	}
	return c.JSON(http.StatusOK, oCallback)
}
