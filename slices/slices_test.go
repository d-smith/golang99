package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLast(t *testing.T) {
	t.Log("given a slice of integers")
	s := GenericSlice{1, 2, 3, 4, 5}

	t.Log("when last is called, the last element of the slice is returned")
	last, err := s.Last()

	assert.Nil(t, err)
	assert.Equal(t, 5, last)

	a, _ := s.Last()
	assert.Equal(t, 10, a.(int)+a.(int))
}

func TestGenericLast(t *testing.T) {
	s := GenericSlice{1, 2, 3, 4, 5}
	last, err := s.Last()

	assert.Nil(t, err)
	assert.Equal(t, 5, last)
}

func TestLastErrCheck(t *testing.T) {
	t.Log("Given an empty slice")
	var s GenericSlice

	t.Log("when passed to last, an error is returned")
	_, err := s.Last()
	assert.NotNil(t, err)
}

func TestPenultimate(t *testing.T) {
	t.Log("Given a slice of 2 or more integers")
	s := GenericSlice{1, 2, 3, 4, 5}

	t.Log("when Penultimate is called, the next to last element is returned")
	n, err := s.Penultimate()
	assert.Nil(t, err)
	assert.Equal(t, 4, n.(int))
}

func TestKth(t *testing.T) {
	s := GenericSlice{1, 2, 3, 4, 5}

	var v interface{}
	var err error

	v, err = s.Kth(-1)
	assert.NotNil(t, err)

	v, err = s.Kth(5)
	assert.NotNil(t, err)

	v, err = s.Kth(0)
	assert.Nil(t, err)
	assert.Equal(t, 1, v.(int))

	v, err = s.Kth(4)
	assert.Nil(t, err)
	assert.Equal(t, 5, v.(int))

	s = append(s, 6)
	v, err = s.Kth(5)
	assert.Nil(t, err)
	assert.Equal(t, 6, v.(int))
}

func TestReverse(t *testing.T) {
	t.Log("Given an empty slice we can Reverse it without error")
		s := GenericSlice{}
		s.Reverse()

	t.Log("A single element slice can be reversed")
	s = append(s,1)
	s.Reverse()
	assert.Equal(t, 1, len(s))
	assert.Equal(t, 1, s[0])

	t.Log("A slice with an even number of elements can be reversed")
	e := GenericSlice{1,2,3,4}
	e.Reverse()
	assert.Equal(t, e, GenericSlice{4,3,2,1} )

	t.Log("A slice with an odd number of elements can be reversed")
	o := GenericSlice{1,2,3,4,5}
	o.Reverse()
	assert.Equal(t, o, GenericSlice{5,4,3,2,1} )
}
