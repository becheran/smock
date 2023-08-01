package testpackage_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/test/testpackage/cmp"
	testpackage_mock "github.com/test/testpackage/golden_test"
)

func TestSimpleWhen(t *testing.T) {
	m := testpackage_mock.NewMockSimple(t)
	m.WHEN().
		Bar().
		ExpectArgs(cmp.Eq(1), nil).
		Return("1")
	m.WHEN().
		Bar().
		ExpectArgs(cmp.Eq(2), nil).
		Return("2")
	m.WHEN().
		Bar().
		ExpectArgs(cmp.AnyOf(23, 3, 2), nil).
		Return("2")
	assert.Equal(t, "1", m.Bar(1, "foo"))
	assert.Equal(t, "2", m.Bar(2, "bzs"))
	assert.Equal(t, "2", m.Bar(3, ""))
}

func TestFallBackToMatchAll(t *testing.T) {
	m := testpackage_mock.NewMockSimple(t)
	m.WHEN().
		Bar().
		ExpectArgs(cmp.Eq(1), nil).
		Return("1")
	m.WHEN().
		Bar().
		Return("2")
	assert.Equal(t, "1", m.Bar(1, "1"))
	assert.Equal(t, "2", m.Bar(2, "1"))
	assert.Equal(t, "2", m.Bar(3, "1"))
}

func TestFallMatchAllTwiceError(t *testing.T) {
	tester := &Tester{}
	m := testpackage_mock.NewMockSimple(tester)
	m.WHEN().
		Bar().
		Return("1")
	m.WHEN().
		Bar().
		ExpectArgs(nil, nil).
		Return("1")
	assert.Equal(t, "Unreachable condition. Call to 'Bar' is already captured by previous WHEN statement.", tester.errStr)

	tester = &Tester{}
	m = testpackage_mock.NewMockSimple(tester)
	m.WHEN().
		Bar().
		ExpectArgs(nil, nil).
		Return("1")
	m.WHEN().
		Bar().
		Return("1")
	assert.Equal(t, "Unreachable condition. Call to 'Bar' is already captured by previous WHEN statement.", tester.errStr)
}

func TestLambda(t *testing.T) {
	m := testpackage_mock.NewMockWithLambda[string](t)

	m.WHEN().Foo().ExpectArgs(cmp.Eq(1), cmp.Eq("a")).Return(true)
	m.WHEN().Bar().ExpectArgs(cmp.Eq(struct{}{}), cmp.Eq(struct{}{})).Return(true)
	m.WHEN().Baz().ExpectArgs(cmp.Eq("other")).Return(true)

	assert.True(t, m.Foo(1, "a"))
	assert.True(t, m.Bar(struct{}{}, struct{}{}))
	assert.True(t, m.Baz("other"))
}

type Tester struct {
	errStr string
}

func (t *Tester) Fatalf(format string, args ...interface{}) {
	t.errStr = fmt.Sprintf(format, args...)
}

func (t *Tester) Helper() {}
