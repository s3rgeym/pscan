package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const str string = `this is
multiline
string`

func TestReadLines(t *testing.T) {
	lines, _ := ReadLines(strings.NewReader(str))
	assert.ElementsMatch(t, lines, []string{"this is", "multiline", "string"})
}
