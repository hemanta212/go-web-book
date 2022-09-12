package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
		// Tip we can use r.Form.Get("username") instead which
		// is better because for all fields it works and
		// returns "" if empty.
		// For checkboxes and radio buttons the current method errors out
		// but it will only return first item for mutiple values tho
		if len(r.Form["username"][0]) == 0 {
			t, _ := template.ParseFiles("login.gtpl")
			t.Execute(w, nil)
		}
		// Number validation
		getint, err = strconv.Atoi(r.Form.Get("password"))

	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":5000", nil)
}
