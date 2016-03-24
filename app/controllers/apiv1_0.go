package controllers

import (
    "github.com/revel/revel"
    //"github.com/gocql/gocql"
		"chaokaset-go/app/models"
    //"golang.org/x/crypto/bcrypt"
)

//Auth for save Structure of Folder Authen (in views)
type Api struct {
	*revel.Controller
}
type ResAuth struct {
    Status      bool
    Username    string
}

func (c Api) Index() revel.Result {
	return c.Render()
}
func (c Api) CheckLogin(Username string,Password string) revel.Result {
    var R *ResAuth
    res := models.CheckPasswordUser(Username,Password)
    R = &ResAuth{Status: res,Username: Username}
    return  c.RenderJson(R)
}
