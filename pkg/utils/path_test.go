package utils

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandPath(t *testing.T) {
	home, _ := os.UserHomeDir()
	filename, _ := ExpandPath("~/foo/bar")
	assert.Equal(t, filename, path.Join(home, "foo", "bar"))
}
