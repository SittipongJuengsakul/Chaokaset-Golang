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
func GetAllCrops(skip int) (results []Crop,error bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("crops")
	qmgo.Find(bson.M{"status": 1}).Skip(skip).Sort("-updated_at").All(&results) //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  return results,true
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
     err = qmgo.Insert(&Crop{UserId : userid,Status: 1,CropName: crop.CropName,PlantId: crop.PlantId,SeedId: crop.SeedId,PlanId: crop.PlanId,Plant: plantnames.PlantName,Seed: seednames.SeedName,StartDate: crop.StartDate,EndDate: crop.EndDate,Duration: crop.Duration,Province: crop.Province,Aumphur: crop.Aumphur,Tumbon: crop.Tumbon,Product: crop.Product,Price: crop.Price,Address : crop.Address})
     if err != nil {
       return false
     }else{
       return true
     }
}
