package controllers

import (
    "github.com/revel/revel"
    //"github.com/gocql/gocql"
		"chaokaset-go/app/models"
    "golang.org/x/crypto/bcrypt"
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
  //models.RegisterUserChaokaset("username" ,"password" ,"prefix" ,"name" ,"lastname" ,"tel")
	return c.Render()
}

func (c Auth) PostRegister(user *models.User) revel.Result {
  user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
  err := models.RegisterUserChaokaset(user.Username ,user.HashedPassword,user.Prefix ,user.Name ,user.Lastname ,user.Tel);
  if err {
    c.Flash.Success("เข้าสู่ระบบสำเร็จ")
    return c.Redirect(App.Index)
  } else {
    c.Flash.Error("เกิดข้อผิดพลาด 1A001 ไม่สามารถสมัครสมาชิกได้ กรุณาสมัครไหม่!!")
    return c.Redirect(Auth.Register)
  }
}
