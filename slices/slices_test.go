package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"reflect"
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
	s = append(s, 1)
	s.Reverse()
	assert.Equal(t, 1, len(s))
	assert.Equal(t, 1, s[0])

	t.Log("A slice with an even number of elements can be reversed")
	e := GenericSlice{1, 2, 3, 4}
	e.Reverse()
	assert.Equal(t, e, GenericSlice{4, 3, 2, 1})

	t.Log("A slice with an odd number of elements can be reversed")
	o := GenericSlice{1, 2, 3, 4, 5}
	o.Reverse()
	assert.Equal(t, o, GenericSlice{5, 4, 3, 2, 1})

	x := o.MakeReverse()
	assert.NotEqual(t, x, o)
	o.Reverse()
	assert.Equal(t, x, o)
}

func equalInts(a interface{}, b interface{}) bool {
	x := a.(int)
	y := b.(int)
	return x == y
}

func TestIsPalindrome(t *testing.T) {

	x := GenericSlice{1, 2, 3, 2, 1}
	assert.True(t, x.IsPalindrome(equalInts))

	y := GenericSlice{1, 2, 3, 4, 5}
	assert.False(t, y.IsPalindrome(equalInts))

	assert.True(t, GenericSlice{}.IsPalindrome(equalInts))
	assert.True(t, GenericSlice{1}.IsPalindrome(equalInts))

}

func TestFlatten(t *testing.T) {
	x := GenericSlice{GenericSlice{1, 1}, 2, GenericSlice{3, GenericSlice{5, 8}}}
	flattened := x.Flatten()
	assert.Equal(t, len(flattened), 6)
	assert.Equal(t, flattened[0], 1)
	assert.Equal(t, flattened[1], 1)
	assert.Equal(t, flattened[2], 2)
	assert.Equal(t, flattened[3], 3)
	assert.Equal(t, flattened[4], 5)
	assert.Equal(t, flattened[5], 8)
}

func TestCompress(t *testing.T) {
	a := GenericSlice{}
	compressed := a.Compress(equalInts)
	assert.Equal(t, 0, len(a))

	b := GenericSlice{1}
	compressed = b.Compress(equalInts)
	assert.Equal(t, 1, len(b))
	assert.Equal(t, 1, b[0])

	x := GenericSlice{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	compressed = x.Compress(equalInts)
	assert.Equal(t, 4, len(compressed))
	assert.Equal(t, 1, compressed[0])
	assert.Equal(t, 2, compressed[1])
	assert.Equal(t, 3, compressed[2])
	assert.Equal(t, 4, compressed[3])
}

func TestCompact(t *testing.T) {
	x := GenericSlice{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	packed := x.Pack(equalInts)
	if assert.Equal(t, 4, len(packed)) {

		s1 := packed[0].(GenericSlice)
		assert.Equal(t, 1, len(s1))
		assert.Equal(t, 1, s1[0])

		s2 := packed[1].(GenericSlice)
		assert.Equal(t, 2, len(s2))
		for _, v := range s2 {
			assert.Equal(t, 2, v)
		}

		s3 := packed[2].(GenericSlice)
		assert.Equal(t, 3, len(s3))
		for _, v := range s3 {
			assert.Equal(t, 3, v)
		}

		s4 := packed[3].(GenericSlice)
		assert.Equal(t, 4, len(s4))
		for _, v := range s4 {
			assert.Equal(t, 4, v)
		}
	}
}

func TestEncode(t *testing.T) {
	x := GenericSlice{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	encoded := x.Encode(equalInts)
	assert.Equal(t, 4, len(encoded))
	for i := 0; i < 4; i++ {
		s := encoded[i].(GenericSlice)
		assert.Equal(t, 2, len(s))
		assert.Equal(t, i+1, s[0])
		assert.Equal(t, i+1, s[1])
	}
}

func TestModifiedEncode(t *testing.T) {
	x := GenericSlice{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	encoded := x.ModifiedEncode(equalInts)
	assert.Equal(t, 4, len(encoded))

	assert.Equal(t, 1, encoded[0].(int))

	for i := 1; i < 4; i++ {
		s := encoded[i].(GenericSlice)
		assert.Equal(t, 2, len(s))
		assert.Equal(t, i+1, s[0])
		assert.Equal(t, i+1, s[1])
	}
}

func TestDecode(t *testing.T) {
	x := GenericSlice{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	encoded := x.Encode(equalInts)
	decoded := encoded.Decode()
	assert.True(t, reflect.DeepEqual(x, decoded))
}

func TestEncodeDirect(t *testing.T) {
	x := GenericSlice{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	encoded := x.EncodeDirect(equalInts)
	assert.Equal(t, 4, len(encoded))
	for i := 0; i < 4; i++ {
		s := encoded[i].(GenericSlice)
		assert.Equal(t, 2, len(s))
		assert.Equal(t, i+1, s[0])
		assert.Equal(t, i+1, s[1])
	}
}

func TestDuplicate(t *testing.T) {
	x := GenericSlice{1,2,3}
	y := x.Duplicate()
	assert.True(t, reflect.DeepEqual(y, GenericSlice{1,1,2,2,3,3}))
}