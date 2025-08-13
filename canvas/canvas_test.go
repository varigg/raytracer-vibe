package canvas_test

import (
	"raytracer-vibe/tuples"
	"testing"

	"raytracer-vibe/canvas"

	"github.com/stretchr/testify/assert"
)

func TestNewCanvas(t *testing.T) {
	c := canvas.NewCanvas(10, 20)
	assert.Equal(t, 10, c.Width)
	assert.Equal(t, 20, c.Height)
	for _, p := range c.Pixels {
		assert.True(t, tuples.Equal(p.Tuple, tuples.NewColor(0, 0, 0).Tuple))
	}
}

func TestWritePixel(t *testing.T) {
	c := canvas.NewCanvas(10, 20)
	red := tuples.NewColor(1, 0, 0)
	c.WritePixel(2, 3, red)
	assert.True(t, tuples.Equal(red.Tuple, c.PixelAt(2, 3).Tuple))
}

func TestToPPMHeader(t *testing.T) {
	c := canvas.NewCanvas(5, 3)
	ppm := c.ToPPM()
	expected := "P3\n5 3\n255\n"
	assert.Equal(t, expected, ppm[:len(expected)])
}

func TestToPPMPixelData(t *testing.T) {
	c := canvas.NewCanvas(5, 3)
	c1 := tuples.NewColor(1.5, 0, 0)
	c2 := tuples.NewColor(0, 0.5, 0)
	c3 := tuples.NewColor(-0.5, 0, 1)
	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)
	ppm := c.ToPPM()
	expected := "P3\n5 3\n255\n255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 128 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 0 0 0 0 0 0 255\n"
	assert.Equal(t, expected, ppm)
}

func TestToPPMSplitLines(t *testing.T) {
	c := canvas.NewCanvas(10, 2)
	for i := range c.Pixels {
		c.Pixels[i] = tuples.NewColor(1, 0.8, 0.6)
	}
	ppm := c.ToPPM()
	expected := "P3\n10 2\n255\n255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n"
	assert.Equal(t, expected, ppm)
}

func TestToPPMEndsWithNewline(t *testing.T) {
	c := canvas.NewCanvas(5, 3)
	ppm := c.ToPPM()
	assert.Equal(t, "\n", ppm[len(ppm)-1:])
}
