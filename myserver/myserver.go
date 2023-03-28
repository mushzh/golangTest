package myserver

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func StartServer() error {
	handler := &MyHandler{}
	http.Handle("/", handler)
	return http.ListenAndServe(":8080", nil)
}

//func main() {
//	if err := myserver.StartServer(); err != nil {
//		log.Fatal(err)
//	}
//}
