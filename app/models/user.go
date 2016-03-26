package models
import (
    "log"
    "github.com/gocql/gocql"
    "golang.org/x/crypto/bcrypt"
    "regexp"
)


type User struct { //สร้าง Struct
	Userid                                    gocql.UUID
	Username,Password,Validpassword           string
  Name,Lastname,Prefix,Tel,Pic              string
  HashedPassword                            []byte
}
var userdb = make(map[string]*User)
/*
//ส่วน Validation Form
var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).
		Key("user.Password")

	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}*/
//RegisterUserChaokaset สมัครสมาชิก
func RegisterUserChaokaset(username string,password []byte,prefix string,name string,lastname string,tel string) (result bool) { //result bool คือประกาศตัวแปรที่ใช้รีเทร์นค่่าเป็น boolean
  // connect to the cluster
	 cluster := gocql.NewCluster("127.0.0.1")
	 cluster.Keyspace = "chaokaset"
	 cluster.Consistency = gocql.Quorum
	 session, _ := cluster.CreateSession()
	 defer session.Close()
   pic := "http://simpleicon.com/wp-content/uploads/multy-user.svg"
   //role_user 1>admin 2>officer 3>farmer 4>user
   role_user := 3
   if err := session.Query("INSERT INTO users_by_chaokaset (userid,username, password, prefix, name, lname,tel,pic,role_user) VALUES (uuid(),?, ?, ?, ?,?,?,?,?)",
        username, password, prefix, name, lastname,tel,pic,role_user).Exec(); err != nil {
       log.Fatal(err)
       result = false;
    } else{
      result = true;
    }
    return result
}

//GetUserData สำหรับเรียกข้อมูลผู้ใช้งาน
func GetUserData(Uusername string) *User {

  cluster := gocql.NewCluster("127.0.0.1")
  cluster.Keyspace = "chaokaset"
  cluster.Consistency = gocql.Quorum
  session, _ := cluster.CreateSession()
  defer session.Close()
  var username,name,lname,pic,prefix,tel string
  var userid gocql.UUID
  if err := session.Query(`SELECT username,userid,name,lname,pic,tel,prefix FROM users_by_chaokaset WHERE username = ? LIMIT 1 ALLOW FILTERING`,
        Uusername).Scan(&username,&userid,&name,&lname,&pic,&tel,&prefix ); err != nil {
        log.Fatal(err)
  }
  user := &User{Userid: userid,Username: username,Name: name,Lastname: lname,Pic: pic,Tel: tel,Prefix: prefix}
  return user
}

//CheckUserLogin สำหรับเรียกข้อมูลผู้ใช้งาน
func CheckUserLogin(Uusername string) *User{

  cluster := gocql.NewCluster("127.0.0.1")
  cluster.Keyspace = "chaokaset"
  cluster.Consistency = gocql.Quorum
  session, _ := cluster.CreateSession()
  defer session.Close()

  var username string
  if err := session.Query(`SELECT username FROM users_by_chaokaset WHERE username = ? LIMIT 1 ALLOW FILTERING`,
        Uusername).Scan(&username); err != nil {
        return nil;
  } else {
    user := &User{Username: username}
  	userdb[user.Username] = user
  	return user
  }
}

//GetPasswordUser สำหรับรับค่า รหัสผ่านของ User
func CheckPasswordUser(Uusername string,Upassword string) (result bool){
  cluster := gocql.NewCluster("127.0.0.1")
  cluster.Keyspace = "chaokaset"
  cluster.Consistency = gocql.Quorum
  session, _ := cluster.CreateSession()
  defer session.Close()
  var username string
  var password []byte
  session.Query(`SELECT username,password FROM users_by_chaokaset WHERE username = ? LIMIT 1 ALLOW FILTERING`,
        Uusername).Scan(&username,&password);
  err := bcrypt.CompareHashAndPassword(password, []byte(Upassword))//ตรวจสอบรหัสผ่าน
   if err == nil{
      return true;
   } else {
      return false;
   }
}
