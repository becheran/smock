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
		Baz().
		Expect(cmp.Eq(1), nil).
		Return("1")
	m.WHEN().
		Baz().
		Expect(cmp.Eq(2), nil).
		Return("2")
	m.WHEN().
		Baz().
		Expect(cmp.Eq(23).Or(cmp.Eq(3)).Or(cmp.Eq(2)), nil).
		Return("2")
	assert.Equal(t, "1", m.Baz(1, "foo"))
	assert.Equal(t, "2", m.Baz(2, "bzs"))
	assert.Equal(t, "2", m.Baz(3, ""))
}

func TestFallBackToMatchAll(t *testing.T) {
	m := testpackage_mock.NewMockSimple(t)
	m.WHEN().
		Baz().
		Expect(cmp.Eq(1), nil).
		Return("1")
	m.WHEN().
		Baz().
		Return("2")
	assert.Equal(t, "1", m.Baz(1, "1"))
	assert.Equal(t, "2", m.Baz(2, "1"))
	assert.Equal(t, "2", m.Baz(3, "1"))
}

func TestFallMatchAllTwiceError(t *testing.T) {
	tester := &Tester{}
	m := testpackage_mock.NewMockSimple(tester)
	m.WHEN().
		Bar().
		Return("1")
	m.WHEN().
		Bar().
		Expect(nil, nil, nil, nil, nil, nil).
		Return("1")
	assert.Equal(t, "Unreachable condition. Call to 'Bar' is already captured by previous WHEN statement.", tester.errStr)

	tester = &Tester{}
	m = testpackage_mock.NewMockSimple(tester)
	m.WHEN().
		Bar().
		Expect(nil, nil, nil, nil, nil, nil).
		Return("1")
	m.WHEN().
		Bar().
		Return("1")
	assert.Equal(t, "Unreachable condition. Call to 'Bar' is already captured by previous WHEN statement.", tester.errStr)
}

func TestLambda(t *testing.T) {
	m := testpackage_mock.NewMockWithLambda[string](t)

	m.WHEN().Foo().Expect(cmp.Eq(1), cmp.Eq("a")).Return(true)
	m.WHEN().Bar().Expect(cmp.Eq(struct{}{}), cmp.Eq(struct{}{})).Return(true)
	m.WHEN().Baz().Expect(cmp.Eq("other")).Return(true)

	assert.True(t, m.Foo(1, "a"))
	assert.True(t, m.Bar(struct{}{}, struct{}{}))
	assert.True(t, m.Baz("other"))
}

func TestUnexpected(t *testing.T) {
	tester := &Tester{t: t}
	m := testpackage_mock.NewMockSimple(tester)

	m.Bar(1, "2", struct{}{}, &struct{}{}, true, []byte{1, 2, 3})

	assert.Equal(t, `Unexpected call Bar(1, "2", {}, &{}, true, [1 2 3])`, tester.errStr)
}

type Tester struct {
	t      *testing.T
	errStr string
}

func (t *Tester) Fatalf(format string, args ...interface{}) {
	t.errStr = fmt.Sprintf(format, args...)
}

func (t *Tester) Helper() {
	if t.t != nil {
		t.t.Helper()
	}
}
