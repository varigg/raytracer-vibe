// nolint: mnd  // these are basic math operations
package spheres

import (
	"math"
	"raytracer-vibe/intersections"
	"raytracer-vibe/matrices"
	"raytracer-vibe/rays"
	"raytracer-vibe/tuples"
)

type Sphere struct {
	Transform matrices.Matrix
}

func NewSphere() *Sphere {
	return &Sphere{
		Transform: matrices.Identity(4),
	}
}

func (s *Sphere) SetTransform(m matrices.Matrix) {
	s.Transform = m
}

func (s *Sphere) Intersect(r rays.Ray) intersections.Intersections {
	ray2 := r.Transform(s.Transform.Inverse())

	sphereToRay := ray2.Origin.Subtract(tuples.Point(0, 0, 0))
	a := ray2.Direction.Dot(ray2.Direction)
	b := 2 * ray2.Direction.Dot(sphereToRay)
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

func (s *Sphere) NormalAt(worldPoint tuples.Tuple) tuples.Tuple {
	objectPoint := s.Transform.Inverse().MultiplyTuple(worldPoint)
	objectNormal := objectPoint.Subtract(tuples.Point(0, 0, 0))
	worldNormal := s.Transform.Inverse().Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0
	return tuples.Normalize(worldNormal)
}
