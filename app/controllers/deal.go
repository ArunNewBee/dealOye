package controllers

import "github.com/revel/revel"

type Profile struct {
	*revel.Controller
}

func (c Profile) DealProfile() revel.Result {
	return c.Render()
}