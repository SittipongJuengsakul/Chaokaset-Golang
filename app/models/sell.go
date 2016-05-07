package models
import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
   // "golang.org/x/crypto/bcrypt"
  //  "github.com/revel/revel"
  //  "regexp"
  //  "time"
  //  "math/rand"
    //"fmt"
    "time"
   "github.com/alouche/go-geolib"
    "math"
)


type Sell struct{
  Sellid          bson.ObjectId `bson:"_id,omitempty"`
  Name            string
  Category        string
  Pic             string
  PicUp           []byte
  Price           int
  Distance        float64
  Address         Address
  Unit            string
  Detail          string
  Expire          string
  TimeCreate      time.Time
  OwnerId         bson.ObjectId
}
type Address struct{
  Lat             float64
  Long            float64
}
type  Owner struct{
  Name,Lastname,Prefix,Tel       string
}

type SellDetail struct{
  Sellid          bson.ObjectId `bson:"_id,omitempty"`
  Name            string
  Category        string
  Pic             string
  Price           int
  Address         Address
  Unit            string
  Detail          string
  Expire          string
  TimeCreate      time.Time
  OwnerId         bson.ObjectId
  Owner           Owner
}


//var Selldb = make(map[string]*Sell)

func (sell *Sell) SetDistance(data float64) {
  sell.Distance = data
}

func (SellDetail *SellDetail) SetOwnerName(data string) {
  SellDetail.Owner.Name = data
}

func (SellDetail *SellDetail) SetOwnerLastname(data string) {
  SellDetail.Owner.Lastname = data
}

func (SellDetail *SellDetail) SetOwnerPrefix(data string) {
  SellDetail.Owner.Prefix = data
}

func (SellDetail *SellDetail) SetOwnerTel(data string) {
  SellDetail.Owner.Tel = data
}
//GetSellData
func GetSellData(Lat float64, Long float64) []Sell {
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()
  //var data *Sell
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")
  /*result := Sell{}
	qmgo.Find(nil).All(&result)
  data = &Sell{Sellid: result.Sellid,Name: result.Name,Category: result.Category,Price: result.Price}
  return data*/
  var result []Sell
  
  qmgo.Find(nil).Sort("price").All(&result)
  //&Sell{Sellid: result.Sellid,Name: result.Name,Category: result.Category,Price: result.Price}
   for i := range result {
      lat1 := Lat
      lat2 := 13.286727
      lon1 := Long
      lon2 := 100.925619
      theta := lon1 - lon2
      dist := math.Sin(geolib.Deg2Rad(lat1)) * math.Sin(geolib.Deg2Rad(lat2)) + math.Cos(geolib.Deg2Rad(lat1)) * math.Cos(geolib.Deg2Rad(lat2)) * math.Cos(geolib.Deg2Rad(theta))
      dist = math.Acos(dist)
      dist = geolib.Rad2Deg(dist)
      result[i].SetDistance(dist * 60 * 1.1515 * 1.609344)
  }
  
  return result
}

func AddSellData(name string,category string, price int, unit string, detail string, expire string, ownerId bson.ObjectId) (result bool) {
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")
  err = qmgo.Insert(&Sell{Name: name, Category: category, Price: price,TimeCreate: time.Now(), Detail: detail, Expire: expire, Unit: unit, OwnerId: ownerId,Pic: "public/img/pic/rice1.jpg"})
  if err != nil {
    return false
  }else{
    return true
  }

}
func AddSellData2(name string,category string, price int, unit string, detail string, expire string, ownerId string) (result bool) {
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")
  err = qmgo.Insert(&Sell{Name: name, Category: category, Price: price,TimeCreate: time.Now(), Detail: detail, Expire: expire, Unit: unit, OwnerId: bson.ObjectIdHex(ownerId),Pic: "public/img/pic/rice1.jpg"})
  if err != nil {
    return false
  }else{
    return true
  }

}

func GetSellDetail(Idsell string) *SellDetail {
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()
 // var user *SellDetail
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")
  //result := SellDetail{}
  var result *SellDetail
  qmgo.Find(bson.M{"_id": bson.ObjectIdHex(Idsell)}).One(&result)
  //user = &SellDetail{Sellid:result.Sellid, Name:result.Name, Category:result.Category, Pic:result.Pic, Price:result.Price, Address.result.Address }
  data := GetOwnerData(result.OwnerId.Hex())
  result.SetOwnerName(data.Name)
  result.SetOwnerLastname(data.Lastname)
  result.SetOwnerPrefix(data.Prefix)
  result.SetOwnerTel(data.Tel)
  return result
}


func GetSearchSell(Name string,Lat float64,Long float64) []Sell {
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")
  var result []Sell
  
  qmgo.Find(bson.M{"name": bson.RegEx{".*"+Name, "s"}}).All(&result)
   for i := range result {
      lat1 := Lat
      lat2 := 13.286727
      lon1 := Long
      lon2 := 100.925619
      theta := lon1 - lon2
      dist := math.Sin(geolib.Deg2Rad(lat1)) * math.Sin(geolib.Deg2Rad(lat2)) + math.Cos(geolib.Deg2Rad(lat1)) * math.Cos(geolib.Deg2Rad(lat2)) * math.Cos(geolib.Deg2Rad(theta))
      dist = math.Acos(dist)
      dist = geolib.Rad2Deg(dist)
      result[i].SetDistance(dist * 60 * 1.1515 * 1.609344)  
  }
  return result
}

func GetOwnerData(id string) *Owner{
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()
  //id = id.Hex()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("users")
  var result *Owner
  qmgo.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
  return result
}











