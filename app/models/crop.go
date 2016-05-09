//Model : การเพาะปลูก สำหรับจัดการการเพาะปลูก
//Author : Sittipong Jungsakul

package models
import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    //"github.com/revel/revel"
    //"regexp"
    "time"
    //"math/rand"
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
  crop = &Crop{CropId: result.CropId,Status: result.Status,CropName: result.CropName,PlantId: result.PlantId,SeedId: result.SeedId,PlanId: result.PlanId,Plant: result.Plant,Seed: result.Seed,StartDate: result.StartDate,EndDate: result.EndDate,Duration: result.Duration,Province: result.Province,Aumphur: result.Aumphur,Tumbon: result.Tumbon,Product: result.Product,Price: result.Price,Address : result.Address}
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
     err = qmgo.Insert(&Crop{Created_at: time.Now(),Updated_at: time.Now(),UserId : userid,Status: 1,CropName: crop.CropName,PlantId: crop.PlantId,SeedId: crop.SeedId,PlanId: crop.PlanId,Plant: plantnames.PlantName,Seed: seednames.SeedName,StartDate: crop.StartDate,EndDate: crop.EndDate,Duration: crop.Duration,Province: crop.Province,Aumphur: crop.Aumphur,Tumbon: crop.Tumbon,Product: crop.Product,Price: crop.Price,Address : crop.Address})
     if err != nil {
       return false
     }else{
       return true
     }
}
