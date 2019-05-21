package main

import (
	"fmt"
	"myanrichal/dot_product/src/service"
)

func main() {
	// var a = [][]int{{3, 2, 1}, {1, 2, 3}}
	// var b = [][]int{{0}, {2}, {4}}

	var a = [][]int{{1, 2}, {3, 4}}
	var b = [][]int{{4, 3}, {2, 1}}

	_, err := service.DotProduct(a, b)
	if err != nil {
		fmt.Println(err)
	}

	// util.SmartPrint(a)
	// util.SmartPrint(b)
	// util.SmartPrint(x)
}
