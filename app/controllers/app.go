package controllers

import (
	"dealOye/app/models"
	"dealOye/app/routes"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	return c.Render()
}

func (c Application) Login() revel.Result {
	return c.Render()
}

func (c Application) DoLogin(username, password string, remember bool) revel.Result {

	user := models.GetUser(username)

	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(routes.Application.Index())
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.Application.Login())
}

func (c Application) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return models.GetUser(username)
	}
	return nil
}
func (c Application) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.Application.Index())
}
