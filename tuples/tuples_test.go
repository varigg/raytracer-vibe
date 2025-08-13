package tuples_test

import (
	"raytracer-vibe/tuples"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTuples(t *testing.T) {
	a1 := tuples.Tuple{X: 3, Y: -2, Z: 5, W: 1}
	a2 := tuples.Tuple{X: -2, Y: 3, Z: 1, W: 0}
	expected := tuples.Tuple{X: 1, Y: 1, Z: 6, W: 1}
	result := tuples.Add(a1, a2)
	assert.True(t, tuples.Equal(expected, result))
}

func TestSubtractPoints(t *testing.T) {
	p1 := tuples.Point(3, 2, 1)
	p2 := tuples.Point(5, 6, 7)
	expected := tuples.Vector(-2, -4, -6)
	result := tuples.Subtract(p1, p2)
	assert.True(t, tuples.Equal(expected, result))
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p := tuples.Point(3, 2, 1)
	v := tuples.Vector(5, 6, 7)
	expected := tuples.Point(-2, -4, -6)
	result := tuples.Subtract(p, v)
	assert.True(t, tuples.Equal(expected, result))
}

func TestSubtractVectors(t *testing.T) {
	v1 := tuples.Vector(3, 2, 1)
	v2 := tuples.Vector(5, 6, 7)
	expected := tuples.Vector(-2, -4, -6)
	result := tuples.Subtract(v1, v2)
	assert.True(t, tuples.Equal(expected, result))
}

func TestNegate(t *testing.T) {
	a := tuples.Tuple{X: 1, Y: -2, Z: 3, W: -4}
	expected := tuples.Tuple{X: -1, Y: 2, Z: -3, W: 4}
	result := tuples.Negate(a)
	assert.True(t, tuples.Equal(expected, result))
}

func TestMultiplyByScalar(t *testing.T) {
	a := tuples.Tuple{X: 1, Y: -2, Z: 3, W: -4}
	expected := tuples.Tuple{X: 3.5, Y: -7, Z: 10.5, W: -14}
	result := tuples.Multiply(a, 3.5)
	assert.True(t, tuples.Equal(expected, result))
}

func TestMultiplyByFraction(t *testing.T) {
	a := tuples.Tuple{X: 1, Y: -2, Z: 3, W: -4}
	expected := tuples.Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2}
	result := tuples.Multiply(a, 0.5)
	assert.True(t, tuples.Equal(expected, result))
}

func TestDivideByScalar(t *testing.T) {
	a := tuples.Tuple{X: 1, Y: -2, Z: 3, W: -4}
	expected := tuples.Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2}
	result := tuples.Divide(a, 2)
	assert.True(t, tuples.Equal(expected, result))
}

func TestMagnitude(t *testing.T) {
	v := tuples.Vector(1, 0, 0)
	assert.True(t, tuples.FloatEqual(1.0, tuples.Magnitude(v)))

	v = tuples.Vector(0, 1, 0)
	assert.True(t, tuples.FloatEqual(1.0, tuples.Magnitude(v)))

	v = tuples.Vector(0, 0, 1)
	assert.True(t, tuples.FloatEqual(1.0, tuples.Magnitude(v)))

	v = tuples.Vector(1, 2, 3)
	assert.True(t, tuples.FloatEqual(3.7416573867739413, tuples.Magnitude(v)))

	v = tuples.Vector(-1, -2, -3)
	assert.True(t, tuples.FloatEqual(3.7416573867739413, tuples.Magnitude(v)))
}

func TestNormalize(t *testing.T) {
	v := tuples.Vector(4, 0, 0)
	expected := tuples.Vector(1, 0, 0)
	result := tuples.Normalize(v)
	assert.True(t, tuples.Equal(expected, result))

	v = tuples.Vector(1, 2, 3)
	expected = tuples.Vector(0.26726, 0.53452, 0.80178)
	result = tuples.Normalize(v)
	assert.True(t, tuples.Equal(expected, result))
}

func TestMagnitudeOfNormalizedVector(t *testing.T) {
	v := tuples.Vector(1, 2, 3)
	norm := tuples.Normalize(v)
	assert.True(t, tuples.FloatEqual(1.0, tuples.Magnitude(norm)))
}

func TestDotProduct(t *testing.T) {
	a := tuples.Vector(1, 2, 3)
	b := tuples.Vector(2, 3, 4)
	assert.True(t, tuples.FloatEqual(20.0, tuples.Dot(a, b)))
}

func TestCrossProduct(t *testing.T) {
	a := tuples.Vector(1, 2, 3)
	b := tuples.Vector(2, 3, 4)

	expected1 := tuples.Vector(-1, 2, -1)
	result1 := tuples.Cross(a, b)
	assert.True(t, tuples.Equal(expected1, result1))

	expected2 := tuples.Vector(1, -2, 1)
	result2 := tuples.Cross(b, a)
	assert.True(t, tuples.Equal(expected2, result2))
}

func TestColor(t *testing.T) {
	// Scenario: Colors are (red, green, blue) tuples
	c := tuples.NewColor(-0.5, 0.4, 1.7)
	assert.True(t, tuples.FloatEqual(-0.5, c.Red()))
	assert.True(t, tuples.FloatEqual(0.4, c.Green()))
	assert.True(t, tuples.FloatEqual(1.7, c.Blue()))
}
