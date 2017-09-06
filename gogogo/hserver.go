package gogogo

import (
	"net/http"
)

var Port, Dest string

// type HServer struct{
// 	Port string
// 	Dest string
// 	Handler
// }

func init() {

}

func SetDomainPort(des, port string) {
	// server := new(Hserver)
	Dest = des
	Port = port
	return
}

func StartServer() {
	http.HandleFunc("/", ServerHttp)
	http.ListenAndServe(Dest+":"+Port, nil)
}

func ServerHttp(w http.ResponseWriter, r *http.Request) {
	go func(w http.ResponseWriter, r *http.Request) {
		res := GetResult(ParseRoute(r))
		w.Write([]byte(res))
	}(w, r)
}
