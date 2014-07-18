package emacsthemecreator

import (
	"bytes"
	"fmt"
	htemplate "html/template"
	"net/http"
	"text/template"

	"github.com/lucasb-eyer/go-colorful"
)

const test = `'((mainbg . {{.bg1}}) (mainfg . {{ .fg1}}))`

// func decodeTheme(r io.ReadCloser) (*Theme, error) {
// 	defer r.Close()
// 	var theme Theme
// 	err := json.NewDecoder(r).Decode(&theme)
// 	return &theme, err
// }

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/savetheme", saveThemeHandler)
}
func selectedColors(r *http.Request) map[string]string {
	facemap := make(map[string]string)
	faces := []string{"deffacebg", "deffacefg", "builtinface", "stringface",
		"keywordface", "functionnameface", "commentface", "variableface",
		"typeface", "constantface", "warningface"}
	for _, i := range faces {
		facemap[i] = r.FormValue(i)
	}
	facemap = addColors(facemap)
	for key, value := range facemap {
		facemap[key] = "\"" + value + "\""
	}
	facemap["themename"] = r.FormValue("themename")
	return facemap
}
func addColors(colors map[string]string) map[string]string {
	fg := colors["deffacefg"]
	bg := colors["deffacebg"]
	bg2 := ""
	bg3 := ""
	fg2 := ""
	fg3 := ""
	fgcol, _ := colorful.Hex(fg)
	bgcol, _ := colorful.Hex(bg)
	if hasDarkBg(&bgcol) {
		fg2 = darken(fgcol, 0.05)
		fg3 = darken(fgcol, 0.1)
		bg2 = lighten(bgcol, 0.05)
		bg3 = lighten(bgcol, 0.1)

	} else {
		fg2 = lighten(fgcol, 0.05)
		fg3 = lighten(fgcol, 0.1)
		bg2 = darken(bgcol, 0.05)
		bg3 = darken(bgcol, 0.1)
	}
	colors["fore2"] = fg2
	colors["fore3"] = fg3
	colors["back2"] = bg2
	colors["back3"] = bg3
	return colors
}

func darken(col colorful.Color, factor float64) string {
	black, _ := colorful.Hex("#000000")
	return col.BlendLab(black, factor).Hex()
}

func lighten(col colorful.Color, factor float64) string {
	white, _ := colorful.Hex("#ffffff")
	return col.BlendLab(white, factor).Hex()
}

// hasDarkBg - Check if supplied color is dark.
func hasDarkBg(c *colorful.Color) bool {
	l, _, _ := c.Lab()
	return l < 0.5
}

func saveThemeHandler(w http.ResponseWriter, r *http.Request) {

	fmap := selectedColors(r)
	var res bytes.Buffer
	themefile := template.Must(template.New("theme").ParseFiles("static/themetemplate.txt"))
	if err := themefile.ExecuteTemplate(&res, "theme", fmap); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	str := res.String()
	webview := htemplate.Must(htemplate.New("_displaytheme").ParseFiles("static/displaytheme.html"))
	if err := webview.ExecuteTemplate(w, "_displaytheme", map[string]interface{}{"emacstheme": str}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func handleTheme(r *http.Request) (interface{}, error) {
	// switch r.Method {
	// case "POST":
	// //	theme, err := decodeTheme(r.Body)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	fmt.Println(theme)

	// case "GET":
	// 	var theme Theme
	// 	return theme, nil
	// }

	return nil, fmt.Errorf("method not implemented")
}
func handler(w http.ResponseWriter, r *http.Request) {
	// val, err := handleTheme(r)
	// if err == nil {
	// 	fmt.Fprintf(w, "Success")
	// 	fmt.Println(val)
	// }
	fmt.Fprintf(w, "success")
}

// func main() {
// 	temp := template.Must(template.New("test").Parse(test))
// 	temp.Execute(os.Stdout, map[string]interface{}{"bg1": "\"#000000\"",
// 		"fg1": "\"#ffffff\""})
// 	fmt.Println()
// 	fmt.Println(rgbToHex(255, 255, 255))
// 	fmt.Println(rgbToHex(0, 0, 0))
// 	fmt.Println(122 * 120 / 100)
// }
