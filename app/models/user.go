package models
import (
    "log"
    "github.com/gocql/gocql"
)


/*type User struct {
	Uid         int
	AccessToken string
}

var db = make(map[int]*User)

func GetUser(id int) *User {
	return db[id]
}


func NewUser() *User {
	user := &User{Uid: rand.Intn(10000)}
	db[user.Uid] = user
	return user
}*/
func RegisterUserChaokaset(username string,password string,prefix string,name string,lastname string,tel string) (result int) {
  // connect to the cluster
	 cluster := gocql.NewCluster("127.0.0.1")
	 cluster.Keyspace = "demo"
	 cluster.Consistency = gocql.Quorum
	 session, _ := cluster.CreateSession()
	 defer session.Close()

    if err := session.Query("INSERT INTO users (username, password, prefix, name, lastname,tel) VALUES (?, ?, ?, ?, ?,?)",
       username, password, prefix, name, lastname,tel).Exec(); err != nil {
       log.Fatal(err)
    } else {
 		   log.Fatal(err)
 	  }
    return result
}
