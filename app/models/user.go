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
func RegisterUserChaokaset(username string,password string,prefix string,name string,lastname string,tel string) (result bool) {
  // connect to the cluster
	 cluster := gocql.NewCluster("127.0.0.1")
	 cluster.Keyspace = "chaokaset"
	 cluster.Consistency = gocql.Quorum
	 session, _ := cluster.CreateSession()
	 defer session.Close()

    if err := session.Query("INSERT INTO users (userid,username, password, prefix, name, lastname,tel) VALUES (uuid(), ?, ?, ?, ?,?,?)",
        username, password, prefix, name, lastname,tel).Exec(); err != nil {
       log.Fatal(err)
       result = true;
    } else{
      result = false;
    }
    return result
}
