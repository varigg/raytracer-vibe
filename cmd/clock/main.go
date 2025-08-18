package main

import (
	"fmt"
	"math"
	"os"
	"raytracer-vibe/canvas"
	"raytracer-vibe/matrices"
	"raytracer-vibe/tuples"
)

const (
	CanvasWidth  = 500
	CanvasHeight = 500
	Scale        = 150
	Translate    = 250
	RotationStep = 6
)

func main() {
	c := canvas.NewCanvas(CanvasWidth, CanvasHeight)
	white := tuples.NewColor(1, 1, 1)

	// Create a transformation matrix that scales and translates the points
	scale := matrices.Scaling(Scale, Scale, 0)
	translate := matrices.Translation(Translate, Translate, 0)

	// Draw 12 points in a circle
	for i := range 12 {
		point := tuples.Point(0, 1, 0)
		rotation := matrices.RotationZ(float64(i) * math.Pi / RotationStep)
		transform := translate.Multiply(scale).Multiply(rotation)
		transformedPoint := transform.MultiplyTuple(point)
		c.WritePixel(int(transformedPoint.X), int(transformedPoint.Y), white)
	}

	// Write the canvas to a PPM file
	ppm := c.ToPPM()
	err := os.WriteFile("clock.ppm", []byte(ppm), 0600)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
