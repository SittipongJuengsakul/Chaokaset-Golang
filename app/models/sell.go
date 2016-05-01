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
)


type Sell struct{
  Sellid          bson.ObjectId `bson:"_id,omitempty"`
  Name            string
  Category        string
  Pic             string
  Price           int
  Distance        float64
  Address         Address
  Unit            string
  Detail          string
  Expire          string
  TimeCreate      time.Time
}
type Address struct{
  Lat       float64
  Long      float64
}

//var Selldb = make(map[string]*Sell)

func (sell *Sell) SetDistance(data float64) {
  sell.Distance = data
}

//GetSellData
func GetSellData() []Sell {
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
    
    result[i].SetDistance(result[i].Address.Lat)
  }
  
  return result
}

func AddSellData(name string,category string, price int, unit string, detail string, expire string) (result bool) {
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")
  err = qmgo.Insert(&Sell{Name: name, Category: category, Price: price,TimeCreate: time.Now(), Detail: detail, Expire: expire, Unit: unit})
  if err != nil {
    return false
  }else{
    return true
  }

}





