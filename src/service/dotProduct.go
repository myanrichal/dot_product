package service

import (
	"myanrichal/dot_product/src/util"
)

func DotProduct(a [][]int, b [][]int) ([][]int, error) {

	err := util.VerifyCompatiable(a, b)
	if err != nil {
		return nil, err
	}
	//Verified

	//transpose for easier calculation
	tb, err := util.Transpose(b)
	if err != nil {
		return nil, err
	}

	//make new dynamic product array
	product := make([][]int, len(a))
	for i := range product {
		product[i] = make([]int, len(tb))
	}

	//perform calculation
	for tbIndex, tbRow := range tb {
		for aIndex, aRow := range a {
			for rowIndex := 0; rowIndex < len(aRow); rowIndex++ {
				product[aIndex][tbIndex] += aRow[rowIndex] * tbRow[rowIndex]
			}
		}
	}
	return product, nil
}
