package emacsthemecreator

import "strconv"

const test = `'((mainbg . {{.bg1}}) (mainfg . {{ .fg1}}))`

// RGB struct, contains an uint8 each for
// red, green and blue.
type RGB struct {
	r, g, b int
}

// NewRGB - Constructor for RGB struct.
// takes 3 uint for red, green and blue and
// returns a new RGB struct
func NewRGB(r, g, b int) RGB {
	rgb := RGB{r: r, g: g, b: b}
	return rgb
}

// rgbToHex - convert 3 rgb color values
// to a hex string.
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

// Lighten - given an rgb value r, multiply each
// value of r by a given factor.
// Example: RGB{122,122,122}.Lighten(0.20) -> RGB{148,148,148}
func (r RGB) Lighten(factor float64) RGB {
	newrgbslice := []int{r.r, r.g, r.b}
	for i, j := range newrgbslice {
		lighter := j + int((factor * (255.0 - float64(j))))
		if lighter > 255 {
			lighter = 255
		}

		newrgbslice[i] = lighter
	}
	return NewRGB(newrgbslice[0], newrgbslice[1], newrgbslice[2])
}

// Darken - given an RGB value r, multiply each
// value of r by factor f, and return a new RGB
// with darker color.
func (r RGB) Darken(factor float64) RGB {
	newrgbslice := []int{r.r, r.g, r.b}
	for i, j := range newrgbslice {
		darker := int(factor * float64(j))
		newrgbslice[i] = darker
	}
	return NewRGB(newrgbslice[0], newrgbslice[1], newrgbslice[2])
}

// func main() {
//	temp := template.Must(template.New("test").Parse(test))
//	temp.Execute(os.Stdout, map[string]interface{}{"bg1": "\"#000000\"",
//		"fg1": "\"#ffffff\""})
//	fmt.Println()
//	fmt.Println(rgbToHex(255, 255, 255))
//	fmt.Println(rgbToHex(0, 0, 0))
//	fmt.Println(122 * 120 / 100)
// }
