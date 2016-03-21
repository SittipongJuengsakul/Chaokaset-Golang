package controllers

import (
    "github.com/revel/revel"
    "golang.org/x/crypto/bcrypt" //สำหรับ hashing password
    "golang.org/x/oauth2" //สำหรับจัดการ Authen
    //"github.com/gocql/gocql"
		"chaokaset-go/app/models" //เรียกไช้ model
)

//App for save Structure of Folder App (in views)
type App struct {
	*revel.Controller
}

//Search for save Structure of Folder Search (in views)
type Search struct {
	*revel.Controller
}

//Auth for save Structure of Folder Authen (in views)
type Auth struct {
	*revel.Controller
}

//Index for Create routing Page Index (localhost/index)
func (c App) Index() revel.Result {
	return c.Render()
}

//Templates for Example Template (localhost/template)
func (c App) Templates() revel.Result {
	return c.Render()
}

//SearchPlant for Create routing Page Index (localhost/searchplant)
func (c Search) SearchPlant() revel.Result {
	return c.Render()
}

//Login for Create routing Page Login (localhost/login)
func (c Auth) Login() revel.Result {
	//models.RegisterUserChaokaset("dddddd","jungsakul","0839915593","123456")
	return c.Render()
}

//Register for Create routing Page Register (localhost/register)
func (c Auth) Register() revel.Result {
	return c.Render()
}

func (c Auth) PostRegister(user models.User) revel.Result {
  user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
  models.RegisterUserChaokaset(user.Username,user.HashedPassword,user.Prefix,user.Name,user.Lastname,user.Tel)
  //c.Session["user"] = user.Username
	//c.Flash.Success("Welcome, " + user.Name)
  return c.Redirect("/index")
}
