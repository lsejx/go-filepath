package fpath

import (
	"os"
	"strings"
)

// AbsHome replaces ~ (user home dir) with os.UserHomeDir().
// Non-nil error is returned if os.UserHomeDir() returned non-nil error.
func AbsHome(path string) (string, error) {
	apply := path == "~" || strings.HasPrefix(path, "~/")
	if !apply {
		return path, nil
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + path[1:], nil
}
