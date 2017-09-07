package main

import (
	"./gogogo"
	"./pro"
)

/**
Go Archture
Controller
	Every Controller neeed Hook
	ManageController single model
	1.Route ablity

View
	Same as htm; source
Model
	Model base on mysql
Service

*/
func main() {

	controller := pro.NewMyController()
	gogogo.Manage.SetRouteMapWithController(*controller)
	gogogo.Manage.Run()
}
