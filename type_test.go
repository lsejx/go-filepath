package fpath

import "testing"

func TestGetType(t *testing.T) {
	type data struct {
		path string
		want Type
	}

	testDatas := []data{
		{"absent_file", 0},
		{"type_test.go", Present},
		{".git", Dir | Present},
	}

	for _, data := range testDatas {
		tp := GetType(data.path)
		if tp != data.want {
			t.Fatalf("path: %v, want: %v, got: %v", data.path, data.want, tp)
		}
	}
}
