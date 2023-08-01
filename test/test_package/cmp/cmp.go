package cmp

import "golang.org/x/exp/slices"

type CompareExpression[T comparable] struct {
	evaluate func(e T) bool
}

func (ce *CompareExpression[T]) Match(val T) bool {
	if ce.evaluate == nil {
		return true
	}
	return ce.evaluate(val)
}

func Eq[T comparable](other T) *CompareExpression[T] {
	return &CompareExpression[T]{
		evaluate: func(val T) bool {
			return val == other
		},
	}
}

func AnyOf[T comparable](other ...T) *CompareExpression[T] {
	return &CompareExpression[T]{
		evaluate: func(val T) bool {
			return slices.Contains(other, val)
		},
	}
}
