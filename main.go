package main

import (
	"fmt"
	"net/http"
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
		// Tip we can use r.FormValue("username") instead which
		// automatically calls ParseForm
		// downside: silences errors when key not found returning ""
		// and if multiple value present returns only first one.
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
