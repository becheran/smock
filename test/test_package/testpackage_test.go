package testpackage_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/becheran/smock/match"
	"github.com/stretchr/testify/assert"
	testpackage_mock "github.com/test/testpackage/golden_test"
)

func TestSimpleWhen(t *testing.T) {
	m := testpackage_mock.NewMockSimple(t)
	m.WHEN().
		Baz().
		Expect(match.Eq(1), nil).
		Return("1")
	m.WHEN().
		Baz().
		Expect(match.Eq(2), nil).
		Return("2")
	m.WHEN().
		Baz().
		Expect(match.Eq(23).Or(match.Eq(3)).Or(match.Eq(2)), nil).
		Return("3")
	assert.Equal(t, "1", m.Baz(1, "foo"))
	assert.Equal(t, "2", m.Baz(2, "bzs"))
	assert.Equal(t, "3", m.Baz(3, ""))
}

func TestFallBackToMatchAll(t *testing.T) {
	m := testpackage_mock.NewMockSimple(t)
	m.WHEN().
		Baz().
		Expect(match.Eq(1), nil).
		Return("1")
	m.WHEN().
		Baz().
		Return("2").
		Times(2)
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

	m.WHEN().Foo().Expect(match.Eq(1), match.Eq("a")).Return(true)
	m.WHEN().Bar().Expect(match.Eq(struct{}{}), match.Eq(struct{}{})).Return(true)
	m.WHEN().Baz().Expect(match.Eq("other")).Return(true)

	assert.True(t, m.Foo(1, "a"))
	assert.True(t, m.Bar(struct{}{}, struct{}{}))
	assert.True(t, m.Baz("other"))
}

func TestUnexpected(t *testing.T) {
	tester := &Tester{t: t}
	m := testpackage_mock.NewMockSimple(tester)

	m.Bar(1, "2", struct{}{}, &struct{}{}, true, []byte{1, 2, 3})

	assert.Equal(t, `Unexpected call to Bar(). If function call is expected add ".WHEN.Bar()" to mock.`, tester.errStr)
}

func TestAnyTimes(t *testing.T) {
	times(t, -1, 30, "")
	times(t, -1, 1, "")
	times(t, -1, 0, "")
}
func TestNeverFail(t *testing.T) {
	times(t, 0, 1, "Expected 'Foo' to be called 0 times, but was called 1 times.")
	times(t, 0, 2, "Expected 'Foo' to be called 0 times, but was called 2 times.")
}
func TestNeverSuccess(t *testing.T) {
	times(t, 0, 0, "")
}
func TestOnceFail(t *testing.T) {
	times(t, 1, 0, "Expected 'Foo' to be called 1 times, but was called 0 times.")
	times(t, 1, 2, "Expected 'Foo' to be called 1 times, but was called 2 times.")
}
func TestOnceSuccess(t *testing.T) {
	times(t, 1, 1, "")
}
func times(t *testing.T, expected, times int, err string) {
	tester := &Tester{t: t}
	m := testpackage_mock.NewMockSimple(tester)
	if expected < 0 {
		m.WHEN().Foo().AnyTimes()
		m.WHEN().Bar().Return("").AnyTimes()
		m.WHEN().SingleArg().Expect(func(i int) bool { return i == 42 }).AnyTimes()
	} else if expected == 0 {
		m.WHEN().Foo().Never()
		m.WHEN().Bar().Return("").Never()
		m.WHEN().SingleArg().Expect(func(i int) bool { return i == 42 }).Never()
	} else if expected == 1 {
		m.WHEN().Foo()
		m.WHEN().Bar().Return("")
		m.WHEN().SingleArg().Expect(func(i int) bool { return i == 42 })
	} else {
		m.WHEN().Foo().Times(expected)
		m.WHEN().Bar().Return("").Times(expected)
		m.WHEN().SingleArg().Expect(func(i int) bool { return i == 42 }).Times(expected)
	}
	m.WHEN().SingleArg().AnyTimes()
	m.SingleArg(1)
	for i := 0; i < times; i++ {
		m.Foo()
		m.Bar(0, "", struct{}{}, nil, nil, nil)
		m.SingleArg(42)
	}
	tester.cleanup()
	if err != "" {
		assert.Contains(t, tester.errStr, err)
	} else {
		assert.Empty(t, tester.errStr)
	}
}

func TestThen(t *testing.T) {
	tester := &Tester{t: t}
	m := testpackage_mock.NewMockSimple(tester)
	m.WHEN().
		Foo().
		Do(func() { fmt.Println("1") }).
		Times(1).
		Then().
		Do(func() { fmt.Println("2") }).
		Times(2).
		Then().
		Do(func() { fmt.Println("3") }).
		Times(3)

	for i := 0; i < 10; i++ {
		m.Foo()
	}

	tester.cleanup()

	assert.Equal(t, `
Expected 'Foo' to be called 3 times, but was called 7 times.`, tester.errStr)
}

func TestThenPass(t *testing.T) {
	m := testpackage_mock.NewMockSimple(t)
	m.WHEN().
		Foo().
		Times(3).
		Then().
		Do(func() { fmt.Println("2") }).
		Times(3).
		Then().
		Do(func() { fmt.Println("3") }).
		Times(3).
		Then()

	for i := 0; i < 10; i++ {
		m.Foo()
	}
}

func TestThenFailWhenUnreachable(t *testing.T) {
	fatalChan := make(chan bool)
	go func() {
		<-fatalChan
	}()
	tester := &Tester{t: t, fatalChan: fatalChan}

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, `Then statement is not reachable. Expected calls of previous statement: AnyTimes`, tester.errStr)
		}
	}()

	m := testpackage_mock.NewMockSimple(tester)
	m.WHEN().
		Foo().
		Do(func() { fmt.Println("1") }).
		Times(-1).
		Then()

	m.Foo()
}

func TestThenWithExpect(t *testing.T) {
	tester := &Tester{t: t}

	m := testpackage_mock.NewMockSimple(tester)
	m.WHEN().
		Bar().Expect(nil, nil, nil, nil, nil, func(b []byte) bool { return len(b) == 0 }).
		Return("One").
		Then().
		Return("Two").
		Then().
		Return("Two").
		Then().
		Return("Three").
		Times(3)

	assert.Equal(t, "One", m.Bar(0, "", struct{}{}, nil, nil, nil))
	assert.Equal(t, "Two", m.Bar(0, "", struct{}{}, nil, nil, nil))
	assert.Equal(t, "Two", m.Bar(0, "", struct{}{}, nil, nil, nil))
	assert.Equal(t, "Three", m.Bar(0, "", struct{}{}, nil, nil, nil))
	assert.Equal(t, "Three", m.Bar(0, "", struct{}{}, nil, nil, nil))
	assert.Equal(t, "Three", m.Bar(0, "", struct{}{}, nil, nil, nil))
}

func TestCallThreadSafe(t *testing.T) {
	m := testpackage_mock.NewMockWithLambda[string](t)

	m.WHEN().Foo().Times(2)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		m.Foo(0, "")
		wg.Done()
	}()
	m.Foo(0, "")
	wg.Wait()
}

type Tester struct {
	t         *testing.T
	errStr    string
	cleanup   func()
	fatalChan chan bool
}

func (t *Tester) Fatalf(format string, args ...any) {
	t.errStr = fmt.Sprintf(format, args...)
	if t.fatalChan != nil {
		t.fatalChan <- true
	}
}

func (t *Tester) Helper() {
	if t.t != nil {
		t.t.Helper()
	}
}

func (t *Tester) Cleanup(c func()) {
	t.cleanup = c
}
