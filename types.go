/*	Copyright (c) 2022, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

// Package types defines constraint type sets for Go Generics.
package types

// Complex is the set of complex number types.
type Complex interface {
	~complex64 | ~complex128
}

// Float is the set of IEEE-754 floating-point types.
type Float interface {
	~float32 | ~float64
}

// Signed is the set of signed integer types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is the set of unsigned integer types.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is the set of integer types.
type Integer interface {
	Signed | Unsigned
}

// Stringy is the set of string-like types.
type Stringy interface {
	~string | ~[]byte
}

// HaveCap is the set of types that have dynamic capacity accessed with cap().
type HaveCap[T any] interface {
	~[]T | ~chan T | ~<-chan T | ~chan<- T
}

// HaveLen is the set of types that have dynamic length accessed with len().
type HaveLen[T any, K comparable] interface {
	HaveCap[T] | ~string | ~map[K]T
}
