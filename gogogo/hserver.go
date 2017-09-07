package gogogo

import (
	"io"
	"net/http"
	"strconv"
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
	func(w http.ResponseWriter, r *http.Request) {
		res := GetResult(ParseRoute(r))
		w.Header().Add("Content-Length", strconv.Itoa(len(res)))
		_, err := io.WriteString(w, res)
		CheckError(err)
	}(w, r)
}
