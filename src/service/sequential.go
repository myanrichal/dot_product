package service

import (
	"fmt"
	"myanrichal/dot_product/src/util"
)

func DotProduct(a [][]int, b [][]int) ([][]int, error) {

	err := util.VerifyCompatiable(a, b)
	if err != nil {
		return nil, err
	}
	//Verified

	tb, err := util.Transpose(b)
	if err != nil {
		return nil, err
	}

	fmt.Println("len: ", len(a), " : ", len(tb))

	product := make([][]int, len(a))
	for i := range product {
		product[i] = make([]int, len(tb))
	}

	var sum, row_sum int
	for hi, _ := range tb {
		for ii, i := range a {
			for ij, j := range i {
				for _, k := range tb {
					fmt.Println(j, " * ", k[ij], " = ", j*k[ij])
					sum = j * k[ij]
				}
				row_sum += sum
			}
			fmt.Println("break: ", row_sum)
			product[ii][hi] = row_sum
			row_sum = 0
		}
	}

	util.SmartPrint(product)

	return nil, nil
}
