package controllers

import (
    "github.com/revel/revel"
    //"github.com/gocql/gocql"
    //"gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"
	  "chaokaset-go/app/models"
   // "golang.org/x/crypto/bcrypt"
    //"fmt"
   // "time"
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

func (c Sell) ProductDetail(id string) revel.Result {

  data := models.GetSellDetail(id)
  return c.Render(data)
}

func (c Sell) Sell() revel.Result {
 /* data := models.GetUserData(c.Session["username"])
  
  return  c.RenderJson(data)*/
  return c.Render()
}

func (c Sell) PostSell(sell *models.Sell) revel.Result {
  c.Validation.Required(sell.Name).Message("จำเป็นต้องกรอก ชื่อสินค้า")
  c.Validation.Required(sell.Price).Message("จำเป็นต้องกรอก ราคาสินค้าเป็นตัวเลข")
  c.Validation.Required(sell.Unit).Message("จำเป็นต้องกรอก หน่วยสินค้า")
  //c.Validation.Required(sell.Pic).Message("จำเป็นต้องกรอก รูปสินค้า")
  c.Validation.Required(sell.Detail).Message("จำเป็นต้องกรอก รายละเอียดสินค้า")
  c.Validation.Required(sell.Expire).Message("จำเป็นต้องกรอก วันหมดอายุ/ปิดการขายสินค้า")

  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect(Sell.Sell)
  }
  data := models.GetUserData(c.Session["username"])
  
  //return  c.RenderJson(data.Userid.Hex)
  err := models.AddSellData(sell.Name, sell.Category, sell.Price, sell.Unit, sell.Detail, sell.Expire, data.Userid)
  if err {
     // c.Flash.Success("สมัครสมาชิกสำเร็จ")
      return  c.Redirect(Sell.IndexSell)
    } else {
      c.Flash.Error("เกิดข้อผิดพลาดไม่สามารถขายสินค้าได้ กรุณากรอกข้อมูลใหม่!!")
      return c.Redirect(Sell.Sell)
    }

}