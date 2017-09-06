package pro

import (
	"../gogogo"
)

type MyController struct {
	gogogo.Controller
	time int
}

func (this *MyController) ActionDemo() (str string) {
	str = "test"
	return
}

func NewMyController() (controller *MyController) {
	return new(MyController)
}
