//Model : การเพาะปลูก สำหรับจัดการการเพาะปลูก
//Author : Sittipong Jungsakul

package models
import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
    "github.com/revel/revel"
)

type Crop struct {
	CropId                                    bson.ObjectId `bson:"_id,omitempty"`
  Created_at,Updated_at                     time.Time
  PlanId,Plant,Seed,CropName                string
  PlantId,SeedId                            string
  StartDate,EndDate                         string
  Description                               string
  Address,Tumbon,Aumphur,Province           string
  Duration,Status                           int
  Product,Price                             float64
  UserId                                    string
  Rai,Ngarn,Wah                             float64
}

type Account struct {
	AccountId                                 bson.ObjectId `bson:"_id,omitempty"`
  Detail                                    string
  Status,TypeAccount                        int
  Price                                     float64
  Created_at,Updated_at                     time.Time
  CropId                                    string
}
type Problem struct {
	ProblemId                                 bson.ObjectId `bson:"_id,omitempty"`
  Detail,Problem                            string
  Status,StatusTopic                        int
  Created_at,Updated_at                     time.Time
  CropId                                    string
}
//ส่วน Validation
func (crop *Crop) ValidateAddCrop(v *revel.Validation) {
  v.Required(crop.CropName).Message("จำเป็นต้องกรอก ชื่อแปลงเพาะปลูก")
  v.Required(crop.Rai).Message("จำเป็นต้องกรอก จำนวนไร่")
  v.Required(crop.Ngarn).Message("จำเป็นต้องกรอก จำนวนงาน")
  v.Required(crop.Wah).Message("จำเป็นต้องกรอก จำนวนตารางวา")
  v.Required(crop.Price).Message("จำเป็นต้องกรอก จำนวนผลผลิตที่คาดหวัง")
  v.Required(crop.Product).Message("จำเป็นต้องกรอก ราคาผลผลิตที่คาดหวัง")
  /*v.Required(user.Name).Message("จำเป็นต้องกรอก ชื่อ")
  v.Required(user.Lastname).Message("จำเป็นต้องกรอก นามสกุล")
  v.Required(user.Tel).Message("จำเป็นต้องกรอก เบอร์โทรศัพท์")
  v.Match(user.Tel, regexp.MustCompile("^\\d*$")).Message("เบอร์โทรศัพท์เป็นตัวเลขเท่านั้น เช่น 08011122233")
  v.MinSize(user.Tel, 9).Message("เบอร์โทรศัพท์ต้องมี 10 ตัวเลข")
  v.MaxSize(user.Tel, 10).Message("เบอร์โทรศัพท์ต้องมี 10 ตัวเลข")
  v.Required(user.Email).Message("จำเป็นต้องกรอก อีเมล์")
  v.Email(user.Email).Message("กรอก Email ในลักษณะ sample@gmail.com")*/
}
//GetAllCrops (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลแผนการเพาะปลูกทั้งหมด
func GetAllCrops(skip int,userid string) (results []Crop,error bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("crops")
	qmgo.Find(bson.M{"status": 1,"userid": userid}).Skip(skip).All(&results) //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  return results,true
}

//GetPlans (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลแผนการเพาะปลูก
func GetOneCrops(cropid string) *Crop {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  var crop *Crop
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("crops")
  result := Crop{}
	qmgo.Find(bson.M{"status": 1,"_id": bson.ObjectIdHex(cropid)}).One(&result)
  crop = &Crop{Rai: result.Rai,Ngarn: result.Ngarn,Wah: result.Wah,CropId: result.CropId,Status: result.Status,CropName: result.CropName,PlantId: result.PlantId,SeedId: result.SeedId,PlanId: result.PlanId,Plant: result.Plant,Seed: result.Seed,StartDate: result.StartDate,EndDate: result.EndDate,Duration: result.Duration,Province: result.Province,Aumphur: result.Aumphur,Tumbon: result.Tumbon,Product: result.Product,Price: result.Price,Address : result.Address}
  return crop
}

//SaveCrop (POST) บันทึกการเพาะปลูก
func SaveCrop(crop *Crop,userid string) (result bool) {
     session, err := mgo.Dial(ip_mgo)
     if err != nil {
         panic(err)
     }
     defer session.Close()
     session.SetMode(mgo.Monotonic, true)
     qmgo := session.DB("chaokaset").C("crops")
     plantnames := GetPlantId(crop.PlantId)
     seednames := GetSeedId(crop.SeedId)
     err = qmgo.Insert(&Crop{Rai: crop.Rai,Ngarn: crop.Ngarn,Wah: crop.Wah,Created_at: time.Now(),Updated_at: time.Now(),UserId : userid,Status: 1,CropName: crop.CropName,PlantId: crop.PlantId,SeedId: crop.SeedId,PlanId: crop.PlanId,Plant: plantnames.PlantName,Seed: seednames.SeedName,StartDate: crop.StartDate,EndDate: crop.EndDate,Duration: crop.Duration,Province: crop.Province,Aumphur: crop.Aumphur,Tumbon: crop.Tumbon,Product: crop.Product,Price: crop.Price,Address : crop.Address})
     if err != nil {
       return false
     }else{
       return true
     }
}
//SaveCrop (POST) บันทึกการเพาะปลูก
func DisableOneCrops(idcrop string) (result bool) {
     session, err := mgo.Dial(ip_mgo)
     if err != nil {
         panic(err)
     }
     defer session.Close()
     session.SetMode(mgo.Monotonic, true)
     //cropqry := bson.ObjectIdHex(cropid)
     qmgo := session.DB("chaokaset").C("crops")
    colQuerier := bson.M{"_id": bson.ObjectIdHex(idcrop)}
    change := bson.M{"$set": bson.M{"status": 0, "Updated_at": time.Now()}}
    err = qmgo.Update(colQuerier, change)
     if err != nil {
       return false
     }else{
       return true
     }
}
//UpdateCrop (POST) บันทึกการเพาะปลูก
func UpdateCrop(crop *Crop,cropid string) (result bool) {
     session, err := mgo.Dial(ip_mgo)
     if err != nil {
         panic(err)
     }
     defer session.Close()
     session.SetMode(mgo.Monotonic, true)
     //cropqry := bson.ObjectIdHex(cropid)
     qmgo := session.DB("chaokaset").C("crops")
    colQuerier := bson.M{"_id": bson.ObjectIdHex(cropid)}
    change := bson.M{"$set": bson.M{"Updated_at": time.Now(),"cropname": crop.CropName,"price": crop.Price,"product": crop.Product,"rai": crop.Rai,"ngarn": crop.Ngarn,"wah": crop.Wah,"duration": crop.Duration,"startdate": crop.StartDate,"enddate": crop.EndDate}}
    err = qmgo.Update(colQuerier, change)
     if err != nil {
       return false
     }else{
       return true
     }
}
//GetAllAccounts (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลแผนการเพาะปลูกทั้งหมด
func GetAllAccounts(idcrop string,skip int) (results []Account,error bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("accounts")
  skip = skip*10;
	qmgo.Find(bson.M{"status": 1,"cropid": idcrop}).Sort("-updated_at").Skip(skip).Limit(10).All(&results) //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  return results,true
}
//GetAllAccounts (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลแผนการเพาะปลูกทั้งหมด
func GetFullAllAccounts(idcrop string) (results []Account,error bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("accounts")
	qmgo.Find(bson.M{"status": 1,"cropid": idcrop}).All(&results) //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  return results,true
}
//GetCountAccount (GET)
func GetCountProblem(idcrop string) (result int,err error) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("problems")
	result,err = qmgo.Find(bson.M{"status": 1,"cropid": idcrop}).Count() //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  return result,err
}
//GetSearchAllAccounts (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลแผนการเพาะปลูกทั้งหมด
func GetSearchAllAccounts(idcrop string,word string) (results []Account,error bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("accounts")
	qmgo.Find(bson.M{"status": 1,"cropid": idcrop}).Sort("-updated_at").All(&results) //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  return results,true
}
//GetOneAccounts (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลแผนการเพาะปลูกทั้งหมด
func GetOneAccount(idcrop string,idaccount string) (results *Account,error bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("accounts")
  qmgo.Find(bson.M{"status": 1,"_id": bson.ObjectIdHex(idaccount)}).One(&results)
  account := &Account{CropId: results.CropId,Status: results.Status,AccountId: results.AccountId,Updated_at: results.Updated_at,Detail: results.Detail,Price: results.Price,TypeAccount: results.TypeAccount}
  return account,true
}
//SaveAccount (POST)
func SaveAccount(idcrop string,typeaccount int,detail string,price float64) (results []Account) {
     session, err := mgo.Dial(ip_mgo)
     if err != nil {
         panic(err)
     }
     defer session.Close()
     session.SetMode(mgo.Monotonic, true)
     qmgo := session.DB("chaokaset").C("accounts")
     err = qmgo.Insert(&Account{CropId: idcrop,Created_at: time.Now(),Updated_at: time.Now(),Status: 1,TypeAccount: typeaccount,Price: price,Detail: detail})
     if err != nil {
       panic(err)
     }

     accountdatas,err_getAccount := GetAllAccounts(idcrop,0)
     if err_getAccount {
       return accountdatas
     }
     return accountdatas

}
//UpdateAccount (PUT)
func UpdateAccount(idaccount string,detail string,price float64) (result bool) {
     session, err := mgo.Dial(ip_mgo)
     if err != nil {
         panic(err)
     }
     defer session.Close()
     session.SetMode(mgo.Monotonic, true)
     qmgo := session.DB("chaokaset").C("accounts")
     colQuerier := bson.M{"_id": bson.ObjectIdHex(idaccount)}
     change := bson.M{"$set": bson.M{"detail": detail,"price": price, "Updated_at": time.Now()}}
     err = qmgo.Update(colQuerier, change)
     if err != nil {
       return false
     }else{
       return true
     }
}
//SaveCrop (POST) บันทึกการเพาะปลูก
func DisableOneAccount(idaccount string) (result bool) {
     session, err := mgo.Dial(ip_mgo)
     if err != nil {
         panic(err)
     }
     defer session.Close()
     session.SetMode(mgo.Monotonic, true)
     //cropqry := bson.ObjectIdHex(cropid)
     qmgo := session.DB("chaokaset").C("accounts")
    colQuerier := bson.M{"_id": bson.ObjectIdHex(idaccount)}
    change := bson.M{"$set": bson.M{"status": 0, "Updated_at": time.Now()}}
    err = qmgo.Update(colQuerier, change)
     if err != nil {
       return false
     }else{
       return true
     }
}


//GetAllProblems (GET)
func GetAllProblems(idcrop string,skip int) (results []Problem,error bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("problems")
  skip = skip*10;
	qmgo.Find(bson.M{"status": 1,"cropid": idcrop}).Sort("-updated_at").Skip(skip).Limit(10).All(&results) //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  return results,true
}

//GetOneProblems (GET)
func GetOneProblem(idcrop string,idproblem string) (results *Problem,error bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("problems")
  qmgo.Find(bson.M{"status": 1,"_id": bson.ObjectIdHex(idproblem)}).One(&results)
  account := &Problem{StatusTopic: results.StatusTopic,Problem: results.Problem,CropId: results.CropId,Status: results.Status,ProblemId: results.ProblemId,Updated_at: results.Updated_at,Detail: results.Detail}
  return account,true
}
//SaveProblem (POST)
func SaveProblem(idcrop string,problem string,detail string) (results []Problem) {
     session, err := mgo.Dial(ip_mgo)
     if err != nil {
         panic(err)
     }
     defer session.Close()
     session.SetMode(mgo.Monotonic, true)
     qmgo := session.DB("chaokaset").C("problems")
     err = qmgo.Insert(&Problem{StatusTopic: 0,CropId: idcrop,Created_at: time.Now(),Updated_at: time.Now(),Status: 1,Problem: problem,Detail: detail})
     if err != nil {
       panic(err)
     }

     accountdatas,err_getProblem := GetAllProblems(idcrop,0)
     if err_getProblem {
       return accountdatas
     }
     return accountdatas

}
//UpdateProblem (PUT)
func UpdateProblem(idproblem string,detail string) (result bool) {
     session, err := mgo.Dial(ip_mgo)
     if err != nil {
         panic(err)
     }
     defer session.Close()
     session.SetMode(mgo.Monotonic, true)
     qmgo := session.DB("chaokaset").C("problems")
     colQuerier := bson.M{"_id": bson.ObjectIdHex(idproblem)}
     change := bson.M{"$set": bson.M{"statustopic": 0,"detail": detail, "Updated_at": time.Now()}}
     err = qmgo.Update(colQuerier, change)
     if err != nil {
       return false
     }else{
       return true
     }
}
//SaveCrop (POST)
func DisableOneProblem(idproblem string) (result bool) {
     session, err := mgo.Dial(ip_mgo)
     if err != nil {
         panic(err)
     }
     defer session.Close()
     session.SetMode(mgo.Monotonic, true)
     qmgo := session.DB("chaokaset").C("problems")
    colQuerier := bson.M{"_id": bson.ObjectIdHex(idproblem)}
    change := bson.M{"$set": bson.M{"status": 0, "Updated_at": time.Now()}}
    err = qmgo.Update(colQuerier, change)
     if err != nil {
       return false
     }else{
       return true
     }
}
