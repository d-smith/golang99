package slices

import (
	"errors"
)

//To do - can we make these functions methods with an interface{} receiver?

//Last returns the last element of a non-empty slice, or an error if given an empty slice.
func Last(s []interface{}) (interface{}, error) {
	length := len(s)
	if length == 0 {
		return -1, errors.New("Can't return last element from zero length slice")
	}

	return s[length-1], nil
}

//Penultimate returns the next to last element of a non-empty slice
func Penultimate(s []interface{}) (interface{}, error) {
	length := len(s); if length < 2 {
		return -1, errors.New("Can't return next to last element from slice of length less than two.")
	}

	return s[length - 2],nil
}

//Kth returns the kth element of a slice, starting with 0. Yah, trivial stuff when using slices...
func Kth(k int, s []interface{})(interface{}, error) {
	if k < 0 || k > len(s) - 1 {
		return -1, errors.New("k index is out of bounds")
	}

	return s[k],nil
}