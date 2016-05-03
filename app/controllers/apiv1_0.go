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
type ResSellAll struct {
    Status      bool
    SellData    []models.Sell
}
type ResSellDetail struct {
    Status      bool
    SellData    *models.SellDetail
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
func (c Api) RegisterUser(Username string,Password string,Prefix string,Name string,Lname string,Tel string,Role_user int) revel.Result {
  var R *ResAuth
  var U *models.User
  HashedPassword, _ := bcrypt.GenerateFromPassword(
    []byte(Password), bcrypt.DefaultCost)
  res := models.RegisterUserChaokaset(Username,HashedPassword,Prefix ,Name ,Lname ,Tel,Role_user); //s
  if res {
    U = models.GetUserData(Username)
  }
  R = &ResAuth{Status: res,UserData: U}
  return  c.RenderJson(R)
}

func (c Api) ProductSell() revel.Result {
  var R *ResSellAll
  var U []models.Sell
  U = models.GetSellData()
  if U == nil{
    R = &ResSellAll{Status: false,SellData: nil}
    return  c.RenderJson(R)
  }

  /*for i := range U {

    U[i].SetDistance(U[i].Address.Lat)
  }*/
  
  R = &ResSellAll{Status: true,SellData: U}
  return  c.RenderJson(R)
}

func (c Api) ProductDetail(Id string) revel.Result{
  var R *ResSellDetail
  var U *models.SellDetail
  U = models.GetSellDetail(Id)
  if U == nil{
    R = &ResSellDetail{Status: false,SellData: nil}
    return  c.RenderJson(R)
  }
  R = &ResSellDetail{Status: true,SellData: U}
  return  c.RenderJson(R)
}

func (c Api)  SearchProduct(Name string, Lat float64, Long float64) revel.Result{
  var R *ResSellAll
  var U []models.Sell
  U = models.GetSearchSell(Name,Lat,Long)
  if U == nil{
    R = &ResSellAll{Status: false,SellData: nil}
    return  c.RenderJson(R)
  }
  R = &ResSellAll{Status: true,SellData: U}
  return  c.RenderJson(R)
}
