package rays

import "raytracer-vibe/tuples"

type Ray struct {
	Origin, Direction tuples.Tuple
}

func New(origin, direction tuples.Tuple) Ray {
	return Ray{Origin: origin, Direction: direction}
}

func (r Ray) Position(t float64) tuples.Tuple {
	return r.Origin.Add(r.Direction.Multiply(t))
}
