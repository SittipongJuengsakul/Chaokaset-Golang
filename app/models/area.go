//Model : พื้นที่ ตำบล อำเภอ จังหวัด
//Author : Sittipong Jungsakul

package models
import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    //github.com/revel/revel"
    //"regexp"
    //"time"
    //"math/rand"
)

type Province struct {
  ProvinceId                            bson.ObjectId `bson:"_id,omitempty"`
  Province_Code,Province_Name           string
  Province_Name_Eng,Geo_Id              string
}

//จังหวัด
func GetProvinces(idprovince string) *Province {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  var province *Province
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("location_thailand").C("provinces")
  result := Province{}
	qmgo.Find(bson.M{"PROVINCE_ID": "1"}).One(&result) //คิวรี่จาก status เป็น 1 หรือ แปลงที่ไช้งานอยู่
  province = &Province{ProvinceId: result.ProvinceId,Province_Code: result.Province_Code}
  return province
}
