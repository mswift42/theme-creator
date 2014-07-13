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

func TestRgbToHex(t *testing.T) {
	assert := assert.New(t)
	for _, test := range convertColors {
		actual := rgbToHex(test.r, test.g, test.b)
		assert.Equal(actual, test.hex)
	}
}
