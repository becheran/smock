package match_test

import (
	"fmt"
	"testing"

	"github.com/becheran/smock/match"
	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	c := make(chan bool)

	assert.True(t, match.Eq("foo")("foo"))
	assert.True(t, match.Eq(1)(1))
	assert.True(t, match.Eq("foo")("foo"))
	assert.True(t, match.Eq("foo")("foo"))
	assert.True(t, match.Eq(c)(c))
	assert.True(t, match.Eq([2]string{"foo", "bar"})([2]string{"foo", "bar"}))

	assert.False(t, match.Eq(c)(nil))
	assert.False(t, match.Eq(c)(make(chan bool)))
	assert.False(t, match.Eq([2]string{"foo", "bar"})([2]string{"foo", "bz"}))
	assert.False(t, match.Eq[string]("")("foo"))
	assert.False(t, match.Eq[string]("foo")("bar"))
}

func TestAnyOf(t *testing.T) {
	assert.True(t, match.AnyOf("foo", "bar")("bar"))
	assert.False(t, match.AnyOf("foo", "bar")("baz"))
}

func TestOr(t *testing.T) {
	assert.True(t, match.Eq("foo").Or(match.Eq("bar"))("foo"))
	assert.True(t, match.Eq("foo").Or(match.Eq("bar"))("bar"))
	assert.False(t, match.Eq("foo").Or(match.Eq("bar"))("other"))
}

func TestNot(t *testing.T) {
	assert.False(t, match.Not(match.Eq("foo"))("foo"))
	assert.True(t, match.Not(match.Eq("foo"))("bar"))
}

func TestAnd(t *testing.T) {
	assert.True(t, match.AnyOf("foo", "bar").And(match.AnyOf("bar", "foo"))("foo"))
	assert.False(t, match.AnyOf("foo", "bar").And(match.AnyOf("bar", "f"))("foo"))
}

func TestSliceEq(t *testing.T) {
	assert.True(t, match.SliceEq[[]int](nil)(nil))
	assert.False(t, match.SliceEq([]int{1, 2})(nil))
	assert.False(t, match.SliceEq([]int{1, 2})([]int{1, 2, 3}))
}

func TestMapEq(t *testing.T) {
	assert.True(t, match.MapEq[map[string]int](nil)(nil))
	assert.False(t, match.MapEq(map[int]int{0: 1})(nil))
	assert.False(t, match.MapEq(map[int]int{0: 1})(map[int]int{1: 1}))
	assert.False(t, match.MapEq(map[int]int{0: 1})(map[int]int{1: 1, 0: 1}))
	assert.True(t, match.MapEq(map[int]int{0: 1})(map[int]int{0: 1}))
}

func TestAny(t *testing.T) {
	assert.True(t, match.Any[string]()(""))
	assert.True(t, match.Any[any]()(nil))
	assert.True(t, match.Any[any]()("bar"))
}

func Example() {
	fmt.Println(match.Eq("foo").Or(match.Eq("bar"))("foo"))
	fmt.Println(match.Eq(1).Or(match.Eq(2))(3))
	fmt.Println(match.Not(match.Eq(1))(1))
	// Output:
	// true
	// false
	// false
}
