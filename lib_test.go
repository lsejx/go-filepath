package filepath

import "testing"

func TestGetType(t *testing.T) {
	type data struct {
		path string
		want Type
	}

	// verify that
	//     "not_existing_file" is not existing,
	//     "test_file" is existing as a normal file, and
	//     "test_dir" is existing as a directory
	// before start testing

	testDatas := []data{
		{"not_existing_file", NotFound},
		{"test_file", 0},
		{"test_dir", Directory},
	}

	for _, data := range testDatas {
		tp := GetType(data.path)
		if tp != data.want {
			t.Fatalf("path: %v, want: %v, got: %v", data.path, data.want, tp)
		}
	}
}

func TestIsExistingFile(t *testing.T) {
	path := "test_file"
	exists := IsExistingFile(GetType(path))
	if !exists {
		t.Fatal()
	}
}
