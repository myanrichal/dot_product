package main

import (
	"fmt"
	"myanrichal/dot_product/src/service"
	"myanrichal/dot_product/src/util"
	"time"
)

func main() {

	a := [][]int{{3, 2, 1}, {1, 2, 3}}
	b := [][]int{{0}, {2}, {4}}

	start := time.Now()
	dp, err := service.DotProduct(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("elapsed: ", time.Since(start))
	util.SmartPrint(dp)
}
