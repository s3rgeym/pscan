package utils

import (
	"bufio"
	"io"
)

// ReadLines fn
func ReadLines(r io.Reader) ([]string, error) {
	s := bufio.NewScanner(r)
	if err := s.Err(); err != nil {
		return nil, err
	}
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, nil
}
