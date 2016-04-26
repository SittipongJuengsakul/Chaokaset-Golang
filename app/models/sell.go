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
)


type Sell struct{
  Sellid          bson.ObjectId `bson:"_id,omitempty"`
  Name,Category   string
  Price           int

}

//var Selldb = make(map[string]*Sell)

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
  qmgo.Find(nil).Sort("name").All(&result)

  //&Sell{Sellid: result.Sellid,Name: result.Name,Category: result.Category,Price: result.Price}
  return result
}





