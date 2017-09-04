package gogogo

import (
	"error"
	"reflect"
	"net/http"
)

type  Controller interface{

}

var Manage ManageController

type ManageController struct{
	 routeMap map[string]Controller

}	

func init() {
	Manage := newManageController()
}

func newManageController()(m * MangeController) {
	m = new(ManageController)
	m.routeMap = make(map[string]Controller)
}

func (this * ManageController) SetRouteMapWithController(controller Controller) {
		v := reflect.ValueOf(controller)
		this.routeMap[reflect.TypeOf(controller) = controller		
}

func GetRouteMapString(controllerName string) (val Controller, isPresent bool) {
		return val,isPresent := this.routeMap[controllerName]
}


func run() {
	server := NewHserver()
	server.Start()
}

func initHttp(){

}


