
### go-sloc: Count SLOC for Go code

By Andrew Stewart ([https://andrewinfosec.com](https://andrewinfosec.com))

A simple program to count the
[SLOC](https://en.wikipedia.org/wiki/Source_lines_of_code) for
[Go](http://golang.org) code.

#### Example usage

    $ go run go-sloc.go < go-sloc.go 
    51

    $ find . -name '*.go' ! -name '*_test.go' -exec cat {} \; | go-sloc
    423

