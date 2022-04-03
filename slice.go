/*	Copyright (c) 2022, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package types

// ComplexSlice is the set of complex number slice types.
type ComplexSlice interface {
	~[]complex64 | ~[]complex128
}

// FloatSlice is the set of IEEE-754 floating-point slice types.
type FloatSlice interface {
	~[]float32 | ~[]float64
}

// SignedSlice is the set of signed integer slice types.
type SignedSlice interface {
	~[]int | ~[]int8 | ~[]int16 | ~[]int32 | ~[]int64
}

// UnsignedSlice is the set of unsigned integer slice types.
type UnsignedSlice interface {
	~[]uint | ~[]uint8 | ~[]uint16 | ~[]uint32 | ~[]uint64 | ~[]uintptr
}

// IntegerSlice is the set of integer slice types.
type IntegerSlice interface {
	SignedSlice | UnsignedSlice
}

// StringySlice is the set of string-like slice types.
type StringySlice interface {
	~[]string | ~[][]byte
}

// HaveCapSlice is the set of slice types whose elements have dynamic capacity
// accessed with cap().
type HaveCapSlice[T any] interface {
	~[][]T | ~[]chan T | ~[]<-chan T | ~[]chan<- T
}

// HaveLenSlice is the set of slice types whose elements have dynamic length
// accessed with len().
type HaveLenSlice[T any, K comparable] interface {
	HaveCapSlice[T] | ~[]string | ~[]map[K]T
}
