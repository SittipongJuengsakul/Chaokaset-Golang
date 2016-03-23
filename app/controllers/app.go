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
func init() {
	revel.InterceptFunc(setuser, revel.BEFORE, &App{})
}
func (c App) connected() *models.User {
  if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	return nil
}

func setuser(c *revel.Controller) revel.Result {
	var user *models.User
  if username, ok := c.Session["username"]; ok {
		user = models.GetUserData(username)
    c.RenderArgs["user"] = user
	} else{
    c.Flash.Error("ยังไม่ล็อคอิน!!")
  }
	return nil
}

//Index for Create routing Page Index (localhost/index)
func (c App) Index(user *models.User) revel.Result {
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
  //user := models.GetUserData("sittipong")
  //c.RenderArgs["user"] = user
	return c.Render()
}
func (c Auth) PostLogin(user *models.User) revel.Result {
  result := models.CheckPasswordUser(user.Username,user.Password)
  if result {
    c.Session["username"] = user.Username
    c.Flash.Success("เข้าสู่ระบบสำเร็จ")
    return c.Redirect(App.Index)
  } else {
    c.Flash.Error("ชื่อผู้ใช้ หรือรหัสผ่านผิดพลาด!!")
    return c.Redirect(Auth.Login)
  }
	return c.Render()
}
//Logout for Create routing Page Login (localhost/Logout)
func (c Auth) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k) //ลบ session ทั้งหมด
	}
	return c.Redirect(App.Index)
}
//Register for Create routing Page Register (localhost/register)
func (c Auth) Register() revel.Result {
	return c.Render()
}
//PostRegister หน้าที่ไช้สำหรับรับค่าจากฟอร์ม Register
func (c Auth) PostRegister(user *models.User) revel.Result {
  user.HashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
  err := models.RegisterUserChaokaset(user.Username ,user.HashedPassword,user.Prefix ,user.Name ,user.Lastname ,user.Tel);
  if err {
    c.Flash.Success("สมัครสมาชิกสำเร็จ")
    return c.Redirect(App.Index)
  } else {
    c.Flash.Error("เกิดข้อผิดพลาดไม่สามารถสมัครสมาชิกได้ กรุณาสมัครไหม่!!")
    return c.Redirect(Auth.Register)
  }
}
