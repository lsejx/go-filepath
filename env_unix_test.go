package fpath

import (
	"testing"
)

func TestResolveEnv(t *testing.T) {
	okDatas := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"aiueo/kakikukeko", "aiueo/kakikukeko"},
		{"$GOPATH", "/go"},
		{"$GOPATH/bin", "/go/bin"},
		{"aiueo$GOPATH", "aiueo/go"},
		{"aiueo$GOPATH/kakikukeko", "aiueo/go/kakikukeko"},
		{"$GOPATH$GOPATH", "/go/go"},
	}

	// frequently used
	e0 := "$XXX is absent"
	badDatas := []struct {
		in      string
		wantErr string
	}{
		{"$XXX", e0},
		{"$XXX/bin", e0},
		{"/aiueo/$XXX", e0},
		{"/aiueo/$XXX/kakikukeko", e0},
		{"$XXX/$YYY", e0},
		{"$#/$XXX", "$# is absent"},
		{"$$", "$$ is absent"},
	}

	for _, okData := range okDatas {
		got, err := ResolveEnv(okData.in)
		if err != nil {
			t.Fatalf("err:%v, in:%q", err, okData.in)
		}
		if okData.want != got {
			t.Fatalf("in:%q, got:%q", okData.in, got)
		}
	}

	for _, badData := range badDatas {
		got, err := ResolveEnv(badData.in)
		if err == nil {
			t.Fatalf("nil err, in %q, got:%q", badData.in, got)
		}
		if badData.wantErr != err.Error() {
			t.Fatalf("wantErr:%v, err:%v", badData.wantErr, err)
		}
	}
}

func TestResolveShellEnv(t *testing.T) {
	okDatas := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"aiueo/kakikukeko", "aiueo/kakikukeko"},
		{"$GOPATH", "/go"},
		{"$GOPATH/bin", "/go/bin"},
		{"aiueo$GOPATH", "aiueo/go"},
		{"aiueo$GOPATH/kakikukeko", "aiueo/go/kakikukeko"},
		{"$GOPATH$GOPATH", "/go/go"},
	}

	for _, okData := range okDatas {
		got, err := ResolveShellEnv(okData.in)
		if err != nil {
			t.Fatalf("err:%v, in:%q", err, okData.in)
		}
		if okData.want != got {
			t.Fatalf("in:%q, got:%q", okData.in, got)
		}
	}

	datas := []string{
		"$$",
		"$?",
		"~",
		"~/aiueo",
	}

	for _, data := range datas {
		got, err := ResolveShellEnv(data)
		if err != nil {
			t.Fatalf("err:%v, in:%q", err, data)
		}
		t.Logf("in:%q, got:%q", data, got)
	}
}
