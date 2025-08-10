package main

import (
	"fmt"
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

func Tick(env Environment, proj Projectile) Projectile {
	position := tuples.Add(proj.Position, proj.Velocity)
	velocity := tuples.Add(tuples.Add(proj.Velocity, env.Gravity), env.Wind)
	return Projectile{Position: position, Velocity: velocity}
}

func main() {
	p := Projectile{
		Position: tuples.Point(0, 1, 0),
		Velocity: tuples.Normalize(tuples.Vector(1, 1, 0)),
	}

	e := Environment{
		Gravity: tuples.Vector(0, -0.1, 0),
		Wind:    tuples.Vector(-0.01, 0, 0),
	}

	ticks := 0
	for p.Position.Y > 0 {
		p = Tick(e, p)
		ticks++
		fmt.Printf("Tick %d: Position: %v\n", ticks, p.Position)
	}

	fmt.Printf("Projectile hit the ground after %d ticks.\n", ticks)
}

