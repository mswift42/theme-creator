package main

import (
	"testing"

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
	old     RGB
	lighter RGB
	factor  float64
}{
	{
		old:     RGB{122, 122, 122},
		lighter: RGB{148, 148, 148},
		factor:  0.20,
	},
	{
		old:     RGB{240, 240, 240},
		lighter: RGB{243, 243, 243},
		factor:  0.20,
	},
	{
		old:     RGB{0, 0, 0},
		lighter: RGB{51, 51, 51},
		factor:  0.20,
	},
}
var darkenedColors = []struct {
	old    RGB
	darker RGB
	factor float64
}{
	{
		old:    RGB{100, 100, 100},
		darker: RGB{80, 80, 80},
		factor: 0.80,
	},
	{
		old:    RGB{255, 255, 255},
		darker: RGB{204, 204, 204},
		factor: 0.80,
	},
	{
		old:    RGB{0, 0, 0},
		darker: RGB{0, 0, 0},
	},
}

func TestRgbToHex(t *testing.T) {
	assert := assert.New(t)
	for _, test := range convertColors {
		actual := rgbToHex(test.r, test.g, test.b)
		assert.Equal(actual, test.hex)
	}
}

func TestRGB(t *testing.T) {
	assert := assert.New(t)
	rgb1 := RGB{r: 255, g: 255, b: 255}
	assert.Equal(rgb1.r, 255)
	rgb2 := RGB{r: 0, g: 0, b: 0}
	assert.Equal(rgb2.r, 0)
}

func TestNewRGB(t *testing.T) {
	assert := assert.New(t)
	r1 := NewRGB(255, 255, 255)
	assert.Equal(r1.r, 255)
	r2 := NewRGB(255, 255, 0)
	assert.Equal(r2.b, 0)
}
func TestLighten(t *testing.T) {
	assert := assert.New(t)
	for _, i := range lightendedColors {
		actual := i.old.Lighten(i.factor)
		assert.Equal(actual, i.lighter)
	}
}
func TestDarken(t *testing.T) {
	assert := assert.New(t)
	for _, i := range darkenedColors {
		actual := i.old.Darken(i.factor)
		assert.Equal(actual, i.darker)
	}
}
