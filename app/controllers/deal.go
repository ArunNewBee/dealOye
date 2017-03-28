package controllers

import (
	"dealOye/app/routes"

	"github.com/revel/revel"
)

type Deal struct {
	Application
}

func (c Deal) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.Application.Login())
	}
	return nil
}

func (c Deal) GetDeal() revel.Result {

	return nil
}
