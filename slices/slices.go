package slices

import (
	"errors"
)

//TODO - can we enhance the type to make clients assert the types of the content the
//GS is meant to contain?
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
	copy(c, s)
	c.Reverse()
	return c
}

//IsPalindrome returns true if the contents of a slice are the same forwards and backwards
func (s GenericSlice) IsPalindrome(f func(interface{}, interface{}) bool) bool {
	r := s.MakeReverse()
	for i, v := range s {
		if !f(v, r[i]) {
			return false
		}
	}

	return true
}

//Flatten looks for GenericSlice elements in a GenericSlice, flattening out their elements such
//that the Generic slice produced by the function is a single level GenericSlice containing all
//embedded elements. For example flattening GenericSlice{GenericSlice{1,1},2,GenericSlice{3, GenericSlice{5,8}}}
//produces GenericSlice{1,1,2,3,5,8}
func (s GenericSlice) Flatten() GenericSlice {
	var ns GenericSlice
	for _, v := range s {
		switch v.(type) {
		default:
			ns = append(ns, v)
		case GenericSlice:
			fs := v.(GenericSlice).Flatten()
			ns = append(ns, fs...)
		}
	}

	return ns
}

//Compress takes a slice and returns a new slice with consecutive duplicate elements remove
func (s GenericSlice) Compress(equal func(interface{}, interface{}) bool) GenericSlice {
	var ns GenericSlice
	if len(s) == 0 {
		return ns
	}

	ns = append(ns, s[0])
	curVal := s[0]
	for i := 1; i < len(s); i++ {
		if !equal(s[i], curVal) {
			curVal = s[i]
			ns = append(ns, curVal)
		}
	}
	return ns
}

//Pack returns a new slice with consecutive duplicates packed into their own slice
func (s GenericSlice) Pack(equal func(interface{}, interface{}) bool) GenericSlice {
	var ns GenericSlice
	if len(s) == 0 {
		return ns
	}

	curVal := s[0]
	ss := GenericSlice{curVal}
	ns = append(ns, ss)

	for i := 1; i < len(s); i++ {
		if equal(s[i], curVal) {
			ss = append(ss, curVal)
			ns[len(ns)-1] = ss
		} else {
			curVal = s[i]
			ss = GenericSlice{curVal}
			ns = append(ns, ss)
		}
	}

	return ns
}

//Encode performs run length encoding on slice contents
func (s GenericSlice) Encode(equal func(interface{}, interface{}) bool) GenericSlice {
	ns := GenericSlice{}
	packed := s.Pack(equal)
	for _, v := range packed {
		ps := GenericSlice{}
		ps = append(ps, len(v.(GenericSlice)))
		ps = append(ps, v.(GenericSlice)[0])
		ns = append(ns, ps)

	}

	return ns
}

//Modified encode performs run length encoding, but copies single elements directly
//into the list
func (s GenericSlice) ModifiedEncode(equal func(interface{}, interface{}) bool) GenericSlice {
	ns := s.Encode(equal)
	for i, v := range ns {
		if v.(GenericSlice)[0] == 1 {
			ns[i] = v.(GenericSlice)[1]
		}
	}

	return ns
}

//Decode decodes a runlength encoded slice
func (s GenericSlice) Decode() GenericSlice {
	ns := GenericSlice{}
	for _,encoded := range s {
		n := encoded.(GenericSlice)[0].(int)
		for i := 0; i < n; i++ {
			ns = append(ns, encoded.(GenericSlice)[1])
		}
	}

	return ns
}


//EncodeDirect implements runlength encoding directly without calling other methods...
func (s GenericSlice) EncodeDirect(equal func(interface{}, interface{}) bool) GenericSlice {
	var ns GenericSlice
	if len(s) == 0 {
		return ns
	}

	curVal := GenericSlice{1,s[0]}
	ns = append(ns, curVal)


	for i := 1; i < len(s); i++ {
		if equal(s[i], curVal[1]) {
			curVal[0] = curVal[0].(int) + 1
			ns[len(ns) - 1] = curVal
		} else {
			curVal = GenericSlice{1,s[i]}
			ns = append(ns, curVal)
		}
	}

	return ns
}

func (s GenericSlice) Duplicate() GenericSlice {
	ns := make([]interface{}, 2 * len(s))
	var i int
	for _, v := range s {
		ns[i] = v
		ns[i + 1] = v
		i += 2
	}
	return ns
}

func (s GenericSlice) DuplicateN(n int) GenericSlice {
	ns := make([]interface{}, n * len(s))
	var i int
	for _,v := range s {
		for j := 0; j < n; j++ {
			ns[i] = v
			i += 1
		}
	}
	return ns
}