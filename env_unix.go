package fpath

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// ResolveEnv resolves environment variables by os.LookupEnv(key).
// Special variables (e.g. $!, $$, and $?) is not resolved by this, ResolveShellEnv(path) resolves them.
// ~ (user home dir) is not resolved, AbsHome(path) or ResolveShellEnv(path) resolves it.
func ResolveEnv(path string) (string, error) {
	for i := 0; i < len(path); i++ {
		if path[i] == '$' && (i != len(path)-1 || path[i+1] != '/' && path[i+1] != '$') {
			var envEnd int
			for envEnd = i + 2; envEnd < len(path); envEnd++ {
				if path[envEnd] == '/' || path[envEnd] == '$' {
					break
				}
			}
			value, present := os.LookupEnv(path[i+1 : envEnd])
			if !present {
				return "", fmt.Errorf("%v is absent", path[i:envEnd])
			}
			path = path[:i] + value + path[envEnd:]
			i += len(value) - 1
		}
	}
	return path, nil
}

// ResolveShellEnv executes $SHELL -c 'echo -n {path}'.
// Non-nil error is returned if $SHELL is absent.
func ResolveShellEnv(path string) (string, error) {
	shell, present := os.LookupEnv("SHELL")
	if !present {
		return "", errors.New("$SHELL is absent")
	}

	cmd := exec.Command(shell, "-c", fmt.Sprintf("echo -n %v", path))
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
