package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
	} else if r.URL.Path == "/test" {
		fmt.Fprint(w, "Test passed")
	} else {
		http.NotFound(w, r)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":5000", mux)
}
