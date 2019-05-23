package test

import (
	"errors"
	"myanrichal/dot_product/src/util"
	"testing"
)

type Arguments struct {
	a [][]int
	b [][]int
}

func assertEquals(t *testing.T, name string, data Arguments) {

	//check size
	width, err := compareSize(t, name, data)
	if err != nil {
		t.FailNow()
	}

	//check values
	for i := 0; i < len(data.a); i++ {
		for j := 0; j < width; j++ {
			if !(data.a[i][j] == data.b[i][j]) {
				t.Log(name, "has failed on assertEquals")
				t.FailNow()
			}
		}
	}
}

func compareSize(t *testing.T, name string, data Arguments) (int, error) {

	err, widthA := util.VerifyMatrix(data.a)
	if err != nil {
		t.Log("Failed on verify A")
		return 0, errors.New("Failed test")
	}

	err, widthB := util.VerifyMatrix(data.b)
	if err != nil {
		t.Log("Failed on verify B")
		return 0, errors.New("Failed test")
	}

	if !((len(data.a) == len(data.b)) && (widthA == widthB)) {
		t.Log("Failed on compareSize")
		return 0, errors.New("Failed test")
	}
	return widthA, err
}
