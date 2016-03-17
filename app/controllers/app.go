package controllers

import "github.com/revel/revel"

//App for save Structure of Folder App (in views)
type App struct {
	*revel.Controller
}

//Search for save Structure of Folder Search (in views)
type Search struct {
	*revel.Controller
}

//Index for Create routing Page Index (localhost/index)
func (c App) Index() revel.Result {
	return c.Render()
}

//SearchPlant for Create routing Page Index (localhost/searchplant)
func (c Search) SearchPlant() revel.Result {
	return c.Render()
}

//Register for Create routing Page Register (localhost/register)
func (c App) Register() revel.Result {
	return c.Render()
}
