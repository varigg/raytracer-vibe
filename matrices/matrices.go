package matrices

import (
	"math"
	"raytracer-vibe/tuples"
)

const epsilon = 0.00001

type Matrix struct {
	rows, cols int
	data       [][]float64
}

func New(rows, cols int, data ...float64) Matrix {
	m := Matrix{
		rows: rows,
		cols: cols,
		data: make([][]float64, rows),
	}
	for i := range rows {
		m.data[i] = make([]float64, cols)
		for j := range cols {
			if len(data) > 0 {
				m.data[i][j] = data[i*cols+j]
			}
		}
	}
	return m
}

func Identity(size int) Matrix {
	m := New(size, size)
	for i := range size {
		m.data[i][i] = 1
	}
	return m
}

func (m Matrix) Get(row, col int) float64 {
	return m.data[row][col]
}

func (m Matrix) Equals(m2 Matrix) bool {
	if m.rows != m2.rows || m.cols != m2.cols {
		return false
	}
	for i := range m.rows {
		for j := range m.cols {
			if math.Abs(m.data[i][j]-m2.data[i][j]) > epsilon {
				return false
			}
		}
	}
	return true
}

func (m Matrix) Multiply(m2 Matrix) Matrix {
	newM := New(m.rows, m2.cols)
	for i := range m.rows {
		for j := range m2.cols {
			for k := range m.cols {
				newM.data[i][j] += m.data[i][k] * m2.data[k][j]
			}
		}
	}
	return newM
}

func (m Matrix) MultiplyTuple(t tuples.Tuple) tuples.Tuple {
	return tuples.New(
		m.data[0][0]*t.X+m.data[0][1]*t.Y+m.data[0][2]*t.Z+m.data[0][3]*t.W,
		m.data[1][0]*t.X+m.data[1][1]*t.Y+m.data[1][2]*t.Z+m.data[1][3]*t.W,
		m.data[2][0]*t.X+m.data[2][1]*t.Y+m.data[2][2]*t.Z+m.data[2][3]*t.W,
		m.data[3][0]*t.X+m.data[3][1]*t.Y+m.data[3][2]*t.Z+m.data[3][3]*t.W,
	)
}

func (m Matrix) Transpose() Matrix {
	t := New(m.cols, m.rows)
	for i := range m.rows {
		for j := range m.cols {
			t.data[j][i] = m.data[i][j]
		}
	}
	return t
}

func (m Matrix) Determinant() float64 {
	var det float64
	const two = 2
	if m.rows == two {
		det = m.data[0][0]*m.data[1][1] - m.data[0][1]*m.data[1][0]
	} else {
		for j := range m.cols {
			det += m.data[0][j] * m.Cofactor(0, j)
		}
	}
	return det
}

func (m Matrix) Submatrix(row, col int) Matrix {
	sub := New(m.rows-1, m.cols-1)
	for i := range m.rows {
		if i == row {
			continue
		}
		for j := range m.cols {
			if j == col {
				continue
			}
			subI, subJ := i, j
			if i > row {
				subI--
			}
			if j > col {
				subJ--
			}
			sub.data[subI][subJ] = m.data[i][j]
		}
	}
	return sub
}

func (m Matrix) Minor(row, col int) float64 {
	return m.Submatrix(row, col).Determinant()
}

func (m Matrix) Cofactor(row, col int) float64 {
	minor := m.Minor(row, col)
	if (row+col)%2 == 1 {
		return -minor
	}
	return minor
}

func (m Matrix) IsInvertible() bool {
	return m.Determinant() != 0
}

func (m Matrix) Inverse() Matrix {
	if !m.IsInvertible() {
		panic("matrix not invertible")
	}
	m2 := New(m.rows, m.cols)
	det := m.Determinant()
	for i := range m.rows {
		for j := range m.cols {
			c := m.Cofactor(i, j)
			m2.data[j][i] = c / det
		}
	}
	return m2
}

func Translation(x, y, z float64) Matrix {
	m := Identity(4)
	m.data[0][3] = x
	m.data[1][3] = y
	m.data[2][3] = z
	return m
}

func Scaling(x, y, z float64) Matrix {
	m := Identity(4)
	m.data[0][0] = x
	m.data[1][1] = y
	m.data[2][2] = z
	return m
}

func RotationX(radians float64) Matrix {
	m := Identity(4)
	m.data[1][1] = math.Cos(radians)
	m.data[1][2] = -math.Sin(radians)
	m.data[2][1] = math.Sin(radians)
	m.data[2][2] = math.Cos(radians)
	return m
}

func RotationY(radians float64) Matrix {
	m := Identity(4)
	m.data[0][0] = math.Cos(radians)
	m.data[0][2] = math.Sin(radians)
	m.data[2][0] = -math.Sin(radians)
	m.data[2][2] = math.Cos(radians)
	return m
}

func RotationZ(radians float64) Matrix {
	m := Identity(4)
	m.data[0][0] = math.Cos(radians)
	m.data[0][1] = -math.Sin(radians)
	m.data[1][0] = math.Sin(radians)
	m.data[1][1] = math.Cos(radians)
	return m
}