package testpackage_test

import (
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
	assert.Equal(t, "1", m.Bar(1, "foo"))
	assert.Equal(t, "2", m.Bar(2, "bzs"))
}
