package fpath

import "testing"

func TestGetType(t *testing.T) {
	datas := []struct {
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

	for _, data := range datas {
		tp := GetType(data.path)
		if tp != data.want {
			t.Fatalf("path:%v, want:%v, got:%v", data.path, data.want, tp)
		}
		if tp.IsFile() != data.isFile {
			t.Fatalf("tp.IsFile():%v, isFile:%v", tp.IsFile(), data.isFile)
		}
		if tp.IsDir() != data.isDir {
			t.Fatalf("tp.IsDir():%v, data.isDir:%v", tp.IsDir(), data.isDir)
		}
		if tp.IsExisting() != data.isExisting {
			t.Fatalf("tp.IsExisting():%v, data.isExisting:%v", tp.IsExisting(), data.isExisting)
		}
		if tp.IsNotExisting() != data.isNotExisting {
			t.Fatalf("tp.IsNotExisting():%v, data.isNotExisting:%v", tp.IsNotExisting(), data.isNotExisting)
		}
	}
}
