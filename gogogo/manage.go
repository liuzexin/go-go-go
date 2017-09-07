package gogogo

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Controller interface {
}

var Manage *ManageController

type ManageController struct {
	routeMap map[string]Controller
	//host,port read config
}

func init() {
	Manage = newManageController()

}

func newManageController() (m *ManageController) {
	m = new(ManageController)
	m.routeMap = make(map[string]Controller)
	return
}

//TODO:: atomic
func (this *ManageController) SetRouteMapWithController(controller Controller) {
	v := reflect.TypeOf(controller).Name()
	fmt.Println(v)
	fmt.Println(controller)
	this.routeMap[v] = controller
}

func (this *ManageController) GetRouteMapString(controllerName string) (val Controller, isPresent bool) {
	val, isPresent = this.routeMap[controllerName]
	return
}

//star sever
func (this *ManageController) Run() {
	SetDomainPort("127.0.0.1", "20000")
	StartServer()
}

func GetResult(route []string) (str string) {
	controller := Manage.routeMap[route[0]]
	v := reflect.New(reflect.TypeOf(controller))
	fmt.Println(reflect.ValueOf(&controller))
	fmt.Println(reflect.ValueOf(controller))
	method := v.MethodByName(route[1])
	fmt.Println(method)
	//Now only need string
	// f := v.Func
	// res := f.Call([]reflect.Value{reflect.Zero(reflect.TypeOf(new(Controller)))})
	fmt.Println(v.Interface())
	res := method.Call([]reflect.Value{})
	str, _ = res[0].Interface().(string)
	return
}

func ParseRoute(r *http.Request) (route []string) {
	path := r.URL.Path
	var separte rune = '/'
	f := func(c rune) bool {
		return c == separte
	}
	route = strings.FieldsFunc(path, f)
	fmt.Println(route)
	route[0] = strings.Title(route[0]) + "Controller"
	route[1] = "Action" + strings.Title(route[1])
	return
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
