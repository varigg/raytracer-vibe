package rays_test

import (
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
