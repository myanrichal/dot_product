package test

import (
	"myanrichal/dot_product/src/service"
	"testing"
)

func TestDotProduct(t *testing.T) {

	a := [][]int{{3, 2, 1}, {1, 2, 3}}
	b := [][]int{{0}, {2}, {4}}
	c := [][]int{{8}, {16}}

	dp, err := service.DotProduct(a, b)
	if err != nil {
		t.Log(err)
		t.Errorf("Failed on DotProduct")
	}

	assertEquals(t, "Dot Product", Arguments{dp, c})

}
