package controllers

import (
    "github.com/revel/revel"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
		"chaokaset-go/app/models"
    "golang.org/x/crypto/bcrypt"
    "regexp"
    "log"
)
var userRegex = regexp.MustCompile("^\\w*$")
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
//Crops
type Crops struct {
	*revel.Controller
}

func init() {
	revel.InterceptFunc(setuser, revel.BEFORE, &App{})
  revel.InterceptFunc(setuser, revel.BEFORE, &Crops{})
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
	}
	return nil
}

//Index for Create routing Page Index (localhost/index)
func (c App) Index() revel.Result {
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()
  var result *models.User
  mgod := session.DB("chaokaset").C("users")
  err = mgod.Find(bson.M{"name": "sittipong"}).One(&result)
  if err != nil {
     log.Fatal(err)
  }
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
  c.Validation.Required(user.Username).Message("จำเป็นต้องกรอก ชื่อผู้ใช้งาน")
  c.Validation.Required(user.Password).Message("จำเป็นต้องกรอก รหัสผ่าน")
  if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Auth.Login)
	}
  result := models.CheckPasswordUser(user.Username,user.Password)
  if result {
    c.Session["username"] = user.Username
    user = models.GetUserData(user.Username)
    c.RenderArgs["user"] = user
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
func (c Auth) PostRegister(user *models.User,Validpassword string) revel.Result {
	user.Validate(c.Validation)
  if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Auth.Register)
	} else{
    user.HashedPassword, _ = bcrypt.GenerateFromPassword(
  		[]byte(user.Password), bcrypt.DefaultCost)
    err := models.RegisterUserChaokaset(user.Username ,user.HashedPassword,user.Prefix ,user.Name ,user.Lastname ,user.Tel);
    if err {
      c.Flash.Success("สมัครสมาชิกสำเร็จ")
      c.Session["username"] = user.Username
      user = models.GetUserData(user.Username)
      c.RenderArgs["user"] = user
      return c.Redirect(App.Index)
    } else {
      c.Flash.Error("เกิดข้อผิดพลาดไม่สามารถสมัครสมาชิกได้ กรุณาสมัครไหม่!!")
      return c.Redirect(Auth.Register)
    }
  }
}

//IndexCrops หน้าหลักของการจัดการการเพาะปลูก
func (c Crops) IndexCrops() revel.Result {
	return c.Render()
}

//Management แสดงข้อมูลการเพาะปลูก
func (c Crops) Management() revel.Result {
	return c.Render()
}
//Account แสดงข้อมูลการเพาะปลูก
func (c Crops) Account() revel.Result {
	return c.Render()
}
//Problem แสดงข้อมูลการเพาะปลูก
func (c Crops) Problem() revel.Result {
	return c.Render()
}
//Problem แสดงข้อมูลการเพาะปลูก
func (c Crops) Product() revel.Result {
	return c.Render()
}
//Board แสดงข้อมูลการเพาะปลูก
func (c Crops) Board() revel.Result {
	return c.Render()
}
