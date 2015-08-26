
### go-sloc: Count SLOC for Go/golang code

By Andrew Stewart ([http://andrewinfosec.com](http://andrewinfosec.com))

A simple program to count the
[SLOC](https://en.wikipedia.org/wiki/Source_lines_of_code) for
[Go](http://golang.org) code.

#### Example

    $ find . -name *.go ! -name '*_test.go' | go-sloc
    213

