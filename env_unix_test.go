package fpath

import (
	"errors"
	"testing"
)

func TestResolveEnv(t *testing.T) {
	e0 := "$XXX is absent"
	tests := []struct {
		arg string
		ret string
		err error
	}{
		{"", "", nil},
		{"aiueo/kakikukeko", "aiueo/kakikukeko", nil},
		{"$GOPATH", "/go", nil},
		{"$GOPATH/bin", "/go/bin", nil},
		{"aiueo$GOPATH", "aiueo/go", nil},
		{"aiueo$GOPATH/kakikukeko", "aiueo/go/kakikukeko", nil},
		{"$GOPATH$GOPATH", "/go/go", nil},
		{"$XXX", "", errors.New(e0)},
		{"$XXX/bin", "", errors.New(e0)},
		{"/aiueo/$XXX", "", errors.New(e0)},
		{"/aiueo/$XXX/kakikukeko", "", errors.New(e0)},
		{"$XXX/$YYY", "", errors.New(e0)},
		{"$#/$XXX", "", errors.New("$# is absent")},
		{"$$", "", errors.New("$$ is absent")},
	}

	for _, tt := range tests {
		s, err := ResolveEnv(tt.arg)
		if tt.err == nil {
			if err != nil {
				t.Fatalf("a:%v, err:%v", tt.arg, err)
			}
			// ok
		} else {
			// non-nil-err case
			if err == nil {
				t.Fatalf("a:%v, nilerr, s:%v", tt.arg, s)
			}
			if tt.err.Error() != err.Error() {
				t.Fatalf("a:%v, err:%v, w:%v", tt.arg, err, tt.err)
			}
		}
	}
}

func TestResolveShellEnv(t *testing.T) {
	tests := []struct {
		arg string
		ret string
		err error
	}{
		{"", "", nil},
		{"aiueo/kakikukeko", "aiueo/kakikukeko", nil},
		{"$GOPATH", "/go", nil},
		{"$GOPATH/bin", "/go/bin", nil},
		{"aiueo$GOPATH", "aiueo/go", nil},
		{"aiueo$GOPATH/kakikukeko", "aiueo/go/kakikukeko", nil},
		{"$GOPATH$GOPATH", "/go/go", nil},
	}

	for _, tt := range tests {
		s, err := ResolveShellEnv(tt.arg)
		if tt.err == nil {
			if err != nil {
				t.Fatalf("a:%v, err:%v", tt.arg, err)
			}
			// ok
		} else {
			if err == nil {
				t.Fatalf("a:%v, nilerr, s:%v", tt.arg, s)
			}
			if tt.err.Error() != err.Error() {
				t.Fatalf("a:%v, err:%v, w:%v", tt.arg, err, tt.err)
			}
		}
	}

	tests2 := []string{
		"$$",
		"$?",
		"~",
		"~/aiueo",
	}

	for _, tt := range tests2 {
		s, err := ResolveShellEnv(tt)
		if err != nil {
			t.Fatalf("err:%v, a:%q", err, tt)
		}
		t.Logf("a:%q, s:%q", tt, s)
	}
}
