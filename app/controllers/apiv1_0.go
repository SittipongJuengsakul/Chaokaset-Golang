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
func (c Api) Index() revel.Result {
  var user *models.User
	user = models.GetUserData("sittipongss")
  /*
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  //var result *models.User
  qmgo := session.DB("chaokaset").C("users")
  result := User{}
	err = qmgo.Find(bson.M{"username": "sittipong"}).One(&result)
  */
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
func (c Api) RegisterUser(Username string,Password string,Prefix string,Name string,Lname string,Tel string) revel.Result {
  var R *ResAuth
  var U *models.User
  HashedPassword, _ := bcrypt.GenerateFromPassword(
    []byte(Password), bcrypt.DefaultCost)
  res := models.RegisterUserChaokaset(Username ,HashedPassword,Prefix ,Name ,Lname ,Tel);
  if res {
    U = models.GetUserData(Username)
  }
  R = &ResAuth{Status: res,UserData: U}
  return  c.RenderJson(R)
}
