package main

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
)

const test = `'((mainbg . {{.bg1}}) (mainfg . {{ .fg1}}))`

func rgbToHex(r, g, b int64) (hex string) {
	red := strconv.FormatInt(r, 16)
	green := strconv.FormatInt(g, 16)
	blue := strconv.FormatInt(b, 16)
	colors := []string{red, green, blue}
	for i, j := range colors {
		if len(j) == 1 {
			colors[i] = "0" + j
		}
	}
	return "#" + colors[0] + colors[1] + colors[2]
}

func main() {
	temp := template.Must(template.New("test").Parse(test))
	temp.Execute(os.Stdout, map[string]interface{}{"bg1": "\"#000000\"",
		"fg1": "\"#ffffff\""})
	fmt.Println()
	fmt.Println(rgbToHex(255, 255, 255))
	fmt.Println(rgbToHex(0, 0, 0))
}
