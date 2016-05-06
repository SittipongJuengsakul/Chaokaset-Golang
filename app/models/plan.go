//Model : แผนการเพาะปลูก สำหรับจัดการแผนการเพาะปลูก
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

type Plan struct { //สร้าง Struct ของ Plan
	PlanId                                    bson.ObjectId `bson:"_id,omitempty"`
  Timestamp                                 time.Time
  Code,NamePlan,Plant,Seed                  string
  Description,Owner,OwnerCompany            string
  Duration,Status                           int
  product,price                             float64
  ConfirmNum,LikeNum,ViewNum,UsedNum        int
}

//GetAllPlans (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลแผนการเพาะปลูกทั้งหมด
func GetAllPlans() *Plan {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  var plans *Plan
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("cropplans")
  result := Plan{}
	qmgo.Find(bson.M{"status": 1}).One(&result) //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  plans = &Plan{PlanId: result.PlanId,Code: result.Code}
  return plans
}
