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

func (c Sell) PostSell(sell *models.Sell) revel.Result {
  c.Validation.Required(sell.Name).Message("จำเป็นต้องกรอก ชื่อสินค้า")
  c.Validation.Required(sell.Price).Message("จำเป็นต้องกรอก ราคาสินค้า")
  c.Validation.Required(sell.Unit).Message("จำเป็นต้องกรอก หน่วยสินค้า")
  c.Validation.Required(sell.Pic).Message("จำเป็นต้องกรอก รูปสินค้า")
  c.Validation.Required(sell.Detail).Message("จำเป็นต้องกรอก รายละเอียดสินค้า")
  c.Validation.Required(sell.Expire).Message("จำเป็นต้องกรอก วันหมดอายุ/ปิดการขายสินค้า")
  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect(Sell.Sell)
  }
 /* result := models.CheckPasswordUser(user.Username,user.Password)
  if result {
    c.Session["username"] = user.Username
    user = models.GetUserData(user.Username)
    c.RenderArgs["user"] = user
    c.Flash.Success("เข้าสู่ระบบสำเร็จ")
    return c.Redirect(App.Index)
  } else {
    c.Flash.Error("ชื่อผู้ใช้ หรือรหัสผ่านผิดพลาด!!")
    return c.Redirect(Auth.Login)
  }*/
  return c.Render()
}