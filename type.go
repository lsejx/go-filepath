package fpath

import "os"

type Type int

const (
	Absent  Type = 0
	Present Type = 1 << iota
	Dir     Type = 1 << iota
)

func GetType(path string) Type {
	flag := Type(0)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return Absent
	}
	flag |= Present
	if info.IsDir() {
		flag |= Dir
	}
	return flag
}

func (t Type) IsFile() bool {
	return t == Present
}

func (t Type) IsDir() bool {
	return t&Dir == Dir
}

func (t Type) IsExisting() bool {
	return t&Present == Present
}

func (t Type) IsNotExisting() bool {
	return t == Absent
}
