package models
import (
    "log"
    "github.com/gocql/gocql"
    "golang.org/x/crypto/bcrypt"
)


type User struct { //สร้าง Struct
	Userid                                    int
	Username,Password,Validpassword           string
  Name,Lastname,Prefix,Tel                  string
  HashedPassword                            []byte
}

//RegisterUserChaokaset สมัครสมาชิก
func RegisterUserChaokaset(username string,password []byte,prefix string,name string,lastname string,tel string) (result bool) { //result bool คือประกาศตัวแปรที่ใช้รีเทร์นค่่าเป็น boolean
  // connect to the cluster
	 cluster := gocql.NewCluster("127.0.0.1")
	 cluster.Keyspace = "chaokaset"
	 cluster.Consistency = gocql.Quorum
	 session, _ := cluster.CreateSession()
	 defer session.Close()
    if err := session.Query("INSERT INTO users_by_chaokaset (userid,username, password, prefix, name, lname,tel) VALUES (uuid(),?, ?, ?, ?,?,?)",
        username, password, prefix, name, lastname,tel).Exec(); err != nil {
       log.Fatal(err)
       result = false;
    } else{
      result = true;
    }
    return result
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
  //Upassword = "sittipongsssss"
  err := bcrypt.CompareHashAndPassword(password, []byte(Upassword))//ตรวจสอบรหัสผ่าน
   if err == nil{
      return true;
   } else {
      return false;
   }
}
