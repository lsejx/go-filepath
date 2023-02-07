package filepath

import "os"

type Type int

const (
	NotFound  Type = 1 << iota
	Directory Type = 1 << iota
)

func GetType(path string) Type {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return NotFound
	}
	if info.IsDir() {
		return Directory
	}
	return 0
}

func IsExistingFile(t Type) bool {
	return t&NotFound == 0 && t&Directory == 0
}
