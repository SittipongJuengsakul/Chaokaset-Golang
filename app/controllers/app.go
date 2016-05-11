package controllers

import (
    "github.com/revel/revel"
    //"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
		"chaokaset-go/app/models"
    "golang.org/x/crypto/bcrypt"
    "regexp"
    //"log"
    //"time"
)
var userRegex = regexp.MustCompile("^\\w*$")

//App for save Structure of Folder App (in views)
type App struct {
	*revel.Controller
}
type Person struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	Username      string
	Phone         string
}
//Search for save Structure of Folder Search (in views)
type Search struct {
	*revel.Controller
}

//Auth for save Structure of Folder Authen (in views)
type Auth struct {
	*revel.Controller
}
//Profile for save Structure of Folder Profile (in views)
type Profile struct {
	*revel.Controller
}
//Howto for save Structure of Folder Howto (in views)
type Howto struct {
	*revel.Controller
}
//plan for save Structure of Folder Plan (in views)
type Plan struct {
	*revel.Controller
}
//Plant
type Plant struct {
	*revel.Controller
}
//Crops
type Crops struct {
	*revel.Controller
}
//Management
type Management struct {
	*revel.Controller
}
//Farmer
type Farmer struct {
	*revel.Controller
}


func init() {
  revel.InterceptFunc(setuser, revel.BEFORE, &App{})
  revel.InterceptFunc(setuser, revel.BEFORE, &Crops{})
  revel.InterceptFunc(setuser, revel.BEFORE, &Profile{})
  revel.InterceptFunc(setuser, revel.BEFORE, &Howto{})
  revel.InterceptFunc(setuser, revel.BEFORE, &Plan{})
  revel.InterceptFunc(setuser, revel.BEFORE, &Management{})
  revel.InterceptFunc(setuser, revel.BEFORE, &Farmer{})
  revel.InterceptFunc(setuser, revel.BEFORE, &Plant{})
  revel.InterceptFunc(checksetuser, revel.BEFORE, &Crops{})
  revel.InterceptFunc(checksetuser, revel.BEFORE, &Profile{})
}

func (c App) connected() *models.User {
  if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	return nil
}
//ตราจสอบ Session
func setuser(c *revel.Controller) revel.Result {
	var user *models.User
  if username, ok := c.Session["username"]; ok {
		user = models.GetUserData(username)
    c.RenderArgs["user"] = user
    if user.Role == 1{
      c.RenderArgs["admin"] = "ผู้ดูแลระบบ"
    } else if user.Role == 2{
      c.RenderArgs["officer"] = "เจ้าหน้าที่"
    } else if user.Role == 3{
      c.RenderArgs["kasetkorn"] = "เกษตรกร"
    } else if user.Role == 4{
      c.RenderArgs["userkaset"] = "เกษตรกร"
    }
	}
	return nil
}
//ตรวจสอบว่ามีสิทธิไช้งานหรือไม่ หากไม่ ล็อกอินก่อน
func checksetuser(c *revel.Controller) revel.Result {
	var user *models.User
  if username, ok := c.Session["username"]; ok {
		user = models.GetUserData(username)
    c.RenderArgs["user"] = user

	} else{
    c.Flash.Error("กรุณาล็อคอินก่อนทำรายการ !!")
    return c.Redirect(Auth.Login)
  }
	return nil
}

//Index for Create routing Page Index (localhost/index)
func (c App) Index() revel.Result {
  //res := models.GenString(6)
	return c.Render()
}
//AboutUs for Create routing Page AboutUs
func (c App) AboutUs() revel.Result {
	return c.Render()
}
//Social for Create routing Page Social
func (c App) Social() revel.Result {
	return c.Render()
}

//Templates for Example Template (localhost/template)
func (c App) Templates() revel.Result {
	return c.Render()
}

//SearchPlant for Create routing Page Index (localhost/searchplant)
func (c Search) SearchPlant() revel.Result {
	return c.Render()
}

//Login for Create routing Page Login (localhost/login)
func (c Auth) Login() revel.Result {
  if _, ok := c.Session["username"]; ok {
    return c.Redirect(App.Index)
	}
	return c.Render()
}
func (c Auth) PostLogin(user *models.User) revel.Result {
  c.Validation.Required(user.Username).Message("จำเป็นต้องกรอก ชื่อผู้ใช้งาน")
  c.Validation.Required(user.Password).Message("จำเป็นต้องกรอก รหัสผ่าน")
  if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Auth.Login)
	}
  result := models.CheckPasswordUser(user.Username,user.Password)
  if result {
    c.Session["username"] = user.Username
    user = models.GetUserData(user.Username)
    c.RenderArgs["user"] = user
    //c.Session["userid"] = user.Userid 
    c.Flash.Success("เข้าสู่ระบบสำเร็จ")
    return c.Redirect(App.Index)
  } else {
    c.Flash.Error("ชื่อผู้ใช้ หรือรหัสผ่านผิดพลาด!!")
    return c.Redirect(Auth.Login)
  }
	return c.Render()
}
//Logout for Create routing Page Login (localhost/Logout)
func (c Auth) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k) //ลบ session ทั้งหมด
	}
	return c.Redirect(App.Index)
}
//Register for Create routing Page Register (localhost/register)
func (c Auth) Register() revel.Result {
  if _, ok := c.Session["username"]; ok {
    return c.Redirect(App.Index)
	}
	return c.Render()
}
//PostRegister หน้าที่ไช้สำหรับรับค่าจากฟอร์ม Register
func (c Auth) PostRegister(user *models.User,Validpassword string) revel.Result {
  resUserData := models.GetUserData(user.Username)
  c.Validation.Required(user.Username != resUserData.Username).Message("ชื่อผู้ไช้มีคนไช้งานแล้ว")
	user.Validate(c.Validation)
  if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Auth.Register)
	} else{
    user.HashedPassword, _ = bcrypt.GenerateFromPassword(
  		[]byte(user.Password), bcrypt.DefaultCost)
    err := models.RegisterUserChaokaset(user.Username ,user.HashedPassword,user.Prefix ,user.Name ,user.Lastname ,user.Tel,user.Role,user.Email);
    if err {
      c.Flash.Success("สมัครสมาชิกสำเร็จ")
      c.Session["username"] = user.Username
      user = models.GetUserData(user.Username)
      c.RenderArgs["user"] = user
      return c.Redirect(App.Index)
    } else {
      c.Flash.Error("เกิดข้อผิดพลาดไม่สามารถสมัครสมาชิกได้ กรุณาสมัครไหม่!!")
      return c.Redirect(Auth.Register)
    }
  }
}
//ChangePassword for Create routing Page ChangePassword
func (c Profile) ChangePassword() revel.Result {
	return c.Render()
}
//EditUser for Create routing Page Register
func (c Profile) EditUser() revel.Result {
  var user *models.UserData
  user = models.GetEditUserData(c.Session["username"])
	return c.Render(user)
}
//PostEditUser for Create routing Page Register
func (c Profile) PostEditUser(user *models.UserData) revel.Result {
	user.ValidateUserData(c.Validation)
  if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Profile.EditUser)
  }
  result := models.EditUserData(c.Session["username"],user.Prefix,user.Name,user.Lastname,user.Tel,user.Email,user.Province,user.Tumbon,user.Aumphur,user.Zipcode,user.Address)
  if result {
    return c.Redirect(App.Index)
  } else{
    return c.Redirect(Profile.EditUser)
  }
}
//SettingUser for Create routing Page Register
func (c Profile) SettingUser() revel.Result {
	return c.Render()
}
//SettingSecurity for Create routing Page Register
func (c Profile) SettingSecurity() revel.Result {
	return c.Render()
}
//IndexCrops หน้าหลักของการจัดการการเพาะปลูก
func (c Crops) IndexCrops() revel.Result {
  user := models.GetUserData(c.Session["username"])
  allcrops,_ := models.GetAllCrops(0,user.Userid.Hex())
	return c.Render(allcrops)
}

//IndexCrops หน้าหลักของการจัดการการเพาะปลูก
func (c Crops) DisbleCrop(idcrop string) revel.Result {
  str := models.DisableOneCrops(idcrop)
  if str == true{
    return c.Redirect(Crops.IndexCrops)
  }else{
    return c.Redirect(App.Index)
  }

}

//Management แสดงข้อมูลการเพาะปลูก
func (c Crops) Management(idcrop string) revel.Result {
  if idcrop == ""{
    return c.Redirect(Crops.IndexCrops)
  }
  cropdata := models.GetOneCrops(idcrop)
  plandata := models.GetPlans(cropdata.PlanId)
  cropdataproduct := cropdata.Product*20.5
	return c.Render(cropdata,plandata,cropdataproduct)
}
//Carlendar
func (c Crops) Calendar(idcrop string) revel.Result {
  if idcrop == ""{
    return c.Redirect(Crops.IndexCrops)
  }
  cropdata := models.GetOneCrops(idcrop)
  plandata := models.GetPlans(cropdata.PlanId)
	return c.Render(cropdata,plandata)
}
//Account แสดงข้อมูลการเพาะปลูก
func (c Crops) Account(idcrop string) revel.Result {
  if idcrop == ""{
    return c.Redirect(Crops.IndexCrops)
  }
  cropdata := models.GetOneCrops(idcrop)
  plandata := models.GetPlans(cropdata.PlanId)
	return c.Render(cropdata,plandata)
}
//Problem แสดงข้อมูลการเพาะปลูก
func (c Crops) Problem(idcrop string) revel.Result {
  if idcrop == ""{
    return c.Redirect(Crops.IndexCrops)
  }
  cropdata := models.GetOneCrops(idcrop)
  plandata := models.GetPlans(cropdata.PlanId)
	return c.Render(cropdata,plandata)
}
//Problem แสดงข้อมูลการเพาะปลูก
func (c Crops) Product(idcrop string) revel.Result {
  if idcrop == ""{
    return c.Redirect(Crops.IndexCrops)
  }
  cropdata := models.GetOneCrops(idcrop)
  plandata := models.GetPlans(cropdata.PlanId)
	return c.Render(cropdata,plandata)
}
//Board แสดงข้อมูลการเพาะปลูก
func (c Crops) Board(idcrop string) revel.Result {
  if idcrop == ""{
    return c.Redirect(Crops.IndexCrops)
  }
  cropdata := models.GetOneCrops(idcrop)
  plandata := models.GetPlans(cropdata.PlanId)
	return c.Render(cropdata,plandata)
}
//AddCrop เพิ่มข้อมูลการเพาะปลูก
func (c Crops) AddCrop() revel.Result {
	return c.Render()
}
//PostAddCrop เพิ่มข้อมูลการเพาะปลูก
func (c Crops) PostAddCrop(crop *models.Crop) revel.Result {
  user := models.GetUserData(c.Session["username"])
  result := models.SaveCrop(crop,user.Userid.Hex())
  if result {
    return c.Redirect(Crops.IndexCrops)
  } else{
    return c.Redirect(Crops.AddCrop)
  }
}
//AddCropPlan เพิ่มข้อมูลแผนการเพาะปลูก
func (c Plan) AddCropPlan() revel.Result {
	return c.Render()
}
//PostAddCropPlan
func (c Plan) PostAddCropPlan(plan *models.Plan) revel.Result {
  plan.CheckSavePlan(c.Validation)
  if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Plan.AddCropPlan)
  }
  result := models.SavePlan(plan)
  if result {
    return c.Redirect(Plan.AddPlanActivity)
  } else{
    return c.Redirect(Plan.AddCropPlan)
  }
	return c.Render()
}
//ShowPlant
func (c Plant) ShowPlant() revel.Result {
	return c.Render()
}
//AddPlant
func (c Plant) AddPlant() revel.Result {
	return c.Render()
}
//ShowSeed
func (c Plant) ShowSeed() revel.Result {
	return c.Render()
}
//AddSeed
func (c Plant) AddSeed() revel.Result {
	return c.Render()
}
//RemoveSeed (DELETE)
func (c Plant) RemoveSeed(idseed string) revel.Result {
  result := models.RemoveSeed(idseed);
  if(result){
    return c.Redirect(Plant.ShowSeed)
  }else{
    return c.Redirect(Plant.ShowSeed)
  }
}
//SavePlant
func (c Plant) PostAddPlant(plant *models.Plant) revel.Result {
    resPlantData := models.GetPlant(plant.PlantName)
    c.Validation.Required(plant.PlantName).Message("กรุณากรอกข้อมูลชื่อพืช")
    c.Validation.Required(plant.PlantName != resPlantData.PlantName).Message(" \""+plant.PlantName+"\" มีอยู่บนระบบแล้วกรุณาตรวจสอบ")
    if c.Validation.HasErrors() {
  		c.Validation.Keep()
  		c.FlashParams()
  		return c.Redirect(Plant.AddPlant)
  	} else{
      Result := models.SavePlant(plant.PlantName);
      if Result {
        return c.Redirect(Plant.ShowPlant)
      } else{
        return c.Redirect(Plant.AddPlant)
      }
    }
}
//PostAddSeed
func (c Plant) PostAddSeed(seed *models.Seed,plant *models.Plant) revel.Result {
    var plantdata *models.Plant
    plantdata = models.GetPlantId(plant.PlantName)
    resSeedData := models.GetSeed(0,plantdata.PlantName,seed.SeedName)
    c.Validation.Required(seed.SeedName).Message("กรุณากรอกข้อมูลชื่อพันธุ์พืช")
    c.Validation.Required(seed.SeedName != resSeedData.SeedName).Message(" \""+seed.SeedName+"\" มีอยู่บนระบบแล้วกรุณาตรวจสอบ")
    if c.Validation.HasErrors() {
  		c.Validation.Keep()
  		c.FlashParams()
  		return c.Redirect(Plant.AddSeed)
  	} else{
      Result := models.SaveSeed(seed.SeedName,plant.PlantName,plantdata.PlantName);
      if Result {
        return c.Redirect(Plant.ShowSeed)
      } else{
        return c.Redirect(Plant.AddSeed)
      }
    }
}
//AddPlanActivity
func (c Plan) AddPlanActivity() revel.Result {
	return c.Render()
}
//ShowPlan
func (c Plan) ShowPlan() revel.Result {
	return c.Render()
}
//PlanManagement
func (c Plan) PlanManagement() revel.Result {
	return c.Render()
}

//ShowFarmer
func (c Farmer) ShowFarmer() revel.Result {
	return c.Render()
}
//FarmerManagement
func (c Farmer) FarmerManagement() revel.Result {
	return c.Render()
}

func (c Howto) HowtoRegister() revel.Result {
	return c.Render()
}
func (c Howto) HowtoCrops() revel.Result {
	return c.Render()
}
func (c Howto) HowtoSetting() revel.Result {
	return c.Render()
}
func (c Howto) HowtoMarket() revel.Result {
	return c.Render()
}

//Management
func (c Management) IndexManagement() revel.Result {
	return c.Render()
}


