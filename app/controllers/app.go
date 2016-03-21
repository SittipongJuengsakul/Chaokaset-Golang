package controllers

import (
    "github.com/revel/revel"
		"log"
    "github.com/gocql/gocql"
)

//App for save Structure of Folder App (in views)
type App struct {
	*revel.Controller
}

//Search for save Structure of Folder Search (in views)
type Search struct {
	*revel.Controller
}

//Auth for save Structure of Folder Authen (in views)
type Auth struct {
	*revel.Controller
}

//Index for Create routing Page Index (localhost/index)
func (c App) Index() revel.Result {
	return c.Render()
}

//Templates for Example Template (localhost/template)
func (c App) Templates() revel.Result {
	return c.Render()
}

//SearchPlant for Create routing Page Index (localhost/searchplant)
func (c Search) SearchPlant() revel.Result {
	return c.Render()
}

//Login for Create routing Page Login (localhost/login)
func (c Auth) Login() revel.Result {
	// connect to the cluster
	 cluster := gocql.NewCluster("127.0.0.1")
	 cluster.Keyspace = "demo"
	 cluster.Consistency = gocql.Quorum
	 session, _ := cluster.CreateSession()
	 defer session.Close()
	 // ทดลองคิวรี่
	 if err := session.Query("INSERT INTO users (lastname, age, city, email, firstname) VALUES ('Joneseee', 38, 'Austin', 'bob@example.com', 'Bob')").Exec(); err != nil {
		 log.Fatal(err)
	}
	return c.Render()
}

//Register for Create routing Page Register (localhost/register)
func (c Auth) Register() revel.Result {
	return c.Render()
}
