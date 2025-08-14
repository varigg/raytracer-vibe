package matrices_test

import (
	"math"
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

// Scenario: Calculating the determinant of a 2x2 matrix.
// Given the following 2x2 matrix A:
// | 1 | 5 |
// | -3 | 2 |
// Then determinant(A) = 17.
func TestDeterminantOf2x2Matrix(t *testing.T) {
	m := matrices.New(2, 2, 1, 5, -3, 2)
	assert.InEpsilon(t, 17, m.Determinant(), epsilon)
}

// Scenario: A submatrix of a 3x3 matrix is a 2x2 matrix.
// Given the following 3x3 matrix A:
// | 1 | 5 | 0 |
// | -3 | 2 | 7 |
// | 0 | 6 | -3 |
// Then submatrix(A, 0, 2) is the following 2x2 matrix:
// | -3 | 2 |
// | 0 | 6 |.
func TestSubmatrixOf3x3Matrix(t *testing.T) {
	m := matrices.New(3, 3, 1, 5, 0, -3, 2, 7, 0, 6, -3)
	expected := matrices.New(2, 2, -3, 2, 0, 6)
	assert.True(t, m.Submatrix(0, 2).Equals(expected))
}

// Scenario: A submatrix of a 4x4 matrix is a 3x3 matrix.
// Given the following 4x4 matrix A:
// | -6 | 1 | 1 | 6 |
// | -8 | 5 | 8 | 6 |
// | -1 | 0 | 8 | 2 |
// | -7 | 1 | -1 | 1 |
// Then submatrix(A, 2, 1) is the following 3x3 matrix:
// | -6 | 1 | 6 |
// | -8 | 8 | 6 |
// | -7 | -1 | 1 |.
func TestSubmatrixOf4x4Matrix(t *testing.T) {
	m := matrices.New(4, 4, -6, 1, 1, 6, -8, 5, 8, 6, -1, 0, 8, 2, -7, 1, -1, 1)
	expected := matrices.New(3, 3, -6, 1, 6, -8, 8, 6, -7, -1, 1)
	assert.True(t, m.Submatrix(2, 1).Equals(expected))
}

// Scenario: Calculating a minor of a 3x3 matrix.
// Given the following 3x3 matrix A:
// | 3 | 5 | 0 |
// | 2 | -1 | -7 |
// | 6 | -1 | 5 |
// And B ← submatrix(A, 1, 0).
// Then determinant(B) = 25.
// And minor(A, 1, 0) = 25.
func TestMinorOf3x3Matrix(t *testing.T) {
	m := matrices.New(3, 3, 3, 5, 0, 2, -1, -7, 6, -1, 5)
	b := m.Submatrix(1, 0)
	assert.InEpsilon(t, 25, b.Determinant(), epsilon)
	assert.InEpsilon(t, 25, m.Minor(1, 0), epsilon)
}

// Scenario: Calculating a cofactor of a 3x3 matrix.
// Given the following 3x3 matrix A:
// | 3 | 5 | 0 |
// | 2 | -1 | -7 |
// | 6 | -1 | 5 |
// Then minor(A, 0, 0) = -12.
// And cofactor(A, 0, 0) = -12.
// And minor(A, 1, 0) = 25.
// And cofactor(A, 1, 0) = -25.
func TestCofactorOf3x3Matrix(t *testing.T) {
	m := matrices.New(3, 3, 3, 5, 0, 2, -1, -7, 6, -1, 5)
	assert.InEpsilon(t, -12, m.Minor(0, 0), epsilon)
	assert.InEpsilon(t, -12, m.Cofactor(0, 0), epsilon)
	assert.InEpsilon(t, 25, m.Minor(1, 0), epsilon)
	assert.InEpsilon(t, -25, m.Cofactor(1, 0), epsilon)
}

// Scenario: Calculating the determinant of a 3x3 matrix.
// Given the following 3x3 matrix A:
// | 1 | 2 | 6 |
// | -5 | 8 | -4 |
// | 2 | 6 | 4 |
// Then cofactor(A, 0, 0) = 56.
// And cofactor(A, 0, 1) = 12.
// And cofactor(A, 0, 2) = -46.
// And determinant(A) = -196.
func TestDeterminantOf3x3Matrix(t *testing.T) {
	m := matrices.New(3, 3, 1, 2, 6, -5, 8, -4, 2, 6, 4)
	assert.InEpsilon(t, 56, m.Cofactor(0, 0), epsilon)
	assert.InEpsilon(t, 12, m.Cofactor(0, 1), epsilon)
	assert.InEpsilon(t, -46, m.Cofactor(0, 2), epsilon)
	assert.InEpsilon(t, -196, m.Determinant(), epsilon)
}

// Scenario: Calculating the determinant of a 4x4 matrix.
// Given the following 4x4 matrix A:
// | -2 | -8 | 3 | 5 |
// | -3 | 1 | 7 | 3 |
// | 1 | 2 | -9 | 6 |
// | -6 | 7 | 7 | -9 |
// Then cofactor(A, 0, 0) = 690.
// And cofactor(A, 0, 1) = 447.
// And cofactor(A, 0, 2) = 210.
// And cofactor(A, 0, 3) = 51.
// And determinant(A) = -4071.
func TestDeterminantOf4x4Matrix(t *testing.T) {
	m := matrices.New(4, 4, -2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9)
	assert.InEpsilon(t, 690, m.Cofactor(0, 0), epsilon)
	assert.InEpsilon(t, 447, m.Cofactor(0, 1), epsilon)
	assert.InEpsilon(t, 210, m.Cofactor(0, 2), epsilon)
	assert.InEpsilon(t, 51, m.Cofactor(0, 3), epsilon)
	assert.InEpsilon(t, -4071, m.Determinant(), epsilon)
}

// Scenario: Testing an invertible matrix for invertibility.
// Given the following 4x4 matrix A:
// | 6 | 4 | 4 | 4 |
// | 5 | 5 | 7 | 6 |
// | 4 | -9 | 3 | -7 |
// | 9 | 1 | 7 | -6 |
// Then determinant(A) = -2120.
// And A is invertible.
func TestInvertibleMatrix(t *testing.T) {
	m := matrices.New(4, 4, 6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 9, 1, 7, -6)
	assert.InEpsilon(t, -2120, m.Determinant(), epsilon)
	assert.True(t, m.IsInvertible())
}

// Scenario: Testing a noninvertible matrix for invertibility.
// Given the following 4x4 matrix A:
// | -4 | 2 | -2 | -3 |
// | 9 | 6 | 2 | 6 |
// | 0 | -5 | 1 | -5 |
// | 0 | 0 | 0 | 0 |
// Then determinant(A) = 0.
// And A is not invertible.
func TestNonInvertibleMatrix(t *testing.T) {
	m := matrices.New(4, 4, -4, 2, -2, -3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0)
	assert.Equal(t, 0.0, m.Determinant()) //nolint:testifylint // InEpsilon doesn't work with 0.
	assert.False(t, m.IsInvertible())
}

// Scenario: Calculating the inverse of a matrix.
// Given the following 4x4 matrix A:
// | -5 | 2 | 6 | -8 |
// | 1 | -5 | 1 | 8 |
// | 7 | 7 | -6 | -7 |
// | 1 | -3 | 7 | 4 |
// And B ← inverse(A).
// Then determinant(A) = 532.
// And cofactor(A, 2, 3) = -160.
// And B[3,2] = -160/532.
// And cofactor(A, 3, 2) = 105.
// And B[2,3] = 105/532.
// And B is the following 4x4 matrix:
// | 0.21805 | 0.45113 | 0.24060 | -0.04511 |
// | -0.80827 | -1.45677 | -0.44361 | 0.52068 |
// | -0.07895 | -0.22368 | -0.05263 | 0.19737 |
// | -0.52256 | -0.81391 | -0.30075 | 0.30639 |.
func TestInverseOfMatrix(t *testing.T) {
	m := matrices.New(4, 4, -5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4)
	b := m.Inverse()
	expected := matrices.New(4, 4,
		0.21805, 0.45113, 0.24060, -0.04511,
		-0.80827, -1.45677, -0.44361, 0.52068,
		-0.07895, -0.22368, -0.05263, 0.19737,
		-0.52256, -0.81391, -0.30075, 0.30639,
	)
	assert.InEpsilon(t, 532, m.Determinant(), epsilon)
	assert.InEpsilon(t, -160, m.Cofactor(2, 3), epsilon)
	assert.InEpsilon(t, -160.0/532.0, b.Get(3, 2), epsilon)
	assert.InEpsilon(t, 105, m.Cofactor(3, 2), epsilon)
	assert.InEpsilon(t, 105.0/532.0, b.Get(2, 3), epsilon)
	assert.True(t, b.Equals(expected))
}

// Scenario: Calculating the inverse of another matrix.
// Given the following 4x4 matrix A:
// | 8 | -5 | 9 | 2 |
// | 7 | 5 | 6 | 1 |
// | -6 | 0 | 9 | 6 |
// | -3 | 0 | -9 | -4 |
// Then inverse(A) is the following 4x4 matrix:
// | -0.15385 | -0.15385 | -0.28205 | -0.53846 |
// | -0.07692 | 0.12308 | 0.02564 | 0.03077 |
// | 0.35897 | 0.35897 | 0.43590 | 0.92308 |
// | -0.69231 | -0.69231 | -0.76923 | -1.92308 |.
func TestInverseOfAnotherMatrix(t *testing.T) {
	m := matrices.New(4, 4, 8, -5, 9, 2, 7, 5, 6, 1, -6, 0, 9, 6, -3, 0, -9, -4)
	expected := matrices.New(4, 4,
		-0.15385, -0.15385, -0.28205, -0.53846,
		-0.07692, 0.12308, 0.02564, 0.03077,
		0.35897, 0.35897, 0.43590, 0.92308,
		-0.69231, -0.69231, -0.76923, -1.92308,
	)
	assert.True(t, m.Inverse().Equals(expected))
}

// Scenario: Calculating the inverse of a third matrix.
// Given the following 4x4 matrix A:
// | 9 | 3 | 0 | 9 |
// | -5 | -2 | -6 | -3 |
// | -4 | 9 | 6 | 4 |
// | -7 | 6 | 6 | 2 |
// Then inverse(A) is the following 4x4 matrix:
// | -0.04074 | -0.07778 | 0.14444 | -0.22222 |
// | -0.07778 | 0.03333 | 0.36667 | -0.33333 |
// | -0.02901 | -0.14630 | -0.10926 | 0.12963 |
// | 0.17778 | 0.06667 | -0.26667 | 0.33333 |.
func TestInverseOfAThirdMatrix(t *testing.T) {
	m := matrices.New(4, 4, 9, 3, 0, 9, -5, -2, -6, -3, -4, 9, 6, 4, -7, 6, 6, 2)
	expected := matrices.New(4, 4,
		-0.04074, -0.07778, 0.14444, -0.22222,
		-0.07778, 0.03333, 0.36667, -0.33333,
		-0.02901, -0.14630, -0.10926, 0.12963,
		0.17778, 0.06667, -0.26667, 0.33333,
	)
	assert.True(t, m.Inverse().Equals(expected))
}

// Scenario: Multiplying a product by its inverse.
// Given the following 4x4 matrix A:
// | 3 | -9 | 7 | 3 |
// | 3 | -8 | 2 | -9 |
// | -4 | 4 | 4 | 1 |
// | -6 | 5 | -1 | 1 |
// And the following 4x4 matrix B:
// | 8 | 2 | 2 | 2 |
// | 3 | -1 | 7 | 0 |
// | 7 | 0 | 5 | 4 |
// | 6 | -2 | 0 | 5 |
// And C ← A * B.
// Then C * inverse(B) = A.
func TestMultiplyProductByInverse(t *testing.T) {
	a := matrices.New(4, 4, 3, -9, 7, 3, 3, -8, 2, -9, -4, 4, 4, 1, -6, 5, -1, 1)
	b := matrices.New(4, 4, 8, 2, 2, 2, 3, -1, 7, 0, 7, 0, 5, 4, 6, -2, 0, 5)
	c := a.Multiply(b)
	assert.True(t, c.Multiply(b.Inverse()).Equals(a))
}

// Scenario: Multiplying by a translation matrix
// Given transform ← translation(5, -3, 2)
// And p ← point(-3, 4, 5)
// Then transform * p = point(2, 1, 7).
func TestTranslation(t *testing.T) {
	transform := matrices.Translation(5, -3, 2)
	p := tuples.Point(-3, 4, 5)
	expected := tuples.Point(2, 1, 7)
	assert.True(t, transform.MultiplyTuple(p).Equals(expected))
}

// Scenario: Multiplying by the inverse of a translation matrix
// Given transform ← translation(5, -3, 2)
// And inv ← inverse(transform)
// And p ← point(-3, 4, 5)
// Then inv * p = point(-8, 7, 3).
func TestTranslationInverse(t *testing.T) {
	transform := matrices.Translation(5, -3, 2)
	inv := transform.Inverse()
	p := tuples.Point(-3, 4, 5)
	expected := tuples.Point(-8, 7, 3)
	assert.True(t, inv.MultiplyTuple(p).Equals(expected))
}

// Scenario: Translation does not affect vectors.
// Given transform ← translation(5, -3, 2)
// And v ← vector(-3, 4, 5)
// Then transform * v = v.
func TestTranslationVector(t *testing.T) {
	transform := matrices.Translation(5, -3, 2)
	v := tuples.Vector(-3, 4, 5)
	assert.True(t, transform.MultiplyTuple(v).Equals(v))
}

// Scenario: A scaling matrix applied to a point
// Given transform ← scaling(2, 3, 4)
// And p ← point(-4, 6, 8)
// Then transform * p = point(-8, 18, 32).
func TestScalingPoint(t *testing.T) {
	transform := matrices.Scaling(2, 3, 4)
	p := tuples.Point(-4, 6, 8)
	expected := tuples.Point(-8, 18, 32)
	assert.True(t, transform.MultiplyTuple(p).Equals(expected))
}

// Scenario: A scaling matrix applied to a vector
// Given transform ← scaling(2, 3, 4)
// And v ← vector(-4, 6, 8)
// Then transform * v = vector(-8, 18, 32).
func TestScalingVector(t *testing.T) {
	transform := matrices.Scaling(2, 3, 4)
	v := tuples.Vector(-4, 6, 8)
	expected := tuples.Vector(-8, 18, 32)
	assert.True(t, transform.MultiplyTuple(v).Equals(expected))
}

// Scenario: Multiplying by the inverse of a scaling matrix
// Given transform ← scaling(2, 3, 4)
// And inv ← inverse(transform)
// And v ← vector(-4, 6, 8)
// Then inv * v = vector(-2, 2, 2).
func TestScalingInverse(t *testing.T) {
	transform := matrices.Scaling(2, 3, 4)
	inv := transform.Inverse()
	v := tuples.Vector(-4, 6, 8)
	expected := tuples.Vector(-2, 2, 2)
	assert.True(t, inv.MultiplyTuple(v).Equals(expected))
}

// Scenario: Reflection is scaling by a negative value
// Given transform ← scaling(-1, 1, 1)
// And p ← point(2, 3, 4)
// Then transform * p = point(-2, 3, 4).
func TestReflection(t *testing.T) {
	transform := matrices.Scaling(-1, 1, 1)
	p := tuples.Point(2, 3, 4)
	expected := tuples.Point(-2, 3, 4)
	assert.True(t, transform.MultiplyTuple(p).Equals(expected))
}

// Scenario: Rotating a point around the x axis
// Given p ← point(0, 1, 0)
// And half_quarter ← rotation_x(π / 4)
// And full_quarter ← rotation_x(π / 2)
// Then half_quarter * p = point(0, √2/2, √2/2)
// And full_quarter * p = point(0, 0, 1).
func TestRotationX(t *testing.T) {
	p := tuples.Point(0, 1, 0)
	halfQuarter := matrices.RotationX(math.Pi / 4)
	fullQuarter := matrices.RotationX(math.Pi / 2)
	expectedHalf := tuples.Point(0, math.Sqrt2/2, math.Sqrt2/2)
	expectedFull := tuples.Point(0, 0, 1)
	assert.True(t, halfQuarter.MultiplyTuple(p).Equals(expectedHalf))
	assert.True(t, fullQuarter.MultiplyTuple(p).Equals(expectedFull))
}

// Scenario: The inverse of an x-rotation rotates in the opposite direction
// Given p ← point(0, 1, 0)
// And half_quarter ← rotation_x(π / 4)
// And inv ← inverse(half_quarter)
// Then inv * p = point(0, √2/2, -√2/2).
func TestRotationXInverse(t *testing.T) {
	p := tuples.Point(0, 1, 0)
	halfQuarter := matrices.RotationX(math.Pi / 4)
	inv := halfQuarter.Inverse()
	expected := tuples.Point(0, math.Sqrt2/2, -math.Sqrt2/2)
	assert.True(t, inv.MultiplyTuple(p).Equals(expected))
}

// Scenario: Rotating a point around the y axis
// Given p ← point(0, 0, 1)
// And half_quarter ← rotation_y(π / 4)
// And full_quarter ← rotation_y(π / 2)
// Then half_quarter * p = point(√2/2, 0, √2/2)
// And full_quarter * p = point(1, 0, 0).
func TestRotationY(t *testing.T) {
	p := tuples.Point(0, 0, 1)
	halfQuarter := matrices.RotationY(math.Pi / 4)
	fullQuarter := matrices.RotationY(math.Pi / 2)
	expectedHalf := tuples.Point(math.Sqrt2/2, 0, math.Sqrt2/2)
	expectedFull := tuples.Point(1, 0, 0)
	assert.True(t, halfQuarter.MultiplyTuple(p).Equals(expectedHalf))
	assert.True(t, fullQuarter.MultiplyTuple(p).Equals(expectedFull))
}

// Scenario: Rotating a point around the z axis
// Given p ← point(0, 1, 0)
// And half_quarter ← rotation_z(π / 4)
// And full_quarter ← rotation_z(π / 2)
// Then half_quarter * p = point(-√2/2, √2/2, 0)
// And full_quarter * p = point(-1, 0, 0).
func TestRotationZ(t *testing.T) {
	p := tuples.Point(0, 1, 0)
	halfQuarter := matrices.RotationZ(math.Pi / 4)
	fullQuarter := matrices.RotationZ(math.Pi / 2)
	expectedHalf := tuples.Point(-math.Sqrt2/2, math.Sqrt2/2, 0)
	expectedFull := tuples.Point(-1, 0, 0)
	assert.True(t, halfQuarter.MultiplyTuple(p).Equals(expectedHalf))
	assert.True(t, fullQuarter.MultiplyTuple(p).Equals(expectedFull))
}

// Scenario: A shearing transformation moves x in proportion to y.
// Given transform ← shearing(1, 0, 0, 0, 0, 0)
// And p ← point(2, 3, 4)
// Then transform * p = point(5, 3, 4).
func TestShearingXinProportionToY(t *testing.T) {
	transform := matrices.Shearing(1, 0, 0, 0, 0, 0)
	p := tuples.Point(2, 3, 4)
	expected := tuples.Point(5, 3, 4)
	assert.True(t, transform.MultiplyTuple(p).Equals(expected))
}

// Scenario: A shearing transformation moves x in proportion to z.
// Given transform ← shearing(0, 1, 0, 0, 0, 0)
// And p ← point(2, 3, 4)
// Then transform * p = point(6, 3, 4).
func TestShearingXinProportionToZ(t *testing.T) {
	transform := matrices.Shearing(0, 1, 0, 0, 0, 0)
	p := tuples.Point(2, 3, 4)
	expected := tuples.Point(6, 3, 4)
	assert.True(t, transform.MultiplyTuple(p).Equals(expected))
}

// Scenario: A shearing transformation moves y in proportion to x.
// Given transform ← shearing(0, 0, 1, 0, 0, 0)
// And p ← point(2, 3, 4)
// Then transform * p = point(2, 5, 4).
func TestShearingYinProportionToX(t *testing.T) {
	transform := matrices.Shearing(0, 0, 1, 0, 0, 0)
	p := tuples.Point(2, 3, 4)
	expected := tuples.Point(2, 5, 4)
	assert.True(t, transform.MultiplyTuple(p).Equals(expected))
}

// Scenario: A shearing transformation moves y in proportion to z.
// Given transform ← shearing(0, 0, 0, 1, 0, 0)
// And p ← point(2, 3, 4)
// Then transform * p = point(2, 7, 4).
func TestShearingYinProportionToZ(t *testing.T) {
	transform := matrices.Shearing(0, 0, 0, 1, 0, 0)
	p := tuples.Point(2, 3, 4)
	expected := tuples.Point(2, 7, 4)
	assert.True(t, transform.MultiplyTuple(p).Equals(expected))
}

// Scenario: A shearing transformation moves z in proportion to x.
// Given transform ← shearing(0, 0, 0, 0, 1, 0)
// And p ← point(2, 3, 4)
// Then transform * p = point(2, 3, 6).
func TestShearingZinProportionToX(t *testing.T) {
	transform := matrices.Shearing(0, 0, 0, 0, 1, 0)
	p := tuples.Point(2, 3, 4)
	expected := tuples.Point(2, 3, 6)
	assert.True(t, transform.MultiplyTuple(p).Equals(expected))
}

// Scenario: A shearing transformation moves z in proportion to y.
// Given transform ← shearing(0, 0, 0, 0, 0, 1)
// And p ← point(2, 3, 4)
// Then transform * p = point(2, 3, 7).
func TestShearingZinProportionToY(t *testing.T) {
	transform := matrices.Shearing(0, 0, 0, 0, 0, 1)
	p := tuples.Point(2, 3, 4)
	expected := tuples.Point(2, 3, 7)
	assert.True(t, transform.MultiplyTuple(p).Equals(expected))
}

// Scenario: Individual transformations are applied in sequence
// Given p ← point(1, 0, 1)
// And A ← rotation_x(π / 2)
// And B ← scaling(5, 5, 5)
// And C ← translation(10, 5, 7)
// # apply rotation first
// When p2 ← A * p
// Then p2 = point(1, -1, 0)
// # then apply scaling
// When p3 ← B * p2
// Then p3 = point(5, -5, 0)
// # then apply translation
// When p4 ← C * p3
// Then p4 = point(15, 0, 7).
func TestIndividualTransformations(t *testing.T) {
	p := tuples.Point(1, 0, 1)
	a := matrices.RotationX(math.Pi / 2)
	b := matrices.Scaling(5, 5, 5)
	c := matrices.Translation(10, 5, 7)

	p2 := a.MultiplyTuple(p)
	assert.True(t, p2.Equals(tuples.Point(1, -1, 0)))

	p3 := b.MultiplyTuple(p2)
	assert.True(t, p3.Equals(tuples.Point(5, -5, 0)))

	p4 := c.MultiplyTuple(p3)
	assert.True(t, p4.Equals(tuples.Point(15, 0, 7)))
}

// Scenario: Chained transformations must be applied in reverse order
// Given p ← point(1, 0, 1)
// And A ← rotation_x(π / 2)
// And B ← scaling(5, 5, 5)
// And C ← translation(10, 5, 7)
// When T ← C * B * A
// Then T * p = point(15, 0, 7).
func TestChainedTransformations(t *testing.T) {
	p := tuples.Point(1, 0, 1)
	a := matrices.RotationX(math.Pi / 2)
	b := matrices.Scaling(5, 5, 5)
	c := matrices.Translation(10, 5, 7)

	transform := c.Multiply(b).Multiply(a)
	assert.True(t, transform.MultiplyTuple(p).Equals(tuples.Point(15, 0, 7)))
}
