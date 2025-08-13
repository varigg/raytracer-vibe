package matrices_test

import (
	"raytracer-vibe/matrices"
	"raytracer-vibe/tuples"
	"testing"

	"github.com/stretchr/testify/assert"
)

const epsilon = 0.00001

// Scenario: Creating and inspecting a 4x4 matrix.
// Given the following 4x4 matrix M:
// | 1 | 2 | 3 | 4 |
// | 5.5 | 6.5 | 7.5 | 8.5 |
// | 9 | 10 | 11 | 12 |
// | 13.5 | 14.5 | 15.5 | 16.5 |
// Then M[0,0] = 1.
// And M[0,3] = 4.
// And M[1,0] = 5.5.
// And M[1,2] = 7.5.
// And M[2,2] = 11.
// And M[3,0] = 13.5.
// And M[3,2] = 15.5.
func TestCreatingAndInspecting4x4Matrix(t *testing.T) {
	m := matrices.New(4, 4, 1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5)
	assert.InEpsilon(t, 1.0, m.Get(0, 0), epsilon)
	assert.InEpsilon(t, 4.0, m.Get(0, 3), epsilon)
	assert.InEpsilon(t, 5.5, m.Get(1, 0), epsilon)
	assert.InEpsilon(t, 7.5, m.Get(1, 2), epsilon)
	assert.InEpsilon(t, 11.0, m.Get(2, 2), epsilon)
	assert.InEpsilon(t, 13.5, m.Get(3, 0), epsilon)
	assert.InEpsilon(t, 15.5, m.Get(3, 2), epsilon)
}

// Scenario: A 2x2 matrix ought to be representable.
// Given the following 2x2 matrix M:
// | -3 | 5 |
// | 1 | -2 |
// Then M[0,0] = -3.
// And M[0,1] = 5.
// And M[1,0] = 1.
// And M[1,1] = -2.
func TestCreatingAndInspecting2x2Matrix(t *testing.T) {
	m := matrices.New(2, 2, -3, 5, 1, -2)
	assert.InEpsilon(t, -3.0, m.Get(0, 0), epsilon)
	assert.InEpsilon(t, 5.0, m.Get(0, 1), epsilon)
	assert.InEpsilon(t, 1.0, m.Get(1, 0), epsilon)
	assert.InEpsilon(t, -2.0, m.Get(1, 1), epsilon)
}

// Scenario: A 3x3 matrix ought to be representable.
// Given the following 3x3 matrix M:
// | -3 | 5 | 0 |
// | 1 | -2 | -7 |
// | 0 | 1 | 1 |
// Then M[0,0] = -3.
// And M[1,1] = -2.
// And M[2,2] = 1.
func TestCreatingAndInspecting3x3Matrix(t *testing.T) {
	m := matrices.New(3, 3, -3, 5, 0, 1, -2, -7, 0, 1, 1)
	assert.InEpsilon(t, -3.0, m.Get(0, 0), epsilon)
	assert.InEpsilon(t, -2.0, m.Get(1, 1), epsilon)
	assert.InEpsilon(t, 1.0, m.Get(2, 2), epsilon)
}

// Scenario: Matrix Equality with identical matrices.
// Given the following matrix A:
// | 1 | 2 | 3 | 4 |
// | 5 | 6 | 7 | 8 |
// | 9 | 8 | 7 | 6 |
// | 5 | 4 | 3 | 2 |
// And the following matrix B:
// | 1 | 2 | 3 | 4 |
// | 5 | 6 | 7 | 8 |
// | 9 | 8 | 7 | 6 |
// | 5 | 4 | 3 | 2 |
// Then A = B.
func TestMatrixEqualityWithIdenticalMatrices(t *testing.T) {
	m1 := matrices.New(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
	m2 := matrices.New(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
	assert.True(t, m1.Equals(m2))
}

// Scenario: Matrix Equality with different matrices.
// Given the following matrix A:
// | 1 | 2 | 3 | 4 |
// | 5 | 6 | 7 | 8 |
// | 9 | 8 | 7 | 6 |
// | 5 | 4 | 3 | 2 |
// And the following matrix B:
// | 2 | 3 | 4 | 5 |
// | 6 | 7 | 8 | 9 |
// | 8 | 7 | 6 | 5 |
// | 4 | 3 | 2 | 1 |
// Then A != B.
func TestMatrixEqualityWithDifferentMatrices(t *testing.T) {
	m1 := matrices.New(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
	m2 := matrices.New(4, 4, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	assert.False(t, m1.Equals(m2))
}

// Scenario: Multiplying two matrices.
// Given the following matrix A:
// | 1 | 2 | 3 | 4 |
// | 5 | 6 | 7 | 8 |
// | 9 | 8 | 7 | 6 |
// | 5 | 4 | 3 | 2 |
// And the following matrix B:
// | -2 | 1 | 2 | 3 |
// | 3 | 2 | 1 | -1 |
// | 4 | 3 | 6 | 5 |
// | 1 | 2 | 7 | 8 |
// Then A * B is the following 4x4 matrix:
// | 20 | 22 | 50 | 48 |
// | 44 | 54 | 114 | 108 |
// | 40 | 58 | 110 | 102 |
// | 16 | 26 | 46 | 42 |.
func TestMatrixMultiplication(t *testing.T) {
	m1 := matrices.New(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
	m2 := matrices.New(4, 4, -2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8)
	expected := matrices.New(4, 4, 20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42)
	assert.True(t, m1.Multiply(m2).Equals(expected))
}

// Scenario: Multiplying a matrix by a tuple.
// Given the following matrix A:
// | 1 | 2 | 3 | 4 |
// | 2 | 4 | 4 | 2 |
// | 8 | 6 | 4 | 1 |
// | 0 | 0 | 0 | 1 |
// And the following tuple b:
// (1, 2, 3, 1)
// Then A * b is the following tuple:
// (18, 24, 33, 1).
func TestMatrixTupleMultiplication(t *testing.T) {
	m := matrices.New(4, 4, 1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1)
	b := tuples.New(1, 2, 3, 1)
	expected := tuples.New(18, 24, 33, 1)
	assert.True(t, m.MultiplyTuple(b).Equals(expected))
}

// Scenario: Multiplying by the identity matrix.
// Given the following matrix A:
// | 0 | 1 | 2 | 4 |
// | 1 | 2 | 4 | 8 |
// | 2 | 4 | 8 | 16 |
// | 4 | 8 | 16 | 32 |
// And the identity matrix I.
// Then A * I = A.
func TestMultiplyingByIdentityMatrix(t *testing.T) {
	m := matrices.New(4, 4, 0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32)
	i := matrices.Identity(4)
	assert.True(t, m.Multiply(i).Equals(m))
}

// Scenario: Transposing a matrix.
// Given the following matrix A:
// | 0 | 9 | 3 | 0 |
// | 9 | 8 | 0 | 8 |
// | 1 | 8 | 5 | 3 |
// | 0 | 0 | 5 | 8 |
// Then transpose(A) is the following matrix:
// | 0 | 9 | 1 | 0 |
// | 9 | 8 | 8 | 0 |
// | 3 | 0 | 5 | 5 |
// | 0 | 8 | 3 | 8 |.
func TestTranspose(t *testing.T) {
	m := matrices.New(4, 4, 0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8)
	expected := matrices.New(4, 4, 0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8)
	assert.True(t, m.Transpose().Equals(expected))
}

// Scenario: Transposing the identity matrix.
// Given the identity matrix I.
// Then transpose(I) = I.
func TestTransposeIdentity(t *testing.T) {
	i := matrices.Identity(4)
	assert.True(t, i.Transpose().Equals(i))
}
