package emacsthemecreator

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lucasb-eyer/go-colorful"
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
		col, _ := colorful.Hex(i.old)
		actual := lighten(col, i.factor)
		assert.Equal(actual, i.lighter)
	}
}
func TestDarken(t *testing.T) {
	assert := assert.New(t)
	for _, i := range darkenedColors {
		col, _ := colorful.Hex(i.old)
		actual := darken(col, i.factor)
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

func TestHasDarkBg(t *testing.T) {
	assert := assert.New(t)
	col, _ := colorful.Hex("#000000")
	col1, _ := colorful.Hex("#002b36")
	col2, _ := colorful.Hex("#586e75")
	col3, _ := colorful.Hex("#fdf6e3")
	assert.True(hasDarkBg(&col))
	assert.True(hasDarkBg(&col1))
	assert.True(hasDarkBg(&col2))
	assert.False(hasDarkBg(&col3))
}
func TestRandomCol(t *testing.T) {
	assert := assert.New(t)
	rands, _ := randomColWarm()
	assert.IsType(rands.Randkey, "string")
	assert.IsType(rands.Randconst, "string")
	assert.Contains(rands.Randconst, "#")
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
