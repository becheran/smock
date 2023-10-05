package smock

import (
	"testing"

	"github.com/becheran/wildmatch-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateMocks(t *testing.T) {
	GenerateMocks(WithDebugLog())
}

func TestSkip(t *testing.T) {
	// Deny list
	assert.False(t, skip("F", []*wildmatch.WildMatch{}, false, true))
	assert.False(t, skip("F", []*wildmatch.WildMatch{wildmatch.NewWildMatch("b")}, false, true))
	assert.True(t, skip("F", []*wildmatch.WildMatch{wildmatch.NewWildMatch("*")}, false, true))
	assert.True(t, skip("F", []*wildmatch.WildMatch{wildmatch.NewWildMatch("b"), wildmatch.NewWildMatch("?")}, false, true))
	assert.True(t, skip("F", []*wildmatch.WildMatch{wildmatch.NewWildMatch("F")}, false, true))

	// Allow list
	assert.False(t, skip("F", []*wildmatch.WildMatch{wildmatch.NewWildMatch("b"), wildmatch.NewWildMatch("?")}, true, true))
	assert.False(t, skip("F", []*wildmatch.WildMatch{wildmatch.NewWildMatch("F")}, true, true))
	assert.True(t, skip("F", []*wildmatch.WildMatch{}, true, true))
	assert.True(t, skip("F", []*wildmatch.WildMatch{}, true, true))

	// Unexported
	assert.True(t, skip("f", nil, false, false))
	assert.False(t, skip("f", nil, false, true))
	assert.False(t, skip("F", nil, false, true))
}
