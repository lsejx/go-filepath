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

func (t Type) IsFile() bool {
	return t.isExisting && !t.isDir
}

func (t Type) IsDir() bool {
	return t.isDir
}

func (t Type) IsExisting() bool {
	return t.isExisting
}

func (t Type) IsNotExisting() bool {
	return !t.isExisting
}
