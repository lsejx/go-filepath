package fpath

import "testing"

func TestAbs(t *testing.T) {
	datas := []struct {
		in   string
		want string
	}{
		{"~", "/root"},
		{"~/.bashrc", "/root/.bashrc"},
		{"go.mod", "/usr/src/filepath/go.mod"},
		{".", "/usr/src/filepath"},
	}

	for _, data := range datas {
		out, err := Abs(data.in)
		t.Log(out)
		if err != nil {
			t.Fatal(err)
		}
		if out != data.want {
			t.Fatalf("want: %v, got: %v", data.want, out)
		}
	}
}
