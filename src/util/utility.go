package util

import (
	"errors"
	"fmt"
)

//Print array
func SmartPrint(a [][]int) {
	for _, i := range a {
		fmt.Print("| ")
		for _, j := range i {
			fmt.Print(j, " ")
		}
		fmt.Println("|")
	}
	fmt.Print("\n")
}

//Verify the Matrix is complete
func VerifyMatrix(a [][]int) (error, int) {
	var b []int
	for _, h := range a {
		b = append(b, len(h))
	}
	for _, k := range b[0:] {
		if b[0] != k {
			//problem
			return errors.New("Matrix is not of consistent size"), 0
		} else {
			//good. consistent with true
		}
	}
	//return success and the second array length
	return nil, b[0]
}

//Verify the matricies can be multipled
func VerifyCompatiable(a [][]int, b [][]int) error {

	err, size_a := VerifyMatrix(a)
	if err != nil {
		return err
	}
	err, _ = VerifyMatrix(b)
	if err != nil {
		return err
	}

	if size_a != len(b) {
		err := errors.New("Sizes are not compatiable")
		return err
	}
	return nil
}

//transpose matrix
func Transpose(matrix [][]int) ([][]int, error) {
	err, width := VerifyMatrix(matrix)
	if err != nil {
		return nil, err
	}

	transposedMatrix := make([][]int, width)
	for i := range transposedMatrix {
		transposedMatrix[i] = make([]int, len(matrix))
		for j := range transposedMatrix[i] {
			transposedMatrix[i][j] = matrix[j][i]
		}
	}
	return transposedMatrix, nil
}
