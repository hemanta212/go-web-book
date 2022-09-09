package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// parse arguments, have to manually call this
	r.ParseForm()
	// form info in server side
	fmt.Println(r.Form)

	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("Key: ", k)
		fmt.Println("Val: ", strings.Join(v, " "))
	}
	fmt.Fprint(w, "Hello Pykancha!")
}

func main() {
	// router set
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:  ", err)
	}
}
