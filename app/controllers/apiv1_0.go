package controllers
import (
    "github.com/revel/revel"
    //"github.com/gocql/gocql"
    //"gopkg.in/mgo.v2"
   // "gopkg.in/mgo.v2/bson"
    "chaokaset-go/app/models"
    "golang.org/x/crypto/bcrypt"
   // "time"
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
type ResAccounts struct {
    Status            bool
    AccountDatas      []models.Account
}
type ResAccount struct {
    Status            bool
    AccountData      *models.Account
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

func (c Api) ApiGetUserData(Username string) revel.Result {
    var R *ResAuth
    var U *models.User
    U = models.GetUserData(Username)
    if U.Username != ""{
      R = &ResAuth{Status: true,UserData: U}
    }else{
      R = &ResAuth{Status: false}
    }
    return  c.RenderJson(R)
}
//PostRegister หน้าที่ไช้สำหรับรับค่าจากฟอร์ม Register
func (c Api) PostRegisterUser(Username string,Password string,Prefix string,Name string,Lastname string,Tel string,Email string,Validpassword string) revel.Result {
  //resUserData := models.GetUserData(user.Username)
    var R *ResAuth
    var U *models.User
    HashedPassword, _ := bcrypt.GenerateFromPassword(
  		[]byte(Password), bcrypt.DefaultCost)
    Role := 3 //เกษตรกร
    err := models.RegisterUserChaokaset(Username,HashedPassword,Prefix ,Name ,Lastname ,Tel,Role,Email);
    if err {
      U = models.GetUserData(Username)
      R = &ResAuth{Status: true,UserData: U}
      return c.RenderJson(R)
    } else {
      R = &ResAuth{Status: false}
      return c.RenderJson(R)
    }
}
//PostEditUser for Create routing Page
func (c Api) PostEditUser(user *models.UserData) revel.Result {
  var R *ResAuth
  var U *models.User
  err := models.EditUserData(user.Username,user.Prefix,user.Name,user.Lastname,user.Tel,user.Email,user.Province,user.Tumbon,user.Aumphur,user.Zipcode,user.Address)
  if err {
    U = models.GetUserData(user.Username)
    R = &ResAuth{Status: err,UserData: U}
    return c.RenderJson(R)
  } else {
    R = &ResAuth{Status: err}
    return c.RenderJson(R)
  }
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

func (c Api) AddProduct(name string,category string, price int, unit string, detail string, expire string, ownerId string, lat float64, long float64) revel.Result {
 err := models.AddSellData2(name,category,price,unit,detail,expire,ownerId,lat,long)
  if err {
      return  c.RenderJson(true)
    } else {
      return  c.RenderJson(false)
    }

    //return c.RenderJson(A)

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

//AddCrop (POST)
func (c Api) AddCrop(cropid string) revel.Result {
  Result := models.GetOneCrops(cropid)
    return c.RenderJson(Result)
}

//DisabledOneCrop (GET)
func (c Api) DisabledOneCrop(cropid string) revel.Result {
  Result := models.DisableOneCrops(cropid)
    return c.RenderJson(Result)
}

//------------------ บัญชีการเพาะปลูก -------------------
//AllAccount (GET)
func (c Api) AllAccount(idcrop string,skip int) revel.Result {
  var R *ResAccounts
  Result,err := models.GetAllAccounts(idcrop,skip)
  if err == true{
    R = &ResAccounts{Status: err,AccountDatas: Result}
    return  c.RenderJson(R)
  }else{
    R = &ResAccounts{Status: err,AccountDatas: Result}
    return c.RenderJson(R)
  }
}
//OneAccount (GET)
func (c Api) OneAccount(idcrop string,idaccount string) revel.Result {
  var R *ResAccount
  Result,err := models.GetOneAccount(idcrop,idaccount)
  if err == true{
    R = &ResAccount{Status: err,AccountData: Result}
    return  c.RenderJson(R)
  }else{
    R = &ResAccount{Status: err,AccountData: Result}
    return c.RenderJson(R)
  }
}
func (c Api) SaveAccount(idcrop string,typeaccount int,detail string,price float64) revel.Result {
  var R *ResAccounts
  Adatas := models.SaveAccount(idcrop,typeaccount,detail,price);
  R = &ResAccounts{Status: true,AccountDatas: Adatas}
  if R.Status {
    return  c.RenderJson(R)
  }else{
    R = &ResAccounts{Status: false}
    return  c.RenderJson(R)
  }
}
func (c Api) EditAccount(idaccount string,detail string,price float64) revel.Result {
  var R *ResAccounts
  Adatas := models.UpdateAccount(idaccount,detail,price);
  R = &ResAccounts{Status: Adatas}
  if R.Status {
    return  c.RenderJson(R)
  }else{
    R = &ResAccounts{Status: false}
    return  c.RenderJson(R)
  }
}
func (c Api) RemoveAccount(idaccount string) revel.Result {
  var R *ResAccounts
  Adatas := models.DisableOneAccount(idaccount);
  R = &ResAccounts{Status: Adatas}
  if R.Status {
    return  c.RenderJson(R)
  }else{
    R = &ResAccounts{Status: false}
    return  c.RenderJson(R)
  }
}
func (c Api) SetStatusSell(idSell string,status int) revel.Result {
  err := models.UpdateStatusSell(idSell,status)
  if err {
    return  c.RenderJson(true)
  } else {
    return  c.RenderJson(false)
  }
}
