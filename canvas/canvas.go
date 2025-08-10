package canvas

import (
	"math"
	"raytracer-vibe/tuples"
	"strconv"
)

const (
	MaxLineLength = 70
	ColorScale    = 255
)

type Canvas struct {
	Width, Height int
	Pixels        []tuples.Tuple
}

func NewCanvas(width, height int) *Canvas {
	pixels := make([]tuples.Tuple, width*height)
	for i := range pixels {
		pixels[i] = tuples.Color(0, 0, 0)
	}
	return &Canvas{Width: width, Height: height, Pixels: pixels}
}

func (c *Canvas) WritePixel(x, y int, color tuples.Tuple) {
	c.Pixels[y*c.Width+x] = color
}

func (c *Canvas) PixelAt(x, y int) tuples.Tuple {
	return c.Pixels[y*c.Width+x]
}

func (c *Canvas) ToPPM() string {
	ppm := "P3\n" + strconv.Itoa(c.Width) + " " + strconv.Itoa(c.Height) + "\n255\n"
	line := ""
	for y := range c.Height {
		for x := range c.Width {

			pixel := c.PixelAt(x, y)
			r := scaleAndClamp(pixel.Red())
			g := scaleAndClamp(pixel.Green())
			b := scaleAndClamp(pixel.Blue())
			rStr := strconv.Itoa(r)
			gStr := strconv.Itoa(g)
			bStr := strconv.Itoa(b)
			if len(line)+len(rStr)+1 > MaxLineLength {
				ppm += line[:len(line)-1] + "\n"
				line = ""
			}
			line += rStr + " "
			if len(line)+len(gStr)+1 > MaxLineLength {
				ppm += line[:len(line)-1] + "\n"
				line = ""
			}
			line += gStr + " "
			if len(line)+len(bStr)+1 > MaxLineLength {
				ppm += line[:len(line)-1] + "\n"
				line = ""
			}
			line += bStr + " "
		}
		ppm += line[:len(line)-1] + "\n"
		line = ""
	}
	return ppm
}

func scaleAndClamp(c float64) int {
	scaled := math.Round(c * ColorScale)
	if scaled < 0 {
		return 0
	}
	if scaled > ColorScale {
		return ColorScale
	}
	return int(scaled)
}
