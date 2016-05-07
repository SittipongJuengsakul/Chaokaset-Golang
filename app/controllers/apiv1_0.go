package controllers

import (
    "github.com/revel/revel"
    //"github.com/gocql/gocql"
    //"gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"
		"chaokaset-go/app/models"
    "golang.org/x/crypto/bcrypt"
)

//Auth for save Structure of Folder Authen (in views)
type Api struct {
	*revel.Controller
}
type ResAuth struct {
    Status      bool
    UserData    *models.User
}
type ResPlan struct {
    Status      bool
    PlanData    *models.Plan
}
func (c Api) Index() revel.Result {
  var user *models.User
	user = models.GetUserData("sittipong")
	return c.Render(user)
}
func (c Api) CheckLogin(Username string,Password string) revel.Result {
    var R *ResAuth
    var U *models.User
    res := models.CheckPasswordUser(Username,Password)
    if res {
      U = models.GetUserData(Username)
    }
    R = &ResAuth{Status: res,UserData: U}
    return  c.RenderJson(R)
}
func (c Api) RegisterUser(Username string,Password string,Prefix string,Name string,Lname string,Tel string,Role_user int,Email string) revel.Result {
  var R *ResAuth
  var U *models.User
  HashedPassword, _ := bcrypt.GenerateFromPassword(
    []byte(Password), bcrypt.DefaultCost)
  res := models.RegisterUserChaokaset(Username,HashedPassword,Prefix ,Name ,Lname ,Tel,Role_user,Email); //s
  if res {
    U = models.GetUserData(Username)
  }
  R = &ResAuth{Status: res,UserData: U}
  return  c.RenderJson(R)
}

//------------------ แผนการเพาะปลูก -------------------
//Plan (GET)
func (c Api) Plans(skip int,word string) revel.Result {
  Result,err := models.GetAllPlans(skip)
  if err == true{
    return  c.RenderJson(Result)
  }else{
    return c.RenderJson(Result)
  }
}

//------------------ พืชและพันธุ์พืช -------------------
//Plan (GET)
func (c Api) Plans(skip int,word string) revel.Result {
  Result,err := models.GetAllPlans(skip)
  if err == true{
    return  c.RenderJson(Result)
  }else{
    return c.RenderJson(Result)
  }
}
