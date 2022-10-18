package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		// to prevent scripts injcetion
		escapedUsername := template.HTMLEscapeString(r.Form.Get("username"))

		fmt.Println("username: ", escapedUsername)
		fmt.Println("age: ", r.Form["age"])
		fmt.Println("password: ", r.Form["password"])
		fmt.Println("fruits: ", r.Form["fruit"])
		fmt.Println("gender: ", r.Form["gender"])
		fmt.Println("Interests: ", r.Form["interest"])
		// Tip we can use r.FormValue("username") instead which
		// automatically calls ParseForm
		// downside: silences errors when key not found returning ""
		// and if multiple value present returns only first one.

		if len(escapedUsername) == 0 {
			r.Method = "GET"
			login(w, r)
		}
		getint, _ := strconv.Atoi(r.Form["age"][0])
		if getint <= 0 || getint > 100 {
			r.Method = "GET"
			login(w, r)
		}
		if !validateFruits(r) {
			r.Method = "GET"
			login(w, r)
		}
		if !validateMaleFemale(r) {
			r.Method = "GET"
			login(w, r)
		}
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
		// sayUsername(w, escapedUsername)

		// if Want to intentionally display unescaped scripts then
		t, _ := template.New("foo").Parse(`{{define "T"}} Hello, {{.}}!{{end}}`)
		t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))

		// checkbox validation is different, should do a Set Diffrentiation A-B, and check for nil
		// interests := []string{"football", "cricket", "tennis"}
		// slice_diff := Slice_diff(r.Form["interest"], interests)
		// if slice_diff == nil { return true}
		// return false

	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
func sayUsername(w http.ResponseWriter, username string) {
	fmt.Fprint(w, username)
}

func validateFruits(r *http.Request) bool {
	fruits := []string{"apple", "banana", "pear"}
	for _, v := range fruits {
		if v == r.Form.Get("fruit") {
			return true
		}
	}
	return false
}

func validateMaleFemale(r *http.Request) bool {
	maleFemale := []int{1, 2}
	for _, v := range maleFemale {
		gender, _ := strconv.Atoi(r.Form.Get("gender"))
		if v == gender {
			return true
		}
	}
	return false
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":5000", nil)
}
