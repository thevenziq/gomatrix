package gomatrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	columns uint
	rows    uint

	matrix [][]int
}

var (
	ErrZeroValue = errors.New("using zero values for the columns or rows of the matrix")
)

// Creates a matrix and fills it with zeros
func New(cols, rows uint) (*Matrix, error) {
	if cols <= 0 || rows <= 0 {
		return nil, ErrZeroValue
	}

	matrix := make([][]int, rows)
	row := make([]int, cols)

	for i := range matrix {
		matrix[i] = row
	}

	return &Matrix{
		columns: cols,
		rows:    rows,

		matrix: matrix,
	}, nil
}

// Generates a string from a matrix suitable for human perception
func (m *Matrix) String() string {
	var rows []string

	for _, row := range m.matrix {
		var elements []string
		for _, num := range row {
			elements = append(elements, strconv.Itoa(num))
		}
		rows = append(rows, strings.Join(elements, " "))
	}

	return "\n" + strings.Join(rows, "\n")
}
