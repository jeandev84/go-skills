package main

import (
	"html/template"
	"os"
)

var tpl = `
======== {{.Name}} =======
Grade: {{.Grade}}
Attendance: {{.Attendance}}
`

func main() {
	t, _ := template.New("student").Parse(tpl)

	s := struct {
		Name       string
		Grade      int
		Attendance int
	}{
		"Bob",
		4,
		8,
	}

	t.Execute(os.Stdout, s)
}
