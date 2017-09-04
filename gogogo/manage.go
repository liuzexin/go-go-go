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
	 //host,port read config
}	

func init() {
	Manage := newManageController()

}

func newManageController()(m * MangeController) {
	m = new(ManageController)
	m.routeMap = make(map[string]Controller)
}

//TODO:: atomic
func (this * ManageController) SetRouteMapWithController(controller Controller) {
		v := reflect.ValueOf(controller)
		this.routeMap[reflect.TypeOf(controller) = controller		
}

func (this * ManageController) GetRouteMapString(controllerName string) (val Controller, isPresent bool) {
		return val,isPresent := this.routeMap[controllerName]
}

//star sever
func (this * ManageController)Run(){

}

func GetResult(route []string)(str string){
	controller := Manage.routeMap[route[0]]
	v := reflect.ValueOf(controller)
	v =  v.MethodByName(route[1])
	//Now only need string
	res := v.Call()
	str,_ := res.(string)
	return 
}

func ParseRoute(r * Request)(route []string) {
	err := nil
	path := r.Url.Path
	route := path.Fields("/")
	route[0] = strings.Title(route[0])
	route[1] = strings.Title(route[1])
	return 
}
