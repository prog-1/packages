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

## Special packages

The only special package is `main`. If package must contain `main` function, which is your application
entry point. Example:

```go
package main
import "fmt"
func main() { fmt.Println("Hello, world!") }
```

There can be multiple applications in the same workspace, and each of them must be contained in a separate
directory. In the same way we always handled our examples - each example in its own directory.

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
