package controllers

import "github.com/revel/revel"

func init() {
	// revel.OnAppStart(InitDB)
	// revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	// revel.InterceptMethod(Application.AddUser, revel.BEFORE)
	revel.InterceptMethod(Deal.checkUser, revel.BEFORE)
	// revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	// revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}
