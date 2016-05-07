//Model : สำหรับข้อมูลพืช และพันธุ์พืช
//Author : Sittipong Jungsakul

package models
import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    //"github.com/revel/revel"
    //"regexp"
    "time"
)

type Plant struct { //สร้าง Struct ของ Plant
	Plant                                     bson.ObjectId `bson:"_id,omitempty"`
  Created_at,Updated_at                     time.Time
  PlantName                                 string
}

//GetAllPlants (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลพืชทั้งหมด
func GetAllPlants(skip int) (results []Plan,error bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  //var plans *Plan
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("plants")
	qmgo.Find(bson.M{"status": 1}).Limit(10).Skip(skip).Sort("-updated_at").All(&results) //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  //plans = &Plan{PlanId: result.PlanId,Created_at: result.Created_at,Updated_at: result.Updated_at,PlanName: result.PlanName,Plant: result.Plant,Seed: result.Seed,OldPlanId: result.OldPlanId,Description: result.Description,Owner: result.Owner,OwnerCompany: result.OwnerCompany,Duration: result.Duration,TypePlan : result.TypePlan,ConfirmNum: result.ConfirmNum,LikeNum: result.LikeNum,ViewNum: result.ViewNum,UsedNum: result.UsedNum}
  return results,true
}

//SavePlanLikeLog (POST) เก็บ log ของการกด like
func SavePlant(PlantName string) (result bool){
    session, err := mgo.Dial(ip_mgo)
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    qmgo := session.DB("chaokaset").C("plants")
    err = qmgo.Insert(&Plant{PlantName: PlantName,Created_at: time.Now(),Updated_at: time.Now()})
    if err != nil {
      return false
    }else{
      return true
    }
}
