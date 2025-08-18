package rays_test

import (
	"raytracer-vibe/matrices"
	"raytracer-vibe/rays"
	"raytracer-vibe/tuples"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRay(t *testing.T) {
	// Scenario: Creating and querying a ray
	origin := tuples.Point(1, 2, 3)
	direction := tuples.Vector(4, 5, 6)
	r := rays.New(origin, direction)
	assert.Equal(t, origin, r.Origin)
	assert.Equal(t, direction, r.Direction)
}

func TestRayPosition(t *testing.T) {
	// Scenario: Computing a point from a distance
	r := rays.New(tuples.Point(2, 3, 4), tuples.Vector(1, 0, 0))
	assert.True(t, tuples.Point(2, 3, 4).Equals(r.Position(0)))
	assert.True(t, tuples.Point(3, 3, 4).Equals(r.Position(1)))
	assert.True(t, tuples.Point(1, 3, 4).Equals(r.Position(-1)))
	assert.True(t, tuples.Point(4.5, 3, 4).Equals(r.Position(2.5)))
}

func TestRayTransform(t *testing.T) {
	// Scenario: Translating a ray
	r := rays.New(tuples.Point(1, 2, 3), tuples.Vector(0, 1, 0))
	m := matrices.Translation(3, 4, 5)
	r2 := r.Transform(m)
	assert.True(t, tuples.Point(4, 6, 8).Equals(r2.Origin))
	assert.True(t, tuples.Vector(0, 1, 0).Equals(r2.Direction))
}

func TestRayScaling(t *testing.T) {
	// Scenario: Scaling a ray
	r := rays.New(tuples.Point(1, 2, 3), tuples.Vector(0, 1, 0))
	m := matrices.Scaling(2, 3, 4)
	r2 := r.Transform(m)
	assert.True(t, tuples.Point(2, 6, 12).Equals(r2.Origin))
	assert.True(t, tuples.Vector(0, 3, 0).Equals(r2.Direction))
}
