package fpath

import (
	"testing"
)

func TestType(t *testing.T) {
	tests := []struct {
		path string
		exi  bool
		reg  bool
		dir  bool
		sym  bool
		dev  bool
		cha  bool
		pip  bool
		soc  bool
		irr  bool
	}{
		{"./test/______", false, false, false, false, false, false, false, false, false},
		{"./test/regular", true, true, false, false, false, false, false, false, false},
		{"./test/", true, false, true, false, false, false, false, false, false},
		{"./test/slink", true, false, false, true, false, false, false, false, false},
		{"/dev/null", true, false, false, false, true, true, false, false, false},
		{"./test/pipe", true, false, false, false, false, false, true, false, false},
	}

	for _, tt := range tests {
		tp := GetType(tt.path)
		switch {
		case tp.IsExisting() != tt.exi:
			t.Fatal(tt.path, "exi")
		case tp.IsRegularFile() != tt.reg:
			t.Fatal(tt.path, "reg")
		case tp.IsDir() != tt.dir:
			t.Fatal(tt.path, "dir")
		case tp.IsSymlink() != tt.sym:
			t.Fatal(tt.path, "sym")
		case tp.IsDevice() != tt.dev:
			t.Fatal(tt.path, "dev")
		case tp.IsCharDevice() != tt.cha:
			t.Fatal(tt.path, "cha")
		case tp.IsPipe() != tt.pip:
			t.Fatal(tt.path, "pip")
		case tp.IsSocket() != tt.soc:
			t.Fatal(tt.path, "soc")
		case tp.IsIrregular() != tt.irr:
			t.Fatal(tt.path, "irr")
		}
	}
}
