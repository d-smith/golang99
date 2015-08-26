package slices

import (
	"errors"
)

type GenericSlice []interface{}

//Last returns the last element in a slice
func (s GenericSlice) Last() (interface{}, error) {
	length := len(s)
	if length == 0 {
		return -1, errors.New("Can't return last element from zero length slice")
	}

	return s[length-1], nil
}

//Penultimate returns the next to last element of a non-empty slice
func (s GenericSlice) Penultimate() (interface{}, error) {
	length := len(s)
	if length < 2 {
		return -1, errors.New("Can't return next to last element from slice of length less than two.")
	}

	return s[length-2], nil
}

//Kth returns the kth element of a slice, starting with 0. Yah, trivial stuff when using slices...
func (s GenericSlice) Kth(k int) (interface{}, error) {
	if k < 0 || k > len(s)-1 {
		return -1, errors.New("k index is out of bounds")
	}

	return s[k], nil
}

//Reverse reverses the elements of a slice.
func (s GenericSlice) Reverse() {
	n := len(s)
	if n == 0 {
		return
	}

	end := n - 1

	for i := 0; i < n/2; i++ {
		tmp := s[i]
		s[i] = s[end]
		s[end] = tmp
		end--
	}
}

//MakeReverse makes a copy of the slice, reverses the elements in the copy, and returns
//the reversed copy
func (s GenericSlice) MakeReverse() GenericSlice {
	c := make(GenericSlice, len(s))
	copy(c,s)
	c.Reverse()
	return c
}

//IsPalindrome returns true if the contents of a slice are the same forwards and backwards
func (s GenericSlice) IsPalindrome(f func(interface{},interface{}) bool) bool {
	r := s.MakeReverse()
	for i, v := range s {
		if !f(v,r[i]) {
			return false
		}
	}

	return true
}

//Flatten looks for GenericSlice elements in a GenericSlice, flattening out their elements such
//that the Generic slice produced by the function is a single level GenericSlice containing all
//embedded elements.
func (s GenericSlice) Flatten() GenericSlice {
	var ns GenericSlice
	for _, v := range s {
		switch v.(type) {
			default:
				ns = append(ns,v)
		case GenericSlice:
			fs := v.(GenericSlice).Flatten()
			ns = append(ns, fs...)
		}
	}

	return ns
}