package matrices

import (
	"raytracer-vibe/tuples"

	"github.com/google/go-cmp/cmp"
)

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
	return cmp.Equal(m.data, m2.data)
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
