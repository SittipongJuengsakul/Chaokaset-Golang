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
  Created_at,Updated_at                     time.Time
  PlanName,Plant,Seed,OldPlanId             string
  Description,Owner,OwnerCompany            string
  Duration,Status,TypePlan                  int //TypePlan คือประเภทของแปลง 0 คือไหม่ 1 คือต่อเนื่องจากอันเดิม
  product,price                             float64
  ConfirmNum,LikeNum,ViewNum,UsedNum        int
  //Like                                     *LogLikePlan
}

type LogLikePlan struct{
  UserId,CompanyId                          string
  Status                                    int
  Created_at,Updated_at                     time.Time
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
  plans = &Plan{PlanId: result.PlanId}
  return plans
}

//SavePlan (POST) บันทึกแผนการเพาะปลูก
func SavePlan(plan *Plan) (result bool) {
     session, err := mgo.Dial(ip_mgo)
     if err != nil {
         panic(err)
     }
     defer session.Close()
     session.SetMode(mgo.Monotonic, true)
     qmgo := session.DB("chaokaset").C("cropplans")
     if plan.TypePlan == 1{
       err = qmgo.Insert(&Plan{Created_at: time.Now(),Updated_at: time.Now(),PlanName: plan.PlanName,OwnerCompany: plan.OwnerCompany,Owner: plan.Owner,Plant: plan.Plant,Seed: plan.Seed,Duration: plan.Duration,Description: plan.Description,TypePlan: plan.TypePlan,OldPlanId : plan.OldPlanId,Status : 1,ViewNum: 1})
     }else{
       err = qmgo.Insert(&Plan{Created_at: time.Now(),Updated_at: time.Now(),PlanName: plan.PlanName,OwnerCompany: plan.OwnerCompany,Owner: plan.Owner,Plant: plan.Plant,Seed: plan.Seed,Duration: plan.Duration,Description: plan.Description,TypePlan: plan.TypePlan,OldPlanId: "",Status : 1,ViewNum: 1})
     }

     if err != nil {
       return false
     }else{
       return true
     }

}

//SavePlanLikeLog (POST) เก็บ log ของการกด like
func SavePlanLikeLog(UserId string,CompanyId string) (result bool){
    session, err := mgo.Dial(ip_mgo)
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    qmgo := session.DB("chaokaset").C("logcropplans")
    err = qmgo.Insert(&LogLikePlan{UserId: UserId,CompanyId: CompanyId,Status: 1,Created_at: time.Now(),Updated_at: time.Now()})
    if err != nil {
      return false
    }else{
      return true
    }
}
