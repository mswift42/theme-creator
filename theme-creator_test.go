package emacsthemecreator

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mswift42/go-colorful"
	"github.com/stretchr/testify/assert"
)

var lightendedColors = []struct {
	old     string
	lighter string
	factor  float64
}{
	{
		old:     "#000000",
		lighter: "#303030",
		factor:  0.20,
	},
	{
		old:     "#202020",
		lighter: "#464646",
		factor:  0.20,
	},
	{
		old:     "#34b156",
		lighter: "#67c177",
		factor:  0.2,
	},
	{
		old:     "#23557a",
		lighter: "#547493",
		factor:  0.2,
	},
	{
		old:     "#ffffff",
		lighter: "#ffffff",
		factor:  0.1,
	},
	{
		old:     "#0e2f17",
		lighter: "#3a543f",
		factor:  0.2,
	},
	{
		old:     "#721d2e",
		lighter: "#a06066",
		factor:  0.3,
	},
	{
		old:     "#8a2436",
		lighter: "#a5525a",
		factor:  0.2,
	},
}
var darkenedColors = []struct {
	old    string
	darker string
	factor float64
}{
	{
		old:    "#ffffff",
		darker: "#e2e2e2",
		factor: 0.1,
	},
	{
		old:    "#e1e1e1",
		darker: "#d4d4d4",
		factor: 0.05,
	},
	{
		old:    "#808080",
		darker: "#727272",
		factor: 0.1,
	},
	{
		old:    "#ff40ff",
		darker: "#f13ef1",
		factor: 0.05,
	},
	{
		old:    "#b72fc1",
		darker: "#ad2eb6",
		factor: 0.05,
	},
	{
		old:    "#000000",
		darker: "#000000",
		factor: 0.05,
	},
	{
		old:    "#a5525a",
		darker: "#824248",
		factor: 0.2,
	},
	{
		old:    "#dab071",
		darker: "#93784f",
		factor: 0.3,
	},
}

func TestLighten(t *testing.T) {
	assert := assert.New(t)
	for _, i := range lightendedColors {
		c, _ := colorful.Hex(i.old)
		actual := lighten(c, i.factor)
		assert.Equal(actual, i.lighter)
	}
}
func TestDarken(t *testing.T) {
	assert := assert.New(t)
	for _, i := range darkenedColors {
		c, _ := colorful.Hex(i.old)
		actual := darken(c, i.factor)
		assert.Equal(actual, i.darker)
	}
}

func TestAddColors(t *testing.T) {
	assert := assert.New(t)
	cmap := map[string]string{"deffacefg": "#000000",
		"deffacebg": "#e2dfd9", "keywordface": "#9e5529"}
	add := addColors(cmap)
	assert.Equal("#181818", add["fore2"])
	assert.Equal("#cecbc6", add["back2"])
	assert.Equal("#bab8b3", add["back3"])
	assert.Equal("#ab673f", add["key2"])
	assert.Equal("#8c4c26", add["key3"])
}

var darkBg = []struct {
	bg  string
	yes bool
}{
	{
		bg:  "#000000",
		yes: true,
	},
	{
		bg:  "#ffffff",
		yes: false,
	},
	{
		bg:  "#002b36",
		yes: true,
	},
	{
		bg:  "#586e75",
		yes: true,
	},
	{
		bg:  "#fdf6e3",
		yes: false,
	},
	{
		bg:  "#d2d2d2",
		yes: false,
	},
	{
		bg:  "#daa1e6",
		yes: false,
	},
	{
		bg:  "#932ad7",
		yes: true,
	},
	{
		bg:  "#d7cf47",
		yes: false,
	},
	{
		bg:  "#87e37e",
		yes: false,
	},
	{
		bg:  "#f0c1bc",
		yes: false,
	},
	{
		bg:  "#dc6c60",
		yes: false,
	},
	{
		bg:  "#39c52b",
		yes: false,
	},
}

func TestHasDarkBg(t *testing.T) {
	assert := assert.New(t)
	for _, i := range darkBg {
		col, _ := colorful.Hex(i.bg)
		actual := hasDarkBg(&col)
		assert.Equal(actual, i.yes)
	}
}
func TestRandomColWarm(t *testing.T) {
	assert := assert.New(t)
	rands, _ := randomColWarm()
	assert.IsType(rands.Randkey, "string")
	assert.IsType(rands.Randconst, "string")
	assert.Contains(rands.Randconst, "#")
}
func TestRandomColHappy(t *testing.T) {
	assert := assert.New(t)
	rands, _ := randomColHappy()
	assert.IsType(rands.Randkey, "string")
	assert.IsType(rands.Randconst, "string")
	assert.Contains(rands.Randbuiltin, "#")
}
func TestRandomColSoft(t *testing.T) {
	assert := assert.New(t)
	rands, _ := randomColSoft()
	assert.IsType(rands.Randkey, "string")
	assert.IsType(rands.Randconst, "string")
	assert.Contains(rands.Randbuiltin, "#")
}

func TestRandomColorHelper(t *testing.T) {
	assert := assert.New(t)
	col1, _ := colorful.Hex("#000000")
	col2, _ := colorful.Hex("#111111")
	col3, _ := colorful.Hex("#222222")
	col4, _ := colorful.Hex("#333333")
	col5, _ := colorful.Hex("#444444")
	col6, _ := colorful.Hex("#555555")
	col7, _ := colorful.Hex("#666666")
	pal := []colorful.Color{col1, col2, col3, col4, col5, col6, col7}
	rand := randomColorHelper(pal)
	assert.Equal(rand.Randkey, "#000000")
	assert.Equal(rand.Randbuiltin, "#111111")
	assert.Equal(rand.Randtype, "#555555")
}

func TestRandomColorHandler(t *testing.T) {
	resp := httptest.NewRecorder()
	uri := "/randomcolorswarm"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	http.DefaultServeMux.ServeHTTP(resp, req)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return error: %s", p)
		} else if !strings.Contains(string(p), "randbuiltin") {
			t.Errorf("header response doesn't match:\n%s", p)
		} else if !strings.Contains(string(p), "randkey") {
			t.Errorf("header response doesn't match:\n%s", p)
		}
	}
}

func TestRandomColorHappyHandler(t *testing.T) {
	resp := httptest.NewRecorder()
	uri := "/randomcolorshappy"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	http.DefaultServeMux.ServeHTTP(resp, req)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return error: %s", p)
		} else if !strings.Contains(string(p), "randbuiltin") {
			t.Errorf("header response doesn't match:\n%s", p)
		}
	}
}
func TestRandomColorSoftHandler(t *testing.T) {
	resp := httptest.NewRecorder()
	uri := "/randomcolorssoft"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	http.DefaultServeMux.ServeHTTP(resp, req)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return error: %s", p)
		} else if !strings.Contains(string(p), "randkey") {
			t.Errorf("header response doesn't match:\n%s", p)
		}
	}
}
func TestHandler(t *testing.T) {
	resp := httptest.NewRecorder()
	uri := "/"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	http.DefaultServeMux.ServeHTTP(resp, req)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return error: %s", p)
		} else if !strings.Contains(string(p), "success") {
			t.Errorf("header response doesn't match:\n%s", p)
		}
	}
}
func TestSaveThemeHandler(t *testing.T) {
	resp := httptest.NewRecorder()
	uri := "/savetheme"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	bodystrings := []string{"custom-theme-set-faces", "font-lock-builtin-face", "warning", "Code", "Emacs", "deftheme", "min-colors"}
	http.DefaultServeMux.ServeHTTP(resp, req)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return error: %s", p)
		}
		for _, i := range bodystrings {
			if !strings.Contains(string(p), i) {
				t.Errorf("header response doesn't match:\n%s", p)
			}
		}
		if strings.Contains(string(p), "sportschau") {
			t.Errorf("header response doesn't match:\n%s", p)
		}
	}
}
