package gomatrix

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	Columns uint `json:"columns"`
	Rows    uint `json:"rows"`

	Matrix [][]int `json:"matrix"`
}

var (
	ErrZeroValue    = errors.New("using zero values for the columns or rows of the matrix")
	ErrSumDiffOrder = errors.New("summing matrices of different order")
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
		Columns: cols,
		Rows:    rows,

		Matrix: matrix,
	}, nil
}

func FromJSON(data []byte) (*Matrix, error) {
	var matrix Matrix

	err := json.Unmarshal(data, &matrix)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling: %v", err)
	}

	return &matrix, err
}

// Generates a string from a matrix suitable for human perception
func (m *Matrix) String() string {
	var rows []string

	for _, row := range m.Matrix {
		var elements []string
		for _, num := range row {
			elements = append(elements, strconv.Itoa(num))
		}
		rows = append(rows, strings.Join(elements, " "))
	}

	return strings.Join(rows, "\n")
}

func Sum(matrices ...*Matrix) (*Matrix, error) {
	if !matricesAreSameOrder(matrices...) {
		return nil, ErrSumDiffOrder
	}

	rows := matrices[0].Rows
	cols := matrices[0].Columns

	matrixSum, _ := New(cols, rows)

	for _, matrix := range matrices {
		for i := uint(0); i < rows; i++ {
			for j := uint(0); j < cols; j++ {
				matrixSum.Matrix[i][j] += matrix.Matrix[i][j]
			}
		}
	}

	return matrixSum, nil
}

func matricesAreSameOrder(matrices ...*Matrix) bool {
	firstRow := matrices[0].Rows
	firstCol := matrices[0].Columns

	for i := 1; i < len(matrices); i++ {
		if firstRow != matrices[i].Rows || firstCol != matrices[i].Columns {
			return false
		}
	}

	return true
}
