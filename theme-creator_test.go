package emacsthemecreator

import (
	"testing"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/stretchr/testify/assert"
)

var convertColors = []struct {
	r, g, b int64
	hex     string
}{
	{
		// white
		r:   255,
		g:   255,
		b:   255,
		hex: "#ffffff",
	},
	{
		// black
		r:   0,
		g:   0,
		b:   0,
		hex: "#000000",
	},
	{
		// red
		r:   255,
		g:   0,
		b:   0,
		hex: "#ff0000",
	},
	{
		// blue
		r:   0,
		g:   0,
		b:   255,
		hex: "#0000ff",
	},
	{
		// pink
		r:   255,
		g:   0,
		b:   255,
		hex: "#ff00ff",
	},
	{
		// maroon
		r:   128,
		g:   0,
		b:   0,
		hex: "#800000",
	},
	{
		// linen
		r:   250,
		g:   240,
		b:   230,
		hex: "#faf0e6",
	},
}
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
		old:     "#ffff40",
		lighter: "#00ff4e",
		factor:  0.05,
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
	cmap := map[string]string{"deffacefg": "#000000", "deffacebg": "e2dfd9"}
	add := addColors(cmap)
	assert.Equal("#000000", add["fore2"])
	assert.Equal("#181818", add["back2"])

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
