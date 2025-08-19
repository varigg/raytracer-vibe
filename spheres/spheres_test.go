package spheres_test

import (
	"math"
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

func TestNormalOnSphereAtPointOnXAxis(t *testing.T) {
	// Scenario: The normal on a sphere at a point on the x axis
	s := spheres.NewSphere()
	n := s.NormalAt(tuples.Point(1, 0, 0))
	assert.True(t, n.Equals(tuples.Vector(1, 0, 0)))
}

func TestNormalOnSphereAtPointOnYAxis(t *testing.T) {
	// Scenario: The normal on a sphere at a point on the y axis
	s := spheres.NewSphere()
	n := s.NormalAt(tuples.Point(0, 1, 0))
	assert.True(t, n.Equals(tuples.Vector(0, 1, 0)))
}

func TestNormalOnSphereAtPointOnZAxis(t *testing.T) {
	// Scenario: The normal on a sphere at a point on the z axis
	s := spheres.NewSphere()
	n := s.NormalAt(tuples.Point(0, 0, 1))
	assert.True(t, n.Equals(tuples.Vector(0, 0, 1)))
}

func TestNormalOnSphereAtNonaxialPoint(t *testing.T) {
	// Scenario: The normal on a sphere at a nonaxial point
	s := spheres.NewSphere()
	val := math.Sqrt(3) / 3
	n := s.NormalAt(tuples.Point(val, val, val))
	assert.True(t, n.Equals(tuples.Vector(val, val, val)))
}

func TestNormalIsNormalizedVector(t *testing.T) {
	// Scenario: The normal is a normalized vector
	s := spheres.NewSphere()
	val := math.Sqrt(3) / 3
	n := s.NormalAt(tuples.Point(val, val, val))
	assert.True(t, n.Equals(tuples.Normalize(n)))
}

func TestComputingNormalOnTranslatedSphere(t *testing.T) {
	// Scenario: Computing the normal on a translated sphere
	s := spheres.NewSphere()
	s.SetTransform(matrices.Translation(0, 1, 0))
	n := s.NormalAt(tuples.Point(0, 1.70711, -0.70711))
	assert.True(t, n.Equals(tuples.Vector(0, 0.70711, -0.70711)))
}

func TestComputingNormalOnTransformedSphere(t *testing.T) {
	// Scenario: Computing the normal on a transformed sphere
	s := spheres.NewSphere()
	m := matrices.Scaling(1, 0.5, 1).Multiply(matrices.RotationZ(math.Pi / 5))
	s.SetTransform(m)
	val := math.Sqrt(2) / 2
	n := s.NormalAt(tuples.Point(0, val, -val))
	assert.True(t, n.Equals(tuples.Vector(0, 0.97014, -0.24254)))
}
