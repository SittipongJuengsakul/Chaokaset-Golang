package controllers
import (
    "github.com/revel/revel"
    //"github.com/gocql/gocql"
    //"gopkg.in/mgo.v2"
   // "gopkg.in/mgo.v2/bson"
    "chaokaset-api/app/models"
    "golang.org/x/crypto/bcrypt"
   // "time"
    "log"
    "time"
    "os"
    "io"
   "fmt"
)

type Api struct {
  *revel.Controller
}

type ResAuth struct {
    Status      bool
    UserData    *models.User
}
type ResSellAll struct {
    Status        bool
    SellData      []models.Sells
}
type ResSellDetail struct {
    Status      bool
    SellData    *models.SellDetail
}

type ResPlan struct {
    Status      bool
    PlanData    *models.Plan
}
type ResPlans struct {
    Status      bool
    PlanData    []models.Plan
}
type ResSeed struct {
    Status      bool
    SeedData    *models.Seed
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

func (c Api) ProductSell(Lat float64, Long float64) revel.Result {
  var R *ResSellAll
  var U []models.Sells
  U = models.GetSellData(Lat,Long)
  if U == nil{
    R = &ResSellAll{Status: false,SellData: nil}
    return  c.RenderJson(R)
  }

  R = &ResSellAll{Status: true,SellData: U}
  return  c.RenderJson(R)
}

func (c Api) ProductDetail(IdSell string,IdUser string) revel.Result{
  var R *ResSellDetail
  var U *models.SellDetail
  U = models.GetSellDetail(IdSell,IdUser)
  if U == nil{
    R = &ResSellDetail{Status: false,SellData: nil}
    return  c.RenderJson(R)
  }
  R = &ResSellDetail{Status: true,SellData: U}
  return  c.RenderJson(R)
}

func (c Api) SearchProduct(Name string, Lat float64, Long float64) revel.Result{
 var R *ResSellAll
  var U []models.Sells
  U = models.GetSearchSell(Name,Lat,Long)
  if U == nil{
    R = &ResSellAll{Status: false,SellData: nil}
    return  c.RenderJson(R)
  }
  R = &ResSellAll{Status: true,SellData: U}
  return  c.RenderJson(R)
}

func (c Api) AddProduct(name string,category string, price int, unit string, detail string, expire string, ownerId string, lat float64, long float64,sellType int) revel.Result {
 err := models.AddSellData2(name,category,price,unit,detail,expire,ownerId,lat,long,sellType)
  /*if err {
      return  c.RenderJson(err)
    } else {
      return  c.RenderJson(err)
    }*/
    return c.RenderJson(err)
}

func (c Api) ManageSell(idUser string) revel.Result {
 var R *ResSellAll
  var U []models.Sells
  U = models.GetManageSell(idUser)
  if U == nil{
    R = &ResSellAll{Status: false,SellData: nil}
    return  c.RenderJson(R)
  }
  R = &ResSellAll{Status: true,SellData: U}
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
//Plan (GET)
func (c Api) PlansAllPlants(skip int,idplant string,idseed string) revel.Result {
  Result,err := models.GetAllPlansPlants(skip,idplant,idseed)
  if err == true{
    return  c.RenderJson(Result)
  }else{
    return c.RenderJson(Result)
  }
}

//------------------ แผนการเพาะปลูก -------------------
//Plan (GET)
func (c Api) Plan(idplan string) revel.Result {
  Result := models.GetPlans(idplan)
  var Res *ResPlan
  if(Result.PlanId == ""){
    Res = &ResPlan{Status: false}
  }else{
    Res = &ResPlan{Status: true,PlanData: Result}
  }
  return  c.RenderJson(Res)
}

//------------------ พืชและพันธุ์พืช -------------------
//Plants (GET)
func (c Api) Plants(skip int,word string) revel.Result {
  Result,err := models.GetAllPlants(skip)
  if err == nil{
    return  c.RenderJson(Result)
  }else{
    return c.RenderJson(Result)
  }
}
//Plant (GET)
func (c Api) Plant(word string) revel.Result {
  //Result := models.GetPlantId("572dfeede3890226904ecba9")
  Result := models.GetPlant(word)
    return  c.RenderJson(Result)
}
//Plants (Post)
func (c Api) SavePlantData(PlantName string) revel.Result {
  Result := models.SavePlant(PlantName);
  return c.RenderJson(Result)
}

//Seed (GET)
func (c Api) Seed(skips int,plantname string,seedname string) revel.Result {
    Result := models.GetSeed(skips,plantname,seedname)
    var Res *ResSeed
    if(Result.SeedId == ""){
      Res = &ResSeed{Status: false}
    }else{
      Res = &ResSeed{Status: true,SeedData: Result}
    }
    return  c.RenderJson(Res)
}
//Seed (GET)
func (c Api) Seeds(skips int,plantid string) revel.Result {
    return  c.RenderJson(models.GetAllSeeds(skips,plantid))
}

//RemoveSeed (GET)
func (c Api) RemoveSeed(idseed string) revel.Result {
  models.RemoveSeed(idseed);
  return c.RenderJson(models.GetAllSeeds(0,""))
}

//Get District Province ProvinceId (GET)
func (c Api) Province(provinceid string) revel.Result {
    return  c.RenderJson(models.GetProvinces(provinceid))
}

//------------------ แผนการเพาะปลูก -------------------
//AllCrops (GET)
func (c Api) AllCrops(skip int,userid string) revel.Result {
  Result,err := models.GetAllCrops(0,userid)
  if err == true{
    return  c.RenderJson(Result)
  }else{
    return c.RenderJson(Result)
  }
}
//OneCrop (GET)
func (c Api) OneCrop(cropid string) revel.Result {
  Result := models.GetOneCrops(cropid)
    return c.RenderJson(Result)
}

//OneCrop (GET)
func (c Api) DisabledOneCrop(cropid string) revel.Result {
  Result := models.DisableOneCrops(cropid)
    return c.RenderJson(Result)
}

func (c Api) SetStatusSell(idSell string,status int) revel.Result {
  err := models.UpdateStatusSell(idSell,status)
  if err {
    return  c.RenderJson(true)
  } else {
    return  c.RenderJson(false)
  }
}

func (c Api) EditProduct(idSell string, name string, category string, price int,detail string,expire string,unit string,lat float64,long float64) revel.Result {
  err := models.EditProductSell(idSell,name,category,price,detail,expire,unit,lat,long)
  if err {
    return  c.RenderJson(true)
  } else {
    return  c.RenderJson(false)
  }
}

func (c Api) PostApi(IdSell string) revel.Result {
  upload_dir := "/var/home/goserver/src/chaokaset-api/public/uploads/"
  m := c.Request.MultipartForm
  result := true
  for fname, _ := range m.File {

    fheaders := m.File[fname]
    for i, _ := range fheaders {
      //for each fileheader, get a handle to the actual file
      file, err := fheaders[i].Open()
      defer file.Close() //close the source file handle on function return
      if err != nil {
         log.Print(err)
        result = false
      }
      //create destination file making sure the path is writeable.
      t := time.Now()
      file_name_db := "mobile-" + t.Format("20060102150405") + "-" +  fheaders[i].Filename
      dst_path := upload_dir + file_name_db
      dst, err := os.Create(dst_path)
      defer dst.Close() //close the destination file handle on function return
      defer os.Chmod(dst_path, (os.FileMode)(0644)) //limit access restrictions
      if err != nil {
        log.Print(err)
       result = false
      }
      //copy the uploaded file to the destination file
      if _, err := io.Copy(dst, file); err != nil {
        log.Print(err)
       result = false
      }
      fmt.Printf("%+v\n", IdSell)
      fmt.Printf("%+v\n", file_name_db)
      models.UpdatePic(IdSell,file_name_db)
    }
  } 
  return  c.RenderJson(result)
}

func (c Api) LikeProduct(idSell string, idUser string) revel.Result {
  err := models.Like(idSell,idUser)
  if err {
    return  c.RenderJson(true)
  } else {
    return  c.RenderJson(false)
  }
}

func (c Api) UnLikeProduct(idSell string, idUser string) revel.Result {
  err := models.UnLike(idSell,idUser)
  if err {
    return  c.RenderJson(true)
  } else {
    return  c.RenderJson(false)
  }
}
