package main

import (
	"os"
	"text/template"
)

const test = `'((mainbg . {{.bg1}}) (mainfg . {{ .fg1}}))`

func main() {
	temp := template.Must(template.New("test").Parse(test))
	temp.Execute(os.Stdout, map[string]interface{}{"bg1": "\"#000000\"",
		"fg1": "\"#ffffff\""})
}
