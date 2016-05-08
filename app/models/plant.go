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
	PlantId                                   bson.ObjectId `bson:"_id,omitempty"`
  Created_at,Updated_at                     time.Time
  PlantName                                 string
}

type Seed struct { //สร้าง Struct ของ Seed
	SeedId                                    bson.ObjectId `bson:"_id,omitempty"`
  Created_at,Updated_at                     time.Time
  PlantName,PlantId,SeedName                string
  OwnerCompany                              *OwnerCompany
}

type OwnerCompany struct {
  CompanyId                                 bson.ObjectId `bson:"_id"`
  CompanyName,OwnerName                     string
  Created_at,Updated_at                     time.Time
}

//GetAllPlants (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลพืชทั้งหมด
func GetAllPlants(skip int) (results []Plant,err_query error) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("plants")
	err_query = qmgo.Find(bson.M{}).Limit(10).Skip(skip).Sort("-updated_at").All(&results) //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  return results,err_query
}
//GetPlant (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลพืชทั้งหมด
func GetPlant(word string) *Plant {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  var plant *Plant
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("plants")
  result := Plant{}
	qmgo.Find(bson.M{"plantname": word}).One(&result)
  plant = &Plant{PlantName: result.PlantName,PlantId: result.PlantId,Created_at: result.Created_at,Updated_at: result.Updated_at}
  return plant
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

//GetAllSeeds (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลพืชทั้งหมด
func GetAllSeeds(skips int) (results []Plant) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("seeds")
	qmgo.Find(bson.M{}).Limit(10).Skip(skips).All(&results) //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  return results
}
//GetSeed (GET) ฟังก์ชั่นสำหรับเรียกข้อมูลพืชทั้งหมด
func GetSeed(skips int,plantname string,seedname string) *Seed {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  var seed *Seed
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("seeds")
  result := Seed{}
	qmgo.Find(bson.M{"plantname": plantname}).One(&result)
  seed = &Seed{PlantName: result.PlantName,PlantId: result.PlantId,Created_at: result.Created_at,Updated_at: result.Updated_at}
  return seed
}

//SaveSeed (POST)
func SaveSeed(seedname string,plantname string) (result bool){
    session, err := mgo.Dial(ip_mgo)
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    qmgo := session.DB("chaokaset").C("seeds")
    var ownerCompany *OwnerCompany
    err = qmgo.Insert(&Seed{SeedName: seedname,PlantName: plantname,OwnerCompany: ownerCompany,Created_at: time.Now(),Updated_at: time.Now()})
    if err != nil {
      return false
    }else{
      return true
    }
}
