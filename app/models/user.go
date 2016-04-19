package models
import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "golang.org/x/crypto/bcrypt"
    "github.com/revel/revel"
    "regexp"
    "time"
    "math/rand"
)


type User struct { //สร้าง Struct
	Userid                                    bson.ObjectId `bson:"_id,omitempty"`
	Username,Password,Validpassword           string
  Role                                      int
  Name,Lastname,Prefix,Tel,Pic              string
  HashedPassword                            []byte
  Timestamp                                 time.Time
}
type UserByChaokaset struct{
  Userid                                    bson.ObjectId `bson:"_id,omitempty"`
  Username,Name,Lastname,Prefix,Tel,Pic     string
  Password                                  []byte
  Timestamp                                 time.Time
  Role                                      int
}

type UserData struct{
  Userid                                    bson.ObjectId `bson:"_id,omitempty"`
  Timestamp                                 time.Time
  Role                                      int
  Username,Name,Lastname,Prefix,Tel,Pic     string
  Email,Province,Aumphur,Tumbon,Address     string
  Zipcode                                   string
}
var userdb = make(map[string]*User)

//ส่วน Validation Form
var userRegex = regexp.MustCompile("^\\w*$")

//ฟังก์ชั่น GenString สำหรับ Generate String
func GenString(num int) (result string){
  rand.Seed(time.Now().UnixNano())
  var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
  b := make([]rune, num)
  for i := range b {
    b[i] = letterRunes[rand.Intn(len(letterRunes))]
  }
  result = string(b[:num])
  return result
}

func (user *User) Validate(v *revel.Validation) {
  v.Required(user.Username).Message("จำเป็นต้องกรอก ชื่อผู้ใช้งาน")
  v.Match(user.Username, regexp.MustCompile("[a-zA-Z0-9_]")).Message("ภาษาอังกฤษ และตัวเลขเท่านั้น")
	v.MinSize(user.Username, 4).Message("ชื่อผู้ใช้ต้องมากกว่า 4 ตัวอักษร")
  v.MaxSize(user.Username, 16).Message("ชื่อผู้ใช้ต้องน้อยกว่า 16 ตัวอักษร")
  v.Required(user.Name).Message("จำเป็นต้องกรอก ชื่อ")
  v.Required(user.Lastname).Message("จำเป็นต้องกรอก นามสกุล")
  v.Required(user.Password).Message("จำเป็นต้องกรอก รหัสผ่าน")
  v.Required(user.Validpassword).Message("จำเป็นต้องกรอก ยืนยันรหัสผ่าน")
  v.Required(user.Tel).Message("จำเป็นต้องกรอก เบอร์โทรศัพท์")
  v.Required(user.Validpassword == user.Password).Message("รหัสผ่านไม่ตรงกัน")
}
//RegisterUserChaokaset สมัครสมาชิก
func RegisterUserChaokaset(username string,password []byte,prefix string,name string,lastname string,tel string,role_user int) (result bool) { //result bool คือประกาศตัวแปรที่ใช้รีเทร์นค่่าเป็น boolean
  // connect to the cluster
     session, err := mgo.Dial("127.0.0.1")
     if err != nil {
         panic(err)
     }
     defer session.Close()
     //var user *User
     session.SetMode(mgo.Monotonic, true)
     qmgo := session.DB("chaokaset").C("users")
     //role_user 1>admin 2>officer 3>farmer 4>user
     pic := "http://simpleicon.com/wp-content/uploads/multy-user.svg"
     err = qmgo.Insert(&UserByChaokaset{Username: username, Password: password,Prefix: prefix,Name: name,Lastname: lastname,Tel: tel,Timestamp: time.Now(),Pic: pic,Role: role_user})
     if err != nil {
       return false
     }else{
       return true
     }

}

//GetUserData สำหรับเรียกข้อมูลผู้ใช้งาน
func GetUserData(Uusername string) *User {
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()
  var user *User
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("users")
  result := User{}
	qmgo.Find(bson.M{"username": Uusername}).One(&result)
  user = &User{Userid: result.Userid,Username: result.Username,Name: result.Name,Lastname: result.Lastname,Pic: result.Pic,Role: result.Role,Tel: result.Tel}
  return user
}
//GetEditUserData สำหรับเรียกข้อมูลแก้ไขผู้ใช้งาน
func GetEditUserData(Uusername string) *UserData {
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()
  var user *UserData
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("users")
  result := UserData{}
	qmgo.Find(bson.M{"username": Uusername}).One(&result)
  user = &UserData{Userid: result.Userid,Username: result.Username,Name: result.Name,Lastname: result.Lastname,Pic: result.Pic,Role: result.Role,Tel: result.Tel,Province: result.Province,Tumbon: result.Tumbon,Aumphur: result.Aumphur,Zipcode: result.Zipcode,Email: result.Email,Address: result.Address}
  return user
}

func EditUserData(Uusername string,prefix string,name string,lastname string,tel string,email string,province string,tumbon string,aumphur string,zipcode string,address string) (result bool) {
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()
  //var user *UserData
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("users")
  result = true
	err = qmgo.Update(bson.M{"username": Uusername}, bson.M{"$set": bson.M{"name": name,"tel" : tel,"email" : email,"province": province,"aumphur": aumphur,"tumbon" : tumbon,"zipcode": zipcode,"lastname" : lastname, "lastedit": time.Now(),"address": address}})
	if err != nil {
		panic(err)
	}
  return result
}

//CheckUserLogin สำหรับเรียกข้อมูลผู้ใช้งาน
func CheckUserLogin(Uusername string) *User{
  session, err := mgo.Dial("127.0.0.1")
  if err != nil {
      panic(err)
  }
  defer session.Close()
  var user *User
  session.SetMode(mgo.Monotonic, true)
  qmgo := session.DB("chaokaset").C("users")
  result := User{}
	qmgo.Find(bson.M{"username": Uusername}).One(&result)
  user = &User{Userid: result.Userid,Username: result.Username,Name: result.Name,Lastname: result.Lastname,Pic: result.Pic,Role: result.Role}
  return user
}

//GetPasswordUser สำหรับรับค่า รหัสผ่านของ User
func CheckPasswordUser(Uusername string,Upassword string) (result bool){
   session, err := mgo.Dial("127.0.0.1")
   if err != nil {
       panic(err)
   }
   defer session.Close()
   session.SetMode(mgo.Monotonic, true)
   qmgo := session.DB("chaokaset").C("users")
   UserResult := UserByChaokaset{}
 	 qmgo.Find(bson.M{"username": Uusername}).One(&UserResult)
   err = bcrypt.CompareHashAndPassword(UserResult.Password, []byte(Upassword))//ตรวจสอบรหัสผ่าน
   if err == nil{
      return true;
   } else {
      return false;
   }
}
