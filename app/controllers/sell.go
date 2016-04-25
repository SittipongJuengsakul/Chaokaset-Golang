package controllers

import (
    "github.com/revel/revel"
    //"github.com/gocql/gocql"
    //"gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"
	  "chaokaset-go/app/models"
   // "golang.org/x/crypto/bcrypt"
)

//Auth for save Structure of Folder Sell (in views)
type Sell struct {
  *revel.Controller
}



func (c Sell) IndexSell() revel.Result {
  //var data *models.Sell
  data := models.GetSellData()
  //data := "5555"
  return c.Render(data)
}

func (c Sell) ProductDetail() revel.Result {
  name := "test"
  return c.Render(name)
}

func (c Sell) Sell() revel.Result {
  return c.Render()
}