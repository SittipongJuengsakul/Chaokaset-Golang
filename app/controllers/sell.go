package controllers

import (
    "github.com/revel/revel"
	  "chaokaset-go/app/models"
    //"net/http"
    "io"
    "bytes"
    "mime/multipart"
    "fmt"
    "os"
    "net/http"
    "io/ioutil"

)

//Auth for save Structure of Folder Sell (in views)
type Sell struct {
  *revel.Controller
}


type PostSell struct{
  Name            string
  Category        string
  Pic             []byte
  Price           int
  Unit            string
  Detail          string
  Expire          string
}

func (c Sell) IndexSell() revel.Result {
  //var data *models.Sell
  //data := models.GetSellData(13,100)
  //data := "5555"
  return c.Render()
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

func (c Sell) PostSell(sell PostSell) revel.Result {
  /*c.Validation.Required(sell.Name).Message("จำเป็นต้องกรอก ชื่อสินค้า")
  c.Validation.Required(sell.Price).Message("จำเป็นต้องกรอก ราคาสินค้าเป็นตัวเลข")
  c.Validation.Required(sell.Unit).Message("จำเป็นต้องกรอก หน่วยสินค้า")
  //c.Validation.Required(sell.Pic).Message("จำเป็นต้องกรอก รูปสินค้า")
  c.Validation.Required(sell.Detail).Message("จำเป็นต้องกรอก รายละเอียดสินค้า")
  c.Validation.Required(sell.Expire).Message("จำเป็นต้องกรอก วันหมดอายุ/ปิดการขายสินค้า")

  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect(Sell.Sell)
  }*/
 // headers := c.Params.Files["sell.Pic"]

 headers := c.Params.Files["sell.Pic"]

   // target_url := "http://188.166.245.68:9000/public/img/pic"
    //filename := "./astaxie.pdf"
    //postFile(headers, target_url)
  //out io.Writer 
  //out := "/public/img/sell/555.jpg" 
  //out := os.OpenFile("./public/img/sell/555.jpg", 0666)
  //io.Copy(fileWriter,fn)
  
  return c.RenderJson(headers)
  
 /* data := models.GetUserid(c.Session["username"])
  
  err := models.AddSellData2(sell.Name,sell.Category,sell.Price,sell.Unit,sell.Detail,sell.Expire,data.Userid.Hex(),13,100,2)
  if err {
      return  c.Redirect(Sell.IndexSell)
    } else {
      c.Flash.Error("เกิดข้อผิดพลาดไม่สามารถขายสินค้าได้ กรุณากรอกข้อมูลใหม่!!")
      return c.Redirect(Sell.Sell)
    }*/

}

func (c Sell) ManageSell() revel.Result {
  userid := c.Session["username"]
  id := models.GetUserid(userid)
  data := models.GetManageSell(id.Userid.Hex())
  return c.Render(data)
}

func postFile(filename string, targetUrl string) error {
    bodyBuf := &bytes.Buffer{}
    bodyWriter := multipart.NewWriter(bodyBuf)

    // this step is very important
    fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
    if err != nil {
        fmt.Println("error writing to buffer")
        return err
    }

    // open file handle
    fh, err := os.Open(filename)
    if err != nil {
        fmt.Println("error opening file")
        return err
    }

    //iocopy
    _, err = io.Copy(fileWriter, fh)
    if err != nil {
        return err
    }

    contentType := bodyWriter.FormDataContentType()
    bodyWriter.Close()

    resp, err := http.Post(targetUrl, contentType, bodyBuf)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    resp_body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    fmt.Println(resp.Status)
    fmt.Println(string(resp_body))
    return nil
}

