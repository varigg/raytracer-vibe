package main

import (
	"fmt"
	"os"

	"raytracer-vibe/canvas"
	"raytracer-vibe/rays"
	"raytracer-vibe/spheres"
	"raytracer-vibe/tuples"
)

func main() {
	rayOrigin := tuples.Point(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 100

	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2 // nolint: mnd

	c := canvas.NewCanvas(canvasPixels, canvasPixels)
	color := tuples.NewColor(1, 0, 0)
	shape := spheres.NewSphere()

	for y := range canvasPixels {
		worldY := half - pixelSize*float64(y)
		for x := range canvasPixels {
			worldX := -half + pixelSize*float64(x)
			position := tuples.Point(worldX, worldY, wallZ)
			r := rays.New(rayOrigin, tuples.Normalize(position.Subtract(rayOrigin)))
			xs := shape.Intersect(r)
			_, found := xs.Hit()
			if found {
				c.WritePixel(x, y, color)
			}
		}
	}

	ppm := c.ToPPM()
	file, err := os.Create("silhouette.ppm")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(ppm)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
