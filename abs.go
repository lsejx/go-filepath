package fpath

import "path/filepath"

// Abs calls AbsHome(path), ResolveEnv(path), and filepath.Abs(path).
func Abs(path string) (string, error) {
	abs, err := AbsHome(path)
	if err != nil {
		return "", err
	}
	abs, err = ResolveEnv(abs)
	if err != nil {
		return "", err
	}
	abs, err = filepath.Abs(abs)
	if err != nil {
		return "", err
	}
	return abs, nil
}

// AbsShell calls ResolveShellEnv(path), and filepath.Abs(path).
func AbsShell(path string) (string, error) {
	abs, err := ResolveShellEnv(path)
	if err != nil {
		return "", err
	}
	abs, err = filepath.Abs(abs)
	if err != nil {
		return "", err
	}
	return abs, nil
}
