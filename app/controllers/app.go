package controllers

import (
	"dealOye/app/models"
	"dealOye/app/routes"

	"fmt"

	"github.com/revel/revel"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	//ma := make(map[string]string)
	var loggedOut, loggedIn bool
	if user := c.connected(); user == nil {
		loggedOut = true
	} else {
		loggedIn = true
	}
	return c.Render(loggedIn, loggedOut)
}

func (c Application) Login() revel.Result {
	return c.Render()
}

func (c Application) DoLogin(username, password string, remember bool) revel.Result {

	var user *models.User
	user = models.GetUser(username)
	fmt.Println(user.Password)
	fmt.Println(password)
	if user != nil {
		// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		// if err != nil {
		// 	fmt.Println(err)
		// }
		if password == user.Password {
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
