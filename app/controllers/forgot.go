package controllers

import "github.com/revel/revel"

type ForgotPassword struct {
	*revel.Controller
}

func (c ForgotPassword) ForgotPassword() revel.Result {
	return c.Render()
}