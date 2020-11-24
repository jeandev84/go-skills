package main

import (
	"html/template"
	"os"
)

var tpl = `
======= {{.Name}} =======
Grade: {{.Grade}}
Attendance: {{.Attendance}}
School: {{.School}}
`

func main() {

	t, _ := template.New("student").Parse(tpl)

	s := struct {
		Name       string
		Grade      int
		Attendance int
		School     string
	}{
		"Bob",
		4,
		8,
		"A",
	}

	t.Execute(os.Stdout, s)
}
