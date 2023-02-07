package filepath

import "os"

type Type int

const (
	NotExist Type = 1 << iota
	Dir      Type = 1 << iota
)

func GetType(path string) Type {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return NotExist
	}
	if info.IsDir() {
		return Dir
	}
	return 0
}

func IsExistingFile(t Type) bool {
	return t&NotExist == 0 && t&Dir == 0
}

func IsDir(t Type) bool {
	return t&Dir == Dir
}

func IsNotExisting(t Type) bool {
	return t&NotExist == NotExist
}
