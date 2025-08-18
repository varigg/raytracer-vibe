package spheres_test

import (
	"raytracer-vibe/matrices"
	"raytracer-vibe/rays"
	"raytracer-vibe/spheres"
	"raytracer-vibe/tuples"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayIntersectsSphereAtTwoPoints(t *testing.T) {
	// Scenario: A ray intersects a sphere at two points
	r := rays.New(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
	s := spheres.NewSphere()
	xs := s.Intersect(r)
	assert.Len(t, xs, 2)
	assert.InEpsilon(t, 4.0, xs[0].T, 0.00001)
	assert.InEpsilon(t, 6.0, xs[1].T, 0.00001)
}

func TestRayIntersectsSphereAtTangent(t *testing.T) {
	// Scenario: A ray intersects a sphere at a tangent
	r := rays.New(tuples.Point(0, 1, -5), tuples.Vector(0, 0, 1))
	s := spheres.NewSphere()
	xs := s.Intersect(r)
	assert.Len(t, xs, 2)
	assert.InEpsilon(t, 5.0, xs[0].T, 0.00001)
	assert.InEpsilon(t, 5.0, xs[1].T, 0.00001)
}

func TestRayMissesSphere(t *testing.T) {
	// Scenario: A ray misses a sphere
	r := rays.New(tuples.Point(0, 2, -5), tuples.Vector(0, 0, 1))
	s := spheres.NewSphere()
	xs := s.Intersect(r)
	assert.Empty(t, xs)
}

func TestRayOriginatesInsideSphere(t *testing.T) {
	// Scenario: A ray originates inside a sphere
	r := rays.New(tuples.Point(0, 0, 0), tuples.Vector(0, 0, 1))
	s := spheres.NewSphere()
	xs := s.Intersect(r)
	assert.Len(t, xs, 2)
	assert.InEpsilon(t, -1.0, xs[0].T, 0.00001)
	assert.InEpsilon(t, 1.0, xs[1].T, 0.00001)
}

func TestSphereBehindRay(t *testing.T) {
	// Scenario: A sphere is behind a ray
	r := rays.New(tuples.Point(0, 0, 5), tuples.Vector(0, 0, 1))
	s := spheres.NewSphere()
	xs := s.Intersect(r)
	assert.Len(t, xs, 2)
	assert.InEpsilon(t, -6.0, xs[0].T, 0.00001)
	assert.InEpsilon(t, -4.0, xs[1].T, 0.00001)
}

func TestIntersectSetsTheObjectOnTheIntersection(t *testing.T) {
	// Scenario: An intersect sets the object on the intersection
	r := rays.New(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
	s := spheres.NewSphere()
	xs := s.Intersect(r)
	assert.Len(t, xs, 2)
	assert.Equal(t, s, xs[0].Object)
	assert.Equal(t, s, xs[1].Object)
}

func TestSphereDefaultTransformation(t *testing.T) {
	// Scenario: A sphere's default transformation
	s := spheres.NewSphere()
	assert.True(t, s.Transform.Equals(matrices.Identity(4)))
}

func TestSphereSetTransform(t *testing.T) {
	// Scenario: Changing a sphere's transformation
	s := spheres.NewSphere()
	tr := matrices.Translation(2, 3, 4)
	s.SetTransform(tr)
	assert.True(t, s.Transform.Equals(tr))
}

func TestIntersectScaledSphereWithRay(t *testing.T) {
	// Scenario: Intersecting a scaled sphere with a ray
	r := rays.New(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
	s := spheres.NewSphere()
	s.SetTransform(matrices.Scaling(2, 2, 2))
	xs := s.Intersect(r)
	assert.Len(t, xs, 2)
	assert.InEpsilon(t, 3, xs[0].T, 0.00001)
	assert.InEpsilon(t, 7, xs[1].T, 0.00001)
}

func TestIntersectTranslatedSphereWithRay(t *testing.T) {
	// Scenario: Intersecting a translated sphere with a ray
	r := rays.New(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
	s := spheres.NewSphere()
	s.SetTransform(matrices.Translation(5, 0, 0))
	xs := s.Intersect(r)
	assert.Empty(t, xs)
}
