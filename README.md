# go-filepath
* only unix-like system is supported now
* resolve filepath which is used on shell commonly but doesn't work correctly on some functions (e.g. os.OpenFile)
* easily check whether a file is existing, whether it's a directory

<br>

# examples
### resolve filepath
	path1 := "$GOPATH/bin"
	resolvedPath1, err := fpath.ResolveEnv(path1)
	// handle err
	// resolvedPath1 == "/go/bin"

	path2 := "~/.bashrc"
	resolvedPath2, err := fpath.AbsHome(path2)
	// handle err
	// resolvedPath2 == "/root/.bashrc"

### check filepath
	path := "/"
	t := fpath.GetType(path)
	t.IsFile()        // false
	t.IsDir()         // true
	t.IsExisting()    // true
	t.IsNotExisting() // false
