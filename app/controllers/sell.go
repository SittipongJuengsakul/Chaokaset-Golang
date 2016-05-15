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

//  "github.com/nfnt/resize"
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


func (c Sell) PostSell(sell PostSell,file io.Reader) revel.Result {
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

 // headers := c.Params.Files["file"]
 
   // target_url := "http://188.166.245.68:9000/public/img/pic"
    //filename := "public/img/pic"
    //postFile(headers, target_url)
  //out io.Writer 
  //out := "/public/img/sell/555.jpg" 
  //out := os.OpenFile("./public/img/sell/555.jpg", 0666)
  //io.Copy(fileWriter,fn)

  //err := c.Params.Files["file"]
 // e := os.File("file")
  //f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)
        
       // fmt.Fprintf(w, "%v", handler.Header)
      /*  f, err := os.OpenFile("./public/img/pic", os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            panic(err)
        }
        defer f.Close()*/
     //m := c.Request.MultipartForm
 
  
       //  m :=c.Params.Files["file"]
  /*       headers, ok := c.Params.Files["file"]
  if ok {

    var (
      err         error
      //bildBuf     bytes.Buffer
      nameSplit   []string
      suffix      string
      uploadDaten image.Image
    )

    nameSplit = strings.Split(strings.ToLower(headers[0].Filename), ".")
    if len(nameSplit) <= 1 {
      //return c.RenderError(SaveFotoError{msg: "no file suffix."})
      return c.RenderJson(1111111111)
    }
    suffix = nameSplit[len(nameSplit)-1]

    switch {
    case suffix == "jpg" || suffix == "jpeg":
      uploadDaten, err = jpeg.Decode(file)
      if err != nil {
        panic(err)
       // return c.RenderError(SaveFotoError{"jpeg.Decode() failed", err})
      }
    case suffix == "png":
      uploadDaten, err = png.Decode(file)
      if err != nil {
        panic(err)
      }
    case suffix == "gif":
      uploadDaten, err = gif.Decode(file)
      if err != nil {
        panic(err)
      }

   resized := resize.Resize(500, 0, uploadDaten, resize.Lanczos3)

    err = png.Encode(&bildBuf, resized)
    if err != nil {
      panic(err)
    }
    f.Data = bildBuf.Bytes()
    f.Width = resized.Bounds().Dx()
    f.Height = resized.Bounds().Dy()
return c.RenderJson(uploadDaten)
  }}*/

        
        //io.Copy(f, io.Reader(c.Params.Files["file"]))
  //target_url := "http://188.166.245.68:9000/" 
  body_buf := bytes.NewBufferString("") 
  body_writer := multipart.NewWriter(body_buf) 
  filename := "rice1.jpg" 
  file_writer, err := body_writer.CreateFormFile("file", filename) 
  if err != nil { 
   panic(err)
  }

  return c.RenderJson(file_writer)
  
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

func uploadHandler(w http.ResponseWriter, r *http.Request) {

  // the FormFile function takes in the POST input id file
  file, header, err := r.FormFile("sell.Pic")

  if err != nil {
    fmt.Fprintln(w, err)
    return
  }

  defer file.Close()

  out, err := os.Create("/public/img/pic")
  if err != nil {
    fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
    return
  }

  defer out.Close()

  // write the content from POST to the file
  _, err = io.Copy(out, file)
  if err != nil {
    fmt.Fprintln(w, err)
  }

  fmt.Fprintf(w, "File uploaded successfully : ")
  fmt.Fprintf(w, header.Filename)
  //return w
 }

