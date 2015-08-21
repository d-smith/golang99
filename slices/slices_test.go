package slices

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLast(t *testing.T) {
	t.Log("given a slice of integers")
	s := []interface{}{1, 2, 3, 4, 5}

	t.Log("when last is called, the last element of the slice is returned")
	last, err := Last(s)

	assert.Nil(t,err)
	assert.Equal(t,5,last)

	a,_ := Last(s)
	assert.Equal(t, 10, a.(int) + a.(int))
}

func TestLastErrCheck(t *testing.T) {
	t.Log("Given an empty slice")
	var s []interface{}

	t.Log("when passed to last, an error is returned")
	_,err := Last(s)
	assert.NotNil(t,err)
}

func TestPenultimate(t *testing.T) {
	t.Log("Given a slice of 2 or more integers")
	s := []interface{}{1,2,3,4,5}

	t.Log("when Penultimate is called, the next to last element is returned")
	n, err := Penultimate(s)
	assert.Nil(t,err)
	assert.Equal(t,4,n.(int))
}

func TestKth(t *testing.T) {
	s := []interface{}{1,2,3,4,5}

	var v interface{}
	var err error

	v,err = Kth(-1, s)
	assert.NotNil(t, err)

	v,err = Kth(300, s)
	assert.NotNil(t, err)

	v,err = Kth(0, s)
	assert.Nil(t,err)
	assert.Equal(t, 1, v.(int))

	v,err = Kth(4, s)
	assert.Nil(t,err)
	assert.Equal(t, 5, v.(int))
}
