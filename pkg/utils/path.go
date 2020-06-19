package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// ExpandPath expand tilde in file path
func ExpandPath(path string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if path == "~" {
		path = home
	} else if strings.HasPrefix(path, "~/") {
		path = filepath.Join(home, path[2:])
	}
	return path, nil
}
