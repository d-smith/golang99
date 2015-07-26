package slices

import (
	"errors"
)

//Last returns the last element of a non-empty slice, or an error if given an empty slice.
func Last(s []int) (int, error) {
	length := len(s)
	if length == 0 {
		return -1, errors.New("Can't return last element from zero length slice")
	}

	return s[length-1], nil
}

//Penultimate returns the next to last element of a non-empty slice
func Penultimate(s []int) (int, error) {
	length := len(s); if length < 2 {
		return -1, errors.New("Can't return next to last element from slice of length less than two.")
	}

	return s[length - 2],nil
}