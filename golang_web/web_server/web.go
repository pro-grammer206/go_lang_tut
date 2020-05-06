package main

import (
	"fmt"
	"log"
	"net/http"
	_"strings"
)

func funcHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to Go")
  fmt.Fprintf(w,string(r.URL.Path))

}

func main() {
	http.HandleFunc("/", funcHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
