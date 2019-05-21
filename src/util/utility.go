package util

import (
	"errors"
	"fmt"
)

func SmartPrint(a [][]int) {
	for _, i := range a {
		fmt.Print("|")
		for _, j := range i {
			fmt.Print(j)
		}
		fmt.Println("|")
	}
	fmt.Print("\n")
}

func verifyMatrix(a [][]int) (error, int) {
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

func VerifyCompatiable(a [][]int, b [][]int) error {

	err, size_a := verifyMatrix(a)
	if err != nil {
		return err
	}
	err, _ = verifyMatrix(b)
	if err != nil {
		return err
	}

	// fmt.Println("size_a: ", size_a)
	// fmt.Println("len: ", len(b))

	if size_a != len(b) {
		err := errors.New("Sizes are not compatiable")
		return err
	}
	return nil
}

func Transpose(matrix [][]int) ([][]int, error) {
	err, width := verifyMatrix(matrix)
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
