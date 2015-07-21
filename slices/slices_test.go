package slices

import (
	"testing"
)

func TestLast(t *testing.T) {
	t.Log("given a slice of integers")
	s := []int{1, 2, 3, 4, 5}

	t.Log("when last is called, the last element of the slice is returned")
	last, err := Last(s)

	if err != nil {
		t.Log("Error returned by last: ", err.Error())
		t.FailNow()
	}

	if last != 5 {
		t.Failed()
	}
}
