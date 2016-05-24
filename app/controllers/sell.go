package controllers

import (
    "github.com/revel/revel"
	  "chaokaset-api/app/models"
    "io"
    "fmt"
    "log"
    "os"
    "time"
  // "sort"
     "strconv"
)

//Auth for save Structure of Folder Sell (in views)
type Sell struct {
  *revel.Controller
}

//type SortData []models.Sells
type SetLatLong struct{
  Lat    string
  Long   string
}

type ByLike []models.Sells

func (a ByLike) Len() int           { return len(a) }
func (a ByLike) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLike) Less(i, j int) bool { return a[i].NumberOfLike < a[j].NumberOfLike }

type ByDistance []models.Sells

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }

func (c Sell) IndexSell() revel.Result {
  //var data *models.Sells
  //Data := models.GetSellData(13,100)
 /*for i := range SortData {
    fmt.Printf("%+v",SortData[i].NumberOfLike)
  }
*/
//  sort.Sort(sort.Reverse(ByLike(Data)))
  Lat := c.Session["Lat"]
  Long := c.Session["Long"]

  return c.Render(Lat,Long)
}

func (c Sell) ProductDetail(id string) revel.Result {
  idUser := models.GetUserid(c.Session["username"])
  data := models.GetSellDetail(id,idUser.Userid.Hex())
  
  return c.Render(data,idUser)
}

func (c Sell) Sell() revel.Result {
 /* data := models.GetUserData(c.Session["username"])

  return  c.RenderJson(data)*/
  return c.Render()
}

func (c Sell) PostSell(sell models.PostSell) revel.Result {
    
  c.Validation.Required(sell.Name).Message("จำเป็นต้องกรอก ชื่อสินค้า")
  c.Validation.Required(sell.Price).Message("จำเป็นต้องกรอก ราคาสินค้าเป็นตัวเลข")
  c.Validation.Required(sell.Unit).Message("จำเป็นต้องกรอก หน่วยสินค้า")
  //c.Validation.Required(sell.Pic).Message("จำเป็นต้องกรอก รูปสินค้า")
  c.Validation.Required(sell.Detail).Message("จำเป็นต้องกรอก รายละเอียดสินค้า")
  c.Validation.Required(sell.Expire).Message("จำเป็นต้องกรอก วันหมดอายุ/ปิดการขายสินค้า")
  var lat_value string = c.Params.Get("lat_value")
  var lon_value string = c.Params.Get("lon_value")
  fmt.Printf("%+v\n", lat_value)
  fmt.Printf("%+v\n", lon_value)

  lat_values,_ := strconv.ParseFloat(lat_value, 64)
  lon_values,_ := strconv.ParseFloat(lon_value, 64)

  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect(Sell.Sell)
  }
    
  data := models.GetUserid(c.Session["username"])

  
  selling := models.AddSellData2(sell.Name,sell.Category,sell.Price,sell.Unit,sell.Detail,sell.Expire,data.Userid.Hex(),lat_values,lon_values,2)

  fmt.Printf("%+v\n", selling.SellId)
  
  upload_dir := "/var/home/goserver/src/chaokaset-api/public/uploads/" 
  m := c.Request.MultipartForm
    //var msg string
  for fname, _ := range m.File {

  fheaders := m.File[fname]
    for i, _ := range fheaders {
      //for each fileheader, get a handle to the actual file
      file, err := fheaders[i].Open()
      defer file.Close() //close the source file handle on function return
      if err != nil {
         log.Print(err)
       //  msg = "upload failed.."
      }
      //create destination file making sure the path is writeable.
      t := time.Now()
      file_name_db := c.Session["username"] + "-" + t.Format("20060102150405") + "-" +  fheaders[i].Filename
      
      fmt.Printf("%+v\n", file_name_db)

      models.UpdatePic(selling.SellId,file_name_db)
      
      dst_path := upload_dir + file_name_db
      dst, err := os.Create(dst_path)
      defer dst.Close() //close the destination file handle on function return
      defer os.Chmod(dst_path, (os.FileMode)(0644)) //limit access restrictions
      if err != nil {
        log.Print(err)
       // msg = "upload failed.."
      }
      //copy the uploaded file to the destination file
      if _, err := io.Copy(dst, file); err != nil {
        log.Print(err)
       // msg = "upload failed.."
      }
    }
  }
   
  //return  c.RenderJson(555)
  
  return  c.Redirect(Sell.IndexSell)

}

func (c Sell) ManageSell() revel.Result {
  userid := c.Session["username"]
  id := models.GetUserid(userid)
  data := models.GetManageSell(id.Userid.Hex())
  return c.Render(data)
}

func (c Sell) EditProductSell(idSell string) revel.Result {
  data := models.GetSellDetail(idSell,"55555555")
  return c.Render(data)
}

func (c Sell) PostEditSell() revel.Result {
  return nil
}

func (c Sell) CloseSell(idSell string) revel.Result {
  models.UpdateStatusSell(idSell,0)
  return  c.Redirect(Sell.ManageSell)
}
func (c Sell) OpenSell(idSell string) revel.Result {
  models.UpdateStatusSell(idSell,1)
  return  c.Redirect(Sell.ManageSell)
}

func (c Sell) ProductCategory(Category string) revel.Result {
 // Data := models.
  return c.Render(Category)
}

func (c Sell) SetLatLong(Lat string, Long string) revel.Result {
 // Data := models.
  //return c.Render(Category)
  c.Session["Lat"] = Lat
  c.Session["Long"] = Long
  var data *SetLatLong
  data = &SetLatLong{Lat : c.Session["Lat"],Long:c.Session["Long"]}
  return c.RenderJson(data)
  //return true
}

func (c Sell) ListSellCrop() revel.Result {
  userid := c.Session["username"]
  id := models.GetUserid(userid)
  data := models.GetCropSell(id.Userid.Hex())
  return c.Render(data)
}

func (c Sell) SellCrop(idcrop string) revel.Result{
  userid := c.Session["username"]
  id := models.GetUserid(userid)
  data := models.GetCropSellDetail(id.Userid.Hex(),idcrop)
  return c.Render(data)
}

func (c Sell) PostSellCrop() revel.Result {
  var Name string = c.Params.Get("Name")
  //var Category string = c.Params.Get("Category")
  var Price string = c.Params.Get("Price")
  //var Unit string = c.Params.Get("Unit")
  var Detail string = c.Params.Get("Detail")
  var Expire string = c.Params.Get("Expire")



 /* c.Validation.Required(Name).Message("จำเป็นต้องกรอก ชื่อสินค้า")
  c.Validation.Required(Price).Message("จำเป็นต้องกรอก ราคาสินค้าเป็นตัวเลข")
  c.Validation.Required(Unit).Message("จำเป็นต้องกรอก หน่วยสินค้า")
  c.Validation.Required(Category).Message("จำเป็นต้องกรอก รูปสินค้า")
  c.Validation.Required(Detail).Message("จำเป็นต้องกรอก รายละเอียดสินค้า")
  c.Validation.Required(Expire).Message("จำเป็นต้องกรอก วันหมดอายุ/ปิดการขายสินค้า")
  
  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect("/sellcrop/%s","57435d9de3890226904eccd6")
  }*/

  data := models.GetUserid(c.Session["username"])

  Prices,_:=  strconv.Atoi(Price)

  
  selling := models.AddSellData2(Name,"ข้าว",Prices,"กิโลกรัม",Detail,Expire,data.Userid.Hex(),13,100,1)

  fmt.Printf("%+v\n", selling.SellId)
  
  upload_dir := "/var/home/goserver/src/chaokaset-api/public/uploads/" 
  m := c.Request.MultipartForm
    //var msg string
  for fname, _ := range m.File {

  fheaders := m.File[fname]
    for i, _ := range fheaders {
      //for each fileheader, get a handle to the actual file
      file, err := fheaders[i].Open()
      defer file.Close() //close the source file handle on function return
      if err != nil {
         log.Print(err)
       //  msg = "upload failed.."
      }
      //create destination file making sure the path is writeable.
      t := time.Now()
      file_name_db := c.Session["username"] + "-" + t.Format("20060102150405") + "-" +  fheaders[i].Filename
      
      fmt.Printf("%+v\n", file_name_db)

      models.UpdatePic(selling.SellId,file_name_db)
      
      dst_path := upload_dir + file_name_db
      dst, err := os.Create(dst_path)
      defer dst.Close() //close the destination file handle on function return
      defer os.Chmod(dst_path, (os.FileMode)(0644)) //limit access restrictions
      if err != nil {
        log.Print(err)
       // msg = "upload failed.."
      }
      //copy the uploaded file to the destination file
      if _, err := io.Copy(dst, file); err != nil {
        log.Print(err)
       // msg = "upload failed.."
      }
    }
  }
  
  
 // return  c.RenderJson(Name)
  
  return  c.Redirect(Sell.IndexSell)

}
