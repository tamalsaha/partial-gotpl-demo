package main
import (
	"os"
	"text/template"
)
func main() {
	type Inner struct {
		A string
	}
	type Outer struct {
		Inner
	}
	o := Outer{
		Inner: Inner{A : "123"},
	}
	tpl := template.Must(template.New("").Option("missingkey=error").Parse(`{{.A}} {{ .B }}`))
	err := tpl.Execute(os.Stdout, &o)
	if err != nil {
		panic(err)
	}
}