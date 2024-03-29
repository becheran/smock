// Package match contains match functions which can be used to create expressions that evaluate to a truth value.
//
// All functions in this package result in a [Match] type which is simply a function with one generic input and a bool return value.
// The combination functions [Match.And] and [Match.Or] can be used to combine two [Match] expressions.
// [Not] will negate a [Match] expression.
package match

import "reflect"

// Match is a simple expression which evaluates to be either 'true' or 'false'
type Match[T any] func(e T) bool

// Any is the verbose version of 'nil' and will always evaluate to true when passed as expression
func Any[T any]() Match[T] {
	return func(T) bool { return true }
}

// And combines two expressions with a logical AND `&&`
func (m Match[T]) And(o Match[T]) Match[T] {
	if o == nil {
		return m
	}
	return func(e T) bool { return m(e) && o(e) }
}

// Or combines two expressions with a logical OR `||`
func (m Match[T]) Or(o Match[T]) Match[T] {
	if o == nil {
		return m
	}
	return func(e T) bool { return m(e) || o(e) }
}

// Not negates a logical expression
func Not[T any](o Match[T]) Match[T] {
	if o == nil {
		return func(a T) bool { return false }
	}
	return func(e T) bool { return !o(e) }
}

// Comparable

// Eq returns a function which will be true if `other` is equal to the input
func Eq[T comparable](other T) Match[T] {
	return func(val T) bool {
		return val == other
	}
}

// EAnyOfq returns a function which will be true if any value of `other` is equal to the input
func AnyOf[T comparable](other ...T) Match[T] {
	return func(val T) bool {
		for _, o := range other {
			if val == o {
				return true
			}
		}
		return false
	}
}

// Any

// DeepEqual returns a function which will be true if `other`
// is equal to the input using to the [reflect.DeepEqual] function
func DeepEqual[T any](other T) Match[T] {
	return func(val T) bool {
		return reflect.DeepEqual(val, other)
	}
}

// Slice

// Eq returns a function which will be true if all elements of `other` are equal to the input
func SliceEq[T ~[]E, E comparable](other T) Match[T] {
	return func(val T) bool {
		if len(val) != len(other) {
			return false
		}
		for i := range val {
			if val[i] != other[i] {
				return false
			}
		}
		return true
	}
}

// Maps

// MapEq returns a function which will be true if `other` contains the same key/value pairs as in the input
func MapEq[T ~map[K]V, K, V comparable](other T) Match[T] {
	return func(val T) bool {
		if len(val) != len(other) {
			return false
		}
		for k, v1 := range val {
			if v2, ok := other[k]; !ok || v1 != v2 {
				return false
			}
		}
		return true
	}
}
