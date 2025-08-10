package tuples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTuples(t *testing.T) {
	a1 := Tuple{3, -2, 5, 1}
	a2 := Tuple{-2, 3, 1, 0}
	expected := Tuple{1, 1, 6, 1}
	result := Add(a1, a2)
	assert.True(t, Equal(expected, result))
}

func TestSubtractPoints(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)
	expected := Vector(-2, -4, -6)
	result := Subtract(p1, p2)
	assert.True(t, Equal(expected, result))
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p := Point(3, 2, 1)
	v := Vector(5, 6, 7)
	expected := Point(-2, -4, -6)
	result := Subtract(p, v)
	assert.True(t, Equal(expected, result))
}

func TestSubtractVectors(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)
	expected := Vector(-2, -4, -6)
	result := Subtract(v1, v2)
	assert.True(t, Equal(expected, result))
}

func TestNegate(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	expected := Tuple{-1, 2, -3, 4}
	result := Negate(a)
	assert.True(t, Equal(expected, result))
}

func TestMultiplyByScalar(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	expected := Tuple{3.5, -7, 10.5, -14}
	result := Multiply(a, 3.5)
	assert.True(t, Equal(expected, result))
}

func TestMultiplyByFraction(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	expected := Tuple{0.5, -1, 1.5, -2}
	result := Multiply(a, 0.5)
	assert.True(t, Equal(expected, result))
}

func TestDivideByScalar(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	expected := Tuple{0.5, -1, 1.5, -2}
	result := Divide(a, 2)
	assert.True(t, Equal(expected, result))
}

func TestMagnitude(t *testing.T) {
	v := Vector(1, 0, 0)
	assert.Equal(t, 1.0, Magnitude(v))

	v = Vector(0, 1, 0)
	assert.Equal(t, 1.0, Magnitude(v))

	v = Vector(0, 0, 1)
	assert.Equal(t, 1.0, Magnitude(v))

	v = Vector(1, 2, 3)
	assert.Equal(t, 3.7416573867739413, Magnitude(v))

	v = Vector(-1, -2, -3)
	assert.Equal(t, 3.7416573867739413, Magnitude(v))
}

func TestNormalize(t *testing.T) {
	v := Vector(4, 0, 0)
	expected := Vector(1, 0, 0)
	result := Normalize(v)
	assert.True(t, Equal(expected, result))

	v = Vector(1, 2, 3)
	expected = Vector(0.26726, 0.53452, 0.80178)
	result = Normalize(v)
	assert.True(t, Equal(expected, result))
}

func TestMagnitudeOfNormalizedVector(t *testing.T) {
	v := Vector(1, 2, 3)
	norm := Normalize(v)
	assert.InDelta(t, 1.0, Magnitude(norm), epsilon)
}

func TestDotProduct(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)
	assert.Equal(t, 20.0, Dot(a, b))
}

func TestCrossProduct(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)

	expected1 := Vector(-1, 2, -1)
	result1 := Cross(a, b)
	assert.True(t, Equal(expected1, result1))

	expected2 := Vector(1, -2, 1)
	result2 := Cross(b, a)
	assert.True(t, Equal(expected2, result2))
}
