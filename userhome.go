package fpath

import (
	"os"
	"path/filepath"
	"strings"
)

func Abs(path string) (string, error) {
	absolutePath := path
	if strings.HasPrefix(path, userHome) {
		absoluteUserHome, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		absolutePath = absoluteUserHome + path[len(userHome):]
	}
	absolutePath, err := filepath.Abs(absolutePath)
	if err != nil {
		return "", err
	}
	return absolutePath, nil
}
