package cmp

type Match[T any] func(e T) bool

func (m Match[T]) And(o Match[T]) Match[T] {
	if o == nil {
		return m
	}
	return func(e T) bool { return m(e) && o(e) }
}

func (m Match[T]) Or(o Match[T]) Match[T] {
	if o == nil {
		return m
	}
	return func(e T) bool { return m(e) || o(e) }
}

func Not[T any](o Match[T]) Match[T] {
	if o == nil {
		return func(a T) bool { return false }
	}
	return func(e T) bool { return !o(e) }
}

func Eq[T comparable](other T) Match[T] {
	return func(val T) bool {
		return val == other
	}
}
