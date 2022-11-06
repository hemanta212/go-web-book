package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hemanta212/webapp/session"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		currtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// check token validity
		} else {
			// give error if no token
		}
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
		getint, _ := strconv.Atoi(r.Form.Get("age"))
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

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		currtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		w.Write([]byte("File uploaded sucessfully"))
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

func runServer() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":5000", nil)
}

func main() {
	// runServer()
	// runClient()
	// sqlLogin()
	// sqliteConnect()
	// postgresConnect()
	// ormConnect()
	// redisConnect()
	// mongoConnect()
}
