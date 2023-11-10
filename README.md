# go-filepath
* easily check whether a file is existing, whether it's a directory, etc. This is like a shortcut of "os" and "io/fs" in standard library.
<br><br>

# Import
	import "github.com/lsejx/go-filepath"
<br><br>

# Example

### check filepath
	path := "/"
	t := fpath.GetType(path)
	t.IsExisting()    // true
	t.IsNotExisting() // false
	t.IsRegularFile() // false
	t.IsDir()         // true
	t.IsSymlink()     // false
	t.IsDevice()      // false
	t.IsCharDevice()  // false
	t.IsPipe()        // false
	t.IsSocket()      // false
	t.IsIrregular()   // false
