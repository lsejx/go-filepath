package fpath

import "testing"

func TestGetType(t *testing.T) {
	tests := []struct {
		path          string
		want          Type
		isFile        bool
		isDir         bool
		isExisting    bool
		isNotExisting bool
	}{
		{"absent_file", Type{false, false}, false, false, false, true},
		{"type_test.go", Type{true, false}, true, false, true, false},
		{".git", Type{true, true}, false, true, true, false},
	}

	for _, tt := range tests {
		tp := GetType(tt.path)
		if tp != tt.want {
			t.Fatalf("path:%v, want:%v, got:%v", tt.path, tt.want, tp)
		}
		if tp.IsFile() != tt.isFile {
			t.Fatalf("tp.IsFile():%v, isFile:%v", tp.IsFile(), tt.isFile)
		}
		if tp.IsDir() != tt.isDir {
			t.Fatalf("tp.IsDir():%v, data.isDir:%v", tp.IsDir(), tt.isDir)
		}
		if tp.IsExisting() != tt.isExisting {
			t.Fatalf("tp.IsExisting():%v, data.isExisting:%v", tp.IsExisting(), tt.isExisting)
		}
		if tp.IsNotExisting() != tt.isNotExisting {
			t.Fatalf("tp.IsNotExisting():%v, data.isNotExisting:%v", tp.IsNotExisting(), tt.isNotExisting)
		}
	}
}
