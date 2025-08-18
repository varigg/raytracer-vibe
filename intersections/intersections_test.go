package intersections_test

import (
	"testing"

	"raytracer-vibe/intersections"
	"raytracer-vibe/objects"

	"github.com/stretchr/testify/assert"
)

type mockObject struct{}

func TestIntersection(t *testing.T) {
	// Scenario: An intersection encapsulates t and object
	// Given o ← mockObject()
	// When i ← intersection(3.5, o)
	// Then i.t = 3.5
	// And i.object = o
	o := mockObject{}
	i := intersections.NewIntersection(3.5, o)
	assert.InEpsilon(t, 3.5, i.T, 0.00001)
	assert.Equal(t, o, i.Object)
}

func TestIntersections(t *testing.T) {
	// Scenario: Aggregating intersections
	// Given o ← mockObject()
	// And i1 ← intersection(1, o)
	// And i2 ← intersection(2, o)
	// When xs ← intersections(i1, i2)
	// Then xs.count = 2
	// And xs[0].t = 1
	// And xs[1].t = 2
	o := mockObject{}
	i1 := intersections.NewIntersection(1, o)
	i2 := intersections.NewIntersection(2, o)
	xs := intersections.NewIntersections(i1, i2)
	assert.Len(t, xs, 2)
	assert.InEpsilon(t, 1.0, xs[0].T, 0.00001)
	assert.InEpsilon(t, 2.0, xs[1].T, 0.00001)
}

func TestHit(t *testing.T) {
	o := mockObject{}

	t.Run("The hit, when all intersections have positive t", func(t *testing.T) {
		// Scenario: The hit, when all intersections have positive t
		// Given o ← mockObject()
		// And i1 ← intersection(1, o)
		// And i2 ← intersection(2, o)
		// And xs ← intersections(i2, i1)
		// When i ← hit(xs)
		// Then i = i1
		i1 := intersections.NewIntersection(1, o)
		i2 := intersections.NewIntersection(2, o)
		xs := intersections.NewIntersections(i2, i1)
		hit, found := xs.Hit()
		assert.True(t, found)
		assert.Equal(t, i1, hit)
	})

	t.Run("The hit, when some intersections have negative t", func(t *testing.T) {
		// Scenario: The hit, when some intersections have negative t
		// Given o ← mockObject()
		// And i1 ← intersection(-1, o)
		// And i2 ← intersection(1, o)
		// And xs ← intersections(i2, i1)
		// When i ← hit(xs)
		// Then i = i2
		i1 := intersections.NewIntersection(-1, o)
		i2 := intersections.NewIntersection(1, o)
		xs := intersections.NewIntersections(i2, i1)
		hit, found := xs.Hit()
		assert.True(t, found)
		assert.Equal(t, i2, hit)
	})

	t.Run("The hit, when all intersections have negative t", func(t *testing.T) {
		// Scenario: The hit, when all intersections have negative t
		// Given o ← mockObject()
		// And i1 ← intersection(-2, o)
		// And i2 ← intersection(-1, o)
		// And xs ← intersections(i2, i1)
		// When i ← hit(xs)
		// Then i is nothing
		i1 := intersections.NewIntersection(-2, o)
		i2 := intersections.NewIntersection(-1, o)
		xs := intersections.NewIntersections(i2, i1)
		_, found := xs.Hit()
		assert.False(t, found)
	})

	t.Run("The hit is always the lowest nonnegative intersection", func(t *testing.T) {
		// Scenario: The hit is always the lowest nonnegative intersection
		// Given o ← mockObject()
		// And i1 ← intersection(5, o)
		// And i2 ← intersection(7, o)
		// And i3 ← intersection(-3, o)
		// And i4 ← intersection(2, o)
		// And xs ← intersections(i1, i2, i3, i4)
		// When i ← hit(xs)
		// Then i = i4
		i1 := intersections.NewIntersection(5, o)
		i2 := intersections.NewIntersection(7, o)
		i3 := intersections.NewIntersection(-3, o)
		i4 := intersections.NewIntersection(2, o)
		xs := intersections.NewIntersections(i1, i2, i3, i4)
		hit, found := xs.Hit()
		assert.True(t, found)
		assert.Equal(t, i4, hit)
	})
}

var _ objects.Object = mockObject{}
