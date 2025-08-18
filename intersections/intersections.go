package intersections

import "raytracer-vibe/objects"

type Intersection struct {
	T      float64
	Object objects.Object
}

func NewIntersection(t float64, object objects.Object) Intersection {
	return Intersection{
		T:      t,
		Object: object,
	}
}

type Intersections []Intersection

func NewIntersections(intersections ...Intersection) Intersections {
	return intersections
}

func (xs Intersections) Hit() (Intersection, bool) {
	hit := Intersection{}
	found := false

	for _, i := range xs {
		if i.T >= 0 {
			if !found || i.T < hit.T {
				hit = i
				found = true
			}
		}
	}

	return hit, found
}
