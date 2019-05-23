package test

import (
	"myanrichal/dot_product/src/util"
	"testing"
)

func TestTranspose(t *testing.T) {

	b := [][]int{{0}, {2}, {4}}
	tb := [][]int{{0, 2, 4}}

	result, err := util.Transpose(b)
	if err != nil {
		t.Error(err)
	}

	assertEquals(t, "Transpose", Arguments{tb, result})
}
