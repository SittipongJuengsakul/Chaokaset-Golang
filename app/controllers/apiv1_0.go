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


func (c Api) Index(user *models.User) revel.Result {
	return c.Render()
}
func (c Api) CheckLogin(Username string,Password string) revel.Result {
    res := models.CheckPasswordUser(Username,Password)
    if res {
      return  c.RenderJson(res)
    } else {
      return  c.RenderJson(res)
    }
}
