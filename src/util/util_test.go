package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeRandInt(t *testing.T) {
	elements := MakeRandInt(10)
	assert.NotNil(t, elements, "elements array should not be nil")
	assert.Equal(t, len(elements), 10, "elements array should not be nil")
}
