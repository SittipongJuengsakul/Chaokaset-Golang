package models
import (
    "log"
    "github.com/gocql/gocql"
)


type User struct { //สร้าง Struct
	Userid                                    int
	Username,Password,Validpassword           string
  Name,Lastname,Prefix,Tel                  string
  HashedPassword                            []byte
}
/*
var db = make(map[int]*User)

func GetUser(id int) *User {
	return db[id]
}


func NewUser() *User {
	user := &User{Uid: rand.Intn(10000)}
	db[user.Uid] = user
	return user
}*/
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
       result = true;
    } else{
      result = false;
    }
    return result
}
