package emacsthemecreator

import (
	"bytes"
	"fmt"
	htemplate "html/template"
	"net/http"
	"text/template"

	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/savetheme", saveThemeHandler)
}

// selectedColors - takes an http.Request, and maps of its
// submitted Form every face name to it's value.
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
	facemap["author"] = r.FormValue("authorname")
	facemap["url"] = r.FormValue("url")
	return facemap
}

// addColors - takes a map of faces, and stores darker
// and lighter variants of the main fg and bg face, as
// well as a lighter and darker variant of the keyword face.
func addColors(colors map[string]string) map[string]string {
	fg := colors["deffacefg"]
	bg := colors["deffacebg"]
	keyword := colors["keywordface"]
	bg2 := ""
	bg3 := ""
	bg4 := ""
	fg2 := ""
	fg3 := ""
	fg4 := ""
	fgcol, _ := colorful.Hex(fg)
	bgcol, _ := colorful.Hex(bg)
	keycol, _ := colorful.Hex(keyword)
	if hasDarkBg(&bgcol) {
		fg2 = darken(fgcol, 0.08)
		fg3 = darken(fgcol, 0.16)
		fg4 = darken(fgcol, 0.24)
		bg2 = lighten(bgcol, 0.08)
		bg3 = lighten(bgcol, 0.16)
		bg4 = lighten(bgcol, 0.24)
	} else {
		fg2 = lighten(fgcol, 0.08)
		fg3 = lighten(fgcol, 0.16)
		fg4 = lighten(fgcol, 0.24)
		bg2 = darken(bgcol, 0.08)
		bg3 = darken(bgcol, 0.16)
		bg4 = darken(bgcol, 0.24)
	}
	key2 := lighten(keycol, 0.11)
	key3 := darken(keycol, 0.11)
	colors["fore2"] = fg2
	colors["fore3"] = fg3
	colors["fore4"] = fg4
	colors["back2"] = bg2
	colors["back3"] = bg3
	colors["back4"] = bg4
	colors["key2"] = key2
	colors["key3"] = key3
	return colors
}

// darken - return a darker variant of a supplied color.
func darken(col colorful.Color, factor float64) string {
	black, _ := colorful.Hex("#000000")
	return col.BlendLab(black, factor).Hex()
}

// lighten - return a lighter variant of a supplied color.
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
	themefile := template.Must(template.New("theme").ParseFiles("themetemplate.txt"))
	if err := themefile.ExecuteTemplate(&res, "theme", fmap); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	str := res.String()
	webview := htemplate.Must(htemplate.New("_displaytheme").ParseFiles("displaytheme.html"))
	if err := webview.ExecuteTemplate(w, "_displaytheme", map[string]interface{}{"emacstheme": str}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func handleTheme(r *http.Request) (interface{}, error) {

	return nil, fmt.Errorf("method not implemented")
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "success")
}
