package fpath

import "os"

type Type struct {
	isExisting bool
	isDir      bool
}

func GetType(path string) Type {
	tp := Type{false, false}
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return tp
	}
	tp.isExisting = true
	if info.IsDir() {
		tp.isDir = true
	}
	return tp
}

// IsFile returns true if the file is existing and it isn't a directory.
// If it isn't existing, this returns false.
func (t Type) IsFile() bool {
	return t.isExisting && !t.isDir
}

// IsDir returns true if the file is existing and it isn't a directory.
// If it isn't existing, this returns false.
func (t Type) IsDir() bool {
	return t.isDir
}

// IsExisting returns true if the file is existing, despite whether it is a directory.
func (t Type) IsExisting() bool {
	return t.isExisting
}

// IsNotExisting returns just NOT of IsExisting.
func (t Type) IsNotExisting() bool {
	return !t.isExisting
}
