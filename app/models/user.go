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
func RegisterUserChaokaset(name string,lastname string,tel string,password string) (err int) {
  // connect to the cluster
	 cluster := gocql.NewCluster("127.0.0.1")
	 cluster.Keyspace = "demo"
	 cluster.Consistency = gocql.Quorum
	 session, _ := cluster.CreateSession()
	 defer session.Close()

    if err := session.Query("INSERT INTO users (lastname, age, city, email, firstname) VALUES ('aaaa', 38, 'Austin', 'bob@example.com', 'ddd')").Exec(); err != nil {
      log.Fatal(err)
     }
    return err
}
