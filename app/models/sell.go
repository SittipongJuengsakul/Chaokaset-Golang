package models

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
    "github.com/alouche/go-geolib"
    "math"
    "fmt"
    "strings"
    "strconv"
)

type Sells struct{
  Sellid          bson.ObjectId `bson:"_id,omitempty"`
  Name            string
  Category        string
  Pic             string
  Price           int
  Distance        float64
  Address         *Address
  Unit            string
  Detail          string
  Start           string
  Expire          string
  TimeCreate      time.Time
  OwnerId         bson.ObjectId
  Status          int
  SellType        int
  Like 			      []bson.ObjectId
  NumberOfLike	  int
  Comment         []Comments
}

type SellDetail struct{
  Sellid          bson.ObjectId `bson:"_id,omitempty"`
  Name            string
  Category        string
  Pic             string
  Price           int
  Address         *Address
  Unit            string
  Detail          string
  Expire          string
  TimeCreate      time.Time
  OwnerId         bson.ObjectId
  Owner           Owner
  Status          int
  SellType        int
  StatusLike      int
  Like        []bson.ObjectId
  NumberOfLike    int
  Comment         []Comments     
}
type Comments struct{
  Userid          bson.ObjectId
  Name            string
  Data            string
  TimeCreate      time.Time
}

type  Owner struct{
  Name,Lastname,Prefix,Tel       string
}

type Address struct{
  Lat             float64
  Long            float64
}
type UserId struct{
  Userid          bson.ObjectId `bson:"_id,omitempty"`
}

type ReturnSellId struct {
    Status      bool
    SellId      string
}

type PostSell struct{
  Name            string
  Category        string
  Price           int
  Unit            string
  Detail          string
  Expire          string
}



func (sell *Sells) SetDistance(data float64) {
  sell.Distance = data
}

func (sell *Sells) SetNumLike(data int) {
  sell.NumberOfLike = data
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

func (SellDetail *SellDetail) SetStatusLike(data int) {
  SellDetail.StatusLike = data
}

func (SellDetail *SellDetail) SetNumLikeDetail(data int) {
  SellDetail.NumberOfLike = data
}
//var ip_mgo = "188.166.230.170"

func GetSellData(Lat float64, Long float64) []Sells {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")

  var result,result_data []Sells
  
  qmgo.Find(bson.M{"status": 1}).Sort("-timecreate").All(&result)

  t := time.Now()
  tNow := t.Format("2006-01-02")
  var s2 string
   for i := range result {
      lat1 := Lat
      lat2 := result[i].Address.Lat
      lon1 := Long
      lon2 := result[i].Address.Long
      theta := lon1 - lon2
      dist := math.Sin(geolib.Deg2Rad(lat1)) * math.Sin(geolib.Deg2Rad(lat2)) + math.Cos(geolib.Deg2Rad(lat1)) * math.Cos(geolib.Deg2Rad(lat2)) * math.Cos(geolib.Deg2Rad(theta))
      dist = math.Acos(dist)
      dist = geolib.Rad2Deg(dist)
      result[i].SetDistance(dist * 60 * 1.1515 * 1.609344)
      result[i].SetNumLike(len(result[i].Like))

      s := strings.Split(result[i].Expire," ")
     // fmt.Println(s)
      switch s[1] {
        case "มกราคม": 
          s2 = "01"
        case "กุมภาพันธ์":
          s2 = "02"
        case "มีนาคม": 
          s2 = "03"
        case "เมษายน": 
          s2 = "04"
        case "พฤษภาคม": 
          s2 = "05"
        case "มิถุนายน": 
          s2 = "06"
        case "กรกฎาคม": 
          s2 = "07"
        case "สิงหาคม": 
          s2 = "08"
        case "กันยายน": 
          s2 = "09"
        case "ตุลาคม": 
          s2 = "10"
        case "พฤษจิกายน": 
          s2 = "11"
        case "ธันวาคม": 
          s2 = "12"
      }
      //fmt.Println(s2)
      year,_ := strconv.ParseInt(s[2], 10, 64)
    //year = year+543
    yearEng:= strconv.FormatInt(year-543, 10)
      dateNum := yearEng+"-"+s2+"-"+s[0]
      fmt.Println(dateNum)
      if dateNum >= tNow{
        result_data = append(result_data,result[i])
      }
  }
    
  return result_data
}

func GetSellDataByCategory(category string,Lat float64, Long float64) []Sells {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")

  var result,result_data []Sells
  
  qmgo.Find(bson.M{"status": 1,"category": category}).Sort("-timecreate").All(&result)
 // fmt.Printf("%+v\n", result)
  t := time.Now()
  tNow := t.Format("2006-01-02")
  var s2 string
   for i := range result {
      lat1 := Lat
      lat2 := result[i].Address.Lat
      lon1 := Long
      lon2 := result[i].Address.Long
      theta := lon1 - lon2
      dist := math.Sin(geolib.Deg2Rad(lat1)) * math.Sin(geolib.Deg2Rad(lat2)) + math.Cos(geolib.Deg2Rad(lat1)) * math.Cos(geolib.Deg2Rad(lat2)) * math.Cos(geolib.Deg2Rad(theta))
      dist = math.Acos(dist)
      dist = geolib.Rad2Deg(dist)
      //fmt.Printf("%+v\n", dist)
      result[i].SetDistance(dist * 60 * 1.1515 * 1.609344)
      result[i].SetNumLike(len(result[i].Like))

      s := strings.Split(result[i].Expire," ")
    //  fmt.Println(s)
      switch s[1] {
        case "มกราคม": 
          s2 = "01"
        case "กุมภาพันธ์":
          s2 = "02"
        case "มีนาคม": 
          s2 = "03"
        case "เมษายน": 
          s2 = "04"
        case "พฤษภาคม": 
          s2 = "05"
        case "มิถุนายน": 
          s2 = "06"
        case "กรกฎาคม": 
          s2 = "07"
        case "สิงหาคม": 
          s2 = "08"
        case "กันยายน": 
          s2 = "09"
        case "ตุลาคม": 
          s2 = "10"
        case "พฤษจิกายน": 
          s2 = "11"
        case "ธันวาคม": 
          s2 = "12"
      }
     //fmt.Println(s2)
      year,_ := strconv.ParseInt(s[2], 10, 64)
    //year = year+543
    yearEng:= strconv.FormatInt(year-543, 10)
      dateNum := yearEng+"-"+s2+"-"+s[0]
      fmt.Println(dateNum)
      if dateNum >= tNow{
        result_data = append(result_data,result[i])
      }
  }

  return result_data
}

func AddSellData(name string,category string, price int, unit string, detail string, expire string, ownerId bson.ObjectId) (result bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")
  
 // var  A *Address
 // A = &Address{Lat: 13.00,Long: 100.00}
  i := bson.NewObjectId()
  
  err = qmgo.Insert(&Sells{
    Sellid: i,
    Name: name, 
    Category: category, 
    Price: price,
    TimeCreate: time.Now(), 
    Detail: detail, 
    Expire: expire, 
    Unit: unit, 
    OwnerId: ownerId,
    Status: 1,
    SellType: 2,
   // Like: [],
  })

  if err != nil {
    return false
  }else{
    return true
  }

}

func AddSellData2(name string,category string, price int, unit string, detail string, expire string, ownerId string, lat float64, long float64,sellType int) *ReturnSellId {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  
  var  A *Address
  A = &Address{Lat: lat,Long: long}

  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")

  t := time.Now()
  //fmt.Println(t.Year())
 // fmt.Println(t.Month())
 // fmt.Println(t.Day())
  Month := t.Format("01")
  fmt.Println(Month)
  var MonthName string

  switch Month {
      case "01": 
        MonthName = "มกราคม"
      case "02": 
        MonthName = "กุมภาพันธ์"
      case "03": 
        MonthName = "มีนาคม"
      case "04": 
        MonthName = "เมษายน"
      case "05": 
        MonthName = "พฤษภาคม"
      case "06": 
        MonthName = "มิถุนายน"
      case "07": 
        MonthName = "กรกฎาคม"
      case "08": 
        MonthName = "สิงหาคม"
      case "09": 
        MonthName = "กันยายน"
      case "10": 
        MonthName = "ตุลาคม"
      case "11": 
        MonthName = "พฤษจิกายน"
      case "12": 
        MonthName = "ธันวาคม"       
  }
  //year,_ := strconv.ParseInt(t.Year(), 10, 64)
  year := t.Year()+543
  yearThai:= strconv.Itoa(year)

  start := t.Format("02") +" "+MonthName+" "+yearThai
  fmt.Println(start)
  i := bson.NewObjectId()
  
  err = qmgo.Insert(&Sells{
    Sellid: i,
    Name: name, 
    Category: category, 
    Price: price,
    TimeCreate: t, 
    Detail: detail, 
    Expire: expire, 
    Unit: unit, 
    OwnerId: bson.ObjectIdHex(ownerId), 
    //Pic: "public/img/pic/rice1.jpg",
    Address: A ,
    Status: 1,
    SellType: sellType,
    Start: start,
   // Like: [],
  })

  if err != nil {
    return &ReturnSellId{Status: false,SellId: ""}
  }else{
    return &ReturnSellId{Status: true,SellId: i.Hex()}
  }

}

func GetSellDetail(Idsell string,Userid string) *SellDetail {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")
  //ประกาศตัวแปร
  var result *SellDetail
  //คิวลี่ข้อมูลการขายโดยกำหนดเลข id การขาย
  qmgo.Find(bson.M{"_id": bson.ObjectIdHex(Idsell),"status": 1}).One(&result)

  check := CheckLike(result.Sellid.Hex(),Userid)
  if check != true{
    //มีคนไลน์
    result.SetStatusLike(1)
  }else{
    //ไม่มี
    result.SetStatusLike(0)
  }
  //เพิ่มข้อมูลของเจ้าของสินค้า
  data := GetOwnerData(result.OwnerId.Hex())
  result.SetOwnerName(data.Name)
  result.SetOwnerLastname(data.Lastname)
  result.SetOwnerPrefix(data.Prefix)
  result.SetOwnerTel(data.Tel)
  result.SetNumLikeDetail(len(result.Like))

  //fmt.Printf("%+v\n", result.Comment)
  for i := range result.Comment {
    //fmt.Printf("%+v\n", result.Comment[i].Userid)
    Detail := GetOwnerData(result.Comment[i].Userid.Hex()) 
    fmt.Printf("%+v\n", result.Comment[i].Userid.Hex())
    fmt.Printf("%+v\n", Detail.Name)
    fmt.Printf("%+v\n", result.Comment[i].Name)
    Name := Detail.Prefix+Detail.Name+"  "+Detail.Lastname
    result.Comment[i].Name = Name
    fmt.Printf("%+v\n", result.Comment[i].Name)   
    //result[i].SetDetailUserid(Detail.Name,Detail.Lastname,Detail.Prefix,Detail.Tel)
   // func (SellDetail *SellDetail) SetDetailUserid(name string,lastname string,prefix string,tel string) 
    //result.Comment[i].UseridDetail.Name.SetDetailUserid(i,Detail.Name)
   // result.Comment[i].SetDetailUserid(Detail.Name)
  }
  
  return result
}


func GetSearchSell(Name string,Lat float64,Long float64) []Sells {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")
  var result,result_data []Sells
  
  qmgo.Find(bson.M{"name": bson.RegEx{".*"+Name, "s"},"status": 1}).All(&result)
 // fmt.Printf("%+v\n", result)
  t := time.Now()
  tNow := t.Format("2006-01-02")
  var s2 string
  
  for i := range result {
      lat1 := Lat
      lat2 := result[i].Address.Lat
      lon1 := Long
      lon2 := result[i].Address.Long
      theta := lon1 - lon2
      dist := math.Sin(geolib.Deg2Rad(lat1)) * math.Sin(geolib.Deg2Rad(lat2)) + math.Cos(geolib.Deg2Rad(lat1)) * math.Cos(geolib.Deg2Rad(lat2)) * math.Cos(geolib.Deg2Rad(theta))
      dist = math.Acos(dist)
      dist = geolib.Rad2Deg(dist)
      result[i].SetDistance(dist * 60 * 1.1515 * 1.609344)
      result[i].SetNumLike(len(result[i].Like))

      s := strings.Split(result[i].Expire," ")
      
      switch s[1] {
        case "มกราคม": 
          s2 = "01"
        case "กุมภาพันธ์":
          s2 = "02"
        case "มีนาคม": 
          s2 = "03"
        case "เมษายน": 
          s2 = "04"
        case "พฤษภาคม": 
          s2 = "05"
        case "มิถุนายน": 
          s2 = "06"
        case "กรกฎาคม": 
          s2 = "07"
        case "สิงหาคม": 
          s2 = "08"
        case "กันยายน": 
          s2 = "09"
        case "ตุลาคม": 
          s2 = "10"
        case "พฤษจิกายน": 
          s2 = "11"
        case "ธันวาคม": 
          s2 = "12"
      }
     // fmt.Println(s2)
      dateNum := s[2]+"-"+s2+"-"+s[0]
      //fmt.Println(dateNum)
      if dateNum >= tNow{
        result_data = append(result_data,result[i])
      }
  }
  return result_data
}

func GetOwnerData(id string) *Owner{
  session, err := mgo.Dial(ip_mgo)
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

func GetManageSell(id string) []Sells {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")
  var result []Sells
  
  qmgo.Find(bson.M{"ownerid": bson.ObjectIdHex(id)}).Sort("TimeCreate").All(&result)

  return result
}

func GetUserid(username string) *UserId {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("users")
  var result *UserId
  qmgo.Find(bson.M{"username": username}).One(&result)
  return result
}

func UpdateStatusSell(idSell string,status int) (result bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")

  colQuerier := bson.M{ "_id": bson.ObjectIdHex(idSell) }
  change := bson.M{"$set": bson.M{"status": status}}
  
  err = qmgo.Update(colQuerier, change)

  if err != nil {
    return false
  }else{
    return true
  }
}

func EditProductSell(idSell string, name string, category string, price int,detail string,expire string,unit string,lat float64,long float64) (result bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")

  var  A *Address
  A = &Address{Lat: lat,Long: long}
  
  colQuerier := bson.M{ "_id": bson.ObjectIdHex(idSell) }
  
  change := bson.M{"$set": bson.M{
    "name": name, 
    "category": category, 
    "price": price,
    "detail": detail, 
    "expire": expire, 
    "unit": unit, 
    "address": A,
  }}
  
  err = qmgo.Update(colQuerier, change)

  if err != nil {
    return false
  }else{
    return true
  }
}

func UpdatePic(idSell string,pic string) (result bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")

  colQuerier := bson.M{ "_id": bson.ObjectIdHex(idSell) }
  change := bson.M{"$set": bson.M{"pic": pic}}
  
  err = qmgo.Update(colQuerier, change)

  if err != nil {
    return false
  }else{
    return true
  }
}

func Like(idSell string,idUser string) (result bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")

  colQuerier := bson.M{ "_id": bson.ObjectIdHex(idSell) }

  change := bson.M{"$push": bson.M{"like": bson.ObjectIdHex(idUser)}}
  
  err = qmgo.Update(colQuerier, change)

  if err != nil {
    return false
  }else{
    return true
  }
}

func UnLike(idSell string,idUser string) (result bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  
  qmgo := session.DB("chaokaset").C("sell")

  colQuerier := bson.M{ "_id": bson.ObjectIdHex(idSell) }

  change := bson.M{"$pull": bson.M{"like": bson.ObjectIdHex(idUser)}}
  
  err = qmgo.Update(colQuerier, change)

  if err != nil {
    return false
  }else{
    return true
  }
}

func CheckLike(idSell string,idUser string) (result bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  
  qmgo := session.DB("chaokaset").C("sell")

  var data *Sells
  qmgo.Find(bson.M{"_id": bson.ObjectIdHex(idSell),"like": bson.ObjectIdHex(idUser)}).One(&data)

  if data != nil{
    return false
  }else{
    return true
  }
  //return true;
}

func GetComment(idSell string) []Sells {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")

  var result []Sells
  //var result []Comment
  
  qmgo.Find(bson.M{"_id": bson.ObjectIdHex(idSell)}).All(&result)
 
  return result
}


func Comment(idSell string,idUser string,data string) (result bool) {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("sell")

  colQuerier := bson.M{ "_id": bson.ObjectIdHex(idSell) }

  change := bson.M{"$push": bson.M{"comment": bson.M{"userid":bson.ObjectIdHex(idUser),"data": data,"timecreate":time.Now()}}}
  
  err = qmgo.Update(colQuerier, change)

  if err != nil {
    return false
  }else{
    return true
  }
}

func GetCropSell(Userid string) []Crop {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("crops")

  var result,result_data []Crop
  
  t := time.Now()
  tNow := t.Format("2006-01-02")
 
 // fmt.Printf("%q\n", strings.Split("27 พฤษภาคม 2558", " "))
  qmgo.Find(bson.M{"userid" : Userid}).Sort("-timecreate").All(&result)
  //qmgo.Find(bson.M{"userid" : Userid,"enddate": bson.M{"$lte":"26 พฤษภาคม 2016"}}).Sort("-timecreate").All(&result)
 fmt.Println(result)
  for i := range result {
    s := strings.Split(result[i].EndDate," ")
    var s2 string
    switch s[1] {
      case "มกราคม": 
        s2 = "01"
      case "กุมภาพันธ์":
        s2 = "02"
      case "มีนาคม": 
        s2 = "03"
      case "เมษายน": 
        s2 = "04"
      case "พฤษภาคม": 
        s2 = "05"
      case "มิถุนายน": 
        s2 = "06"
      case "กรฏาคม": 
        s2 = "07"
      case "สิงหาคม": 
        s2 = "08"
      case "กันยายน": 
        s2 = "09"
      case "ตุลาคม": 
        s2 = "10"
      case "พฤษจิกายน": 
        s2 = "11"
      case "ธันวาคม": 
        s2 = "12"
    }
    year,_ := strconv.ParseInt(s[2], 10, 64)
    //year = year+543
    yearEng:= strconv.FormatInt(year-543, 10)
    dateNum := yearEng+"-"+s2+"-"+s[0]
   // fmt.Println(s2)
  //  dateNum := s[2]+"-"+s2+"-"+s[0]
    fmt.Println(tNow)
    fmt.Println("----------------")
    fmt.Println("5555"+dateNum)
    if dateNum <= tNow{
      result_data = append(result_data,result[i])
    }
  }

  fmt.Println(result_data)


  return result_data
}

func GetCropSellDetail(Userid string,Cropid string) *Crop {
  session, err := mgo.Dial(ip_mgo)
  if err != nil {
      panic(err)
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("crops")

  var result *Crop
  
  qmgo.Find(bson.M{"status": 1,"userid" : Userid,"_id": bson.ObjectIdHex(Cropid)}).One(&result)

  return result
}