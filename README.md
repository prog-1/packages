# Packages

## What is it

A package is a directory inside Go workspace containing one or more Go source files and/or other packages.

Every Go language source file in this directory belongs to the package. All Go source files should have the
same `package` starting every file. E.g. `package main` or `package mypackage`. If these rules are met, Go
CLI front-end application `go` will handle these packages auto-magically.

Packages allow to group functionality, resolve naming conflicts e.g. having functions with the same name,
but different meaning (or even the same name and the same meaning, but different implementations e.g.
packages`text/template` and `html/template`; or `strings` and `bytes`).

## Importing packages

You are already familiar with how packages can be imported. Example:

```go
import "fmt"

import (
  "encoding/json"
  "math"
)
```

### Export/import names

Symbols (e.g. global variable names, functions, contants, types, etc) that start with capital letters are available
to be imported e.g. `fmt.Println` or `strings.Contains`. Other names are available from inside of the package only.

### Import paths

* Reference: https://go.dev/doc/go1.4#canonicalimports
* Remote import paths: https://pkg.go.dev/cmd/go#hdr-Remote_import_paths

Code often lives in repositories hosted by public services such as github.com, meaning that the import paths
for packages begin with the name of the hosting service, github.com/rsc/pdf for example. One can use an existing
mechanism to provide a "custom" or "vanity" import path such as rsc.io/pdf, but that creates two valid import paths
for the package. 

## Special packages

### Main

The only special package is `main`. If package must contain `main` function, which is your application
entry point. Example:

```go
package main
import "fmt"
func main() { fmt.Println("Hello, world!") }
```

There can be multiple applications in the same workspace, and each of them must be contained in a separate
directory. In the same way we always handled our examples - each example in its own directory.

### Internal

* Reference: https://go.dev/doc/go1.4#internalpackages

Go's package system makes it easy to structure programs into components with clean boundaries, but there
are only two forms of access: local (unexported) and global (exported). Sometimes one wishes to have components
that are not exported, for instance to avoid acquiring clients of interfaces to code that is part of a public
repository but not intended for use outside the program to which it belongs.

The Go language does not have the power to enforce this distinction, but as of Go 1.4 the go command introduces
a mechanism to define "internal" packages that may not be imported by packages outside the source subtree in which they reside.

To create such a package, place it in a directory named internal or in a subdirectory of a directory named internal.
When the go command sees an import of a package with internal in its path, it verifies that the package doing the
import is within the tree rooted at the parent of the internal directory. For example, a package .../a/b/c/internal/d/e/f
can be imported only by code in the directory tree rooted at .../a/b/c. It cannot be imported by code in .../a/b/g or in
any other repository.

## Modules

In Go a module is a directory containing a collected of related packages (packages can be nested). The module
contain `go.mod` file in its root directory. `go.mod` file provides a dependency managing capabilities,
utilizing information about the mobule import path, minimum required Go version and depenedent modules.

`go.mod` files are managed by `go mod` command line:

```text
$ go help mod
Go mod provides access to operations on modules.
...
Usage:
        go mod <command> [arguments]
The commands are:

        download    download modules to local cache
        edit        edit go.mod from tools or scripts
        init        initialize new module in current directory
        ...
```

### Adding remote dependencies

To add a remote dependency use `go get` command e.g. `go get github.com/yarcat/fsm-go` or
`go get github.com/nsf/termbox-go`. Use `go help mod` for more information.

Running `go mod tidy` peridically allows to keep the dependencies under control by removing unused deps, and
hashing (remembering exact revision) used ones.
