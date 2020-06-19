package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, Min(3, 5), 3)
}

func TestMax(t *testing.T) {
	assert.Equal(t, Max(3, 5), 5)
}
