package gogogo

import (
	"net/http"
	"strings"
)

type HServer struct{
	Port string
	Dest string
	Handler
}

func init() {

}



func NewHsever(des, port string)(server * Hserver) {
	server := new(Hserver)
	http.ListenAndServe(str + ":" + port, nil)
	http.HandleFunc("/", server)
	return 
}

func (this * Hserver)ServeHttp(w ResponseWriter, r * Request){
	go func (w ResponseWriter, r * Requset ) {
		res := Manage.GetResult(Manage.PaseRoute(r))
		w.Write([]byte(res))
	}(w, r)
}
