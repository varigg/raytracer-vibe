package main

import (
	"fmt"
	"math"
	"os"

	"raytracer-vibe/canvas"
	"raytracer-vibe/tuples"
)

type Projectile struct {
	Position tuples.Tuple
	Velocity tuples.Tuple
}

type Environment struct {
	Gravity tuples.Tuple
	Wind    tuples.Tuple
}

const (
	CanvasWidth          = 900
	CanvasHeight         = 550
	VelocityMultiplier   = 11.25
	ProjectileUpVelocity = 1.8
)

func Tick(env Environment, proj Projectile) Projectile {
	position := tuples.Add(proj.Position, proj.Velocity)
	velocity := tuples.Add(tuples.Add(proj.Velocity, env.Gravity), env.Wind)
	return Projectile{Position: position, Velocity: velocity}
}

func main() {
	start := tuples.Point(0, 1, 0)
	velocity := tuples.Multiply(tuples.Normalize(tuples.Vector(1, ProjectileUpVelocity, 0)), VelocityMultiplier)
	p := Projectile{Position: start, Velocity: velocity}

	gravity := tuples.Vector(0, -0.1, 0)
	wind := tuples.Vector(-0.01, 0, 0)
	e := Environment{Gravity: gravity, Wind: wind}

	c := canvas.NewCanvas(CanvasWidth, CanvasHeight)
	red := tuples.NewColor(1, 0, 0)

	for p.Position.Y > 0 {
		x := int(math.Round(p.Position.X))
		y := int(math.Round(p.Position.Y))

		if x >= 0 && x < c.Width && y >= 0 && y < c.Height {
			c.WritePixel(x, c.Height-y, red)
		}

		p = Tick(e, p)
	}

	ppm := c.ToPPM()
	err := os.WriteFile("projectile.ppm", []byte(ppm), 0600)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
