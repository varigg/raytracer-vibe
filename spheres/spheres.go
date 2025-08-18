// nolint: mnd  // these are basic math operations
package spheres

import (
	"math"
	"raytracer-vibe/intersections"
	"raytracer-vibe/rays"
	"raytracer-vibe/tuples"
)

type Sphere struct {
}

func NewSphere() Sphere {
	return Sphere{}
}

func (s Sphere) Intersect(r rays.Ray) intersections.Intersections {
	sphereToRay := r.Origin.Subtract(tuples.Point(0, 0, 0))
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return intersections.NewIntersections()
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return intersections.NewIntersections(
		intersections.NewIntersection(t1, s),
		intersections.NewIntersection(t2, s),
	)
}
