package templates

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func SimpleInsertExample() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}! {{.Emails}}")
	p := Person{UserName: "hemanta212", Emails: []string{"a@a.com"}}
	t.Execute(os.Stdout, p)
}

func NestedFieldsInsertion() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}

	t := template.New("nested fieldnames")
	t, _ = t.Parse(`hello {{.UserName}}!
                {{range .Emails}}
                    an email {{.}}
                {{end}}
                {{with .Friends}}
                {{range .}}
                    my  friend name is {{.Fname}}
                {{end}}
                {{end}}
	        `)
	p := Person{UserName: "hemanta212",
		Emails:  []string{"pykancha@gmail.com", "sharmahemanta.212@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

func ConditionsExample() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(
		tEmpty.Parse(`:: Empty pipeline if demo:
{{if ""}}
  will not be outputted.
{{end}}
`))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(
		tWithValue.Parse(`:: Not empty pipeline if demo:
{{if "anything"}}
  will be outputted.
{{end}}
`))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(
		tIfElse.Parse(`:: if-else demo:
{{if .}}
  if part
{{else}}
  else part
{{end}}
`))
	tIfElse.Execute(os.Stdout, "data")
	tIfElse.Execute(os.Stdout, nil)
}

func FunctionExample() {
	t := template.New("template functions")
	t = t.Funcs(template.FuncMap{"emailformat": emailFormat})
	t, _ = t.Parse(`Here is your {{. | emailformat}}`)
	t.Execute(os.Stdout, `sharmahemanta.212@gmail.com`)
}

func emailFormat(args ...interface{}) string {
	fallbackStr := fmt.Sprint(args...)
	var email string
	ok := false
	if len(args) == 1 {
		email, ok = args[0].(string)
	}
	if !ok {
		return fallbackStr
	}

	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return (parts[0] + " at " + parts[1])
	} else {
		return email
	}
}

func MustParseExample() {
	t := template.New("will parse")
	t = template.Must(t.Parse(`hello {{ . }}`))
	t.Execute(os.Stdout, "data")
	fmt.Println(":: First ✅")

	template.Must(template.New("test").Parse(`some text and bad comment /* comment`))
	fmt.Println(":: Second ✅")

	t1 := template.New("wont parse")
	t1 = template.Must(t1.Parse(`hello {{ . }`))
	fmt.Println(":: Third ✅")

}

func Subtemplates() {
	allfiles, err := parseTmplFiles("./templates")
	if err != nil {
		fmt.Println(err)
		return
	}

	allTemplates, err := template.ParseFiles(allfiles...)

	s1 := allTemplates.Lookup("header.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s2 := allTemplates.Lookup("content.tmpl")
	s2.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s3 := allTemplates.Lookup("footer.tmpl")
	s3.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	// wont output anything since there is no default sub template
	s3.Execute(os.Stdout, nil)
}

func parseTmplFiles(dir string) ([]string, error) {
	allfiles := []string{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".tmpl") {
			allfiles = append(allfiles, "./templates/"+filename)
		}
	}
	return allfiles, nil
}
