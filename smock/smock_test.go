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
	assert.False(t, skip("f", []*wildmatch.WildMatch{}, false))
	assert.False(t, skip("f", []*wildmatch.WildMatch{wildmatch.NewWildMatch("b")}, false))
	assert.True(t, skip("f", []*wildmatch.WildMatch{wildmatch.NewWildMatch("*")}, false))
	assert.True(t, skip("f", []*wildmatch.WildMatch{wildmatch.NewWildMatch("b"), wildmatch.NewWildMatch("?")}, false))
	assert.True(t, skip("f", []*wildmatch.WildMatch{wildmatch.NewWildMatch("f")}, false))

	// Allow list
	assert.False(t, skip("f", []*wildmatch.WildMatch{wildmatch.NewWildMatch("b"), wildmatch.NewWildMatch("?")}, true))
	assert.False(t, skip("f", []*wildmatch.WildMatch{wildmatch.NewWildMatch("f")}, true))
	assert.True(t, skip("f", []*wildmatch.WildMatch{}, true))
	assert.True(t, skip("f", []*wildmatch.WildMatch{}, true))
}
