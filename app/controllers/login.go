package controllers

import "github.com/revel/revel"

type Login struct {
	*revel.Controller
}

func (c Login) UserLogin() revel.Result {
	return c.Render()
}

func (c Login) DoLogin() revel.Result {
	txtmail := c.Params.Get("txtmail")
	c.Validation.Required(txtmail).Message("Please enter a email id")
	txtpass := c.Params.Get("txtpass")
	c.Validation.Required(txtpass).Message("Please enter password")
	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Login.UserLogin)
	}
	return c.Redirect(App.Index)
}
