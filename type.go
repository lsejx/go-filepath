package fpath

import (
	"io/fs"
	"os"
)

type Type struct {
	isExisting bool
	mode       fs.FileMode
}

// GetType checks the filepath and returns struct that contains the result of checking.
// If os.Lstat returns non-nil error, GetType returns it.
func GetType(path string) (Type, error) {
	tp := Type{isExisting: false}
	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return tp, nil
		}
		return Type{}, err
	}
	tp.isExisting = true
	tp.mode = info.Mode()
	return tp, nil
}

// IsExisting returns true on existing path despite whether it is a regular file, directory or special file.
// This is based on os.Stat.
func (t Type) IsExisting() bool {
	return t.isExisting
}

// IsNotExisting returns NOT of IsExisting.
func (t Type) IsNotExisting() bool {
	return !t.isExisting
}

// IsRegularFile returns true if it's a regular file (not a directory and a special file.)
// If it's a special file like device, pipe, symlink, socket, etc, this returns false.
// This is based on io/fs.FileMode.
func (t Type) IsRegularFile() bool {
	if !t.isExisting {
		return false
	}
	return t.mode.IsRegular()
}

// IsDir returns true if it's a directory.
// This is based on io/fs.FileMode.
func (t Type) IsDir() bool {
	return t.mode.IsDir()
}

// IsSymlink returns true if it's a symbolic link.
// This is based on io/fs.FileMode.
func (t Type) IsSymlink() bool {
	return t.mode&fs.ModeSymlink != 0
}

// IsDevice returns true if it's a device file despite whether it's a block device or character device.
// This is based on io/fs.FileMode.
func (t Type) IsDevice() bool {
	return t.mode&fs.ModeDevice != 0
}

// IsCharDevice returns true if it's a character device file.
// IsDevice returns true when IsCharDevice returns true.
// This is based on io/fs.FileMode.
func (t Type) IsCharDevice() bool {
	return t.mode&fs.ModeCharDevice != 0
}

// IsPipe returns true if it's a named pipe.
// This is based on io/fs.FileMode.
func (t Type) IsPipe() bool {
	return t.mode&fs.ModeNamedPipe != 0
}

// IsSocket returns true if it's a socket.
// This is based on io/fs.FileMode.
func (t Type) IsSocket() bool {
	return t.mode&fs.ModeSocket != 0
}

// IsIrregular returns true if it's non-regular file; nothing else is known about this file (see pkg.go.dev/io/fs#FileMode).
// This is based on io/fs.FileMode.
func (t Type) IsIrregular() bool {
	return t.mode&fs.ModeIrregular != 0
}
