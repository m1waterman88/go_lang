# Learning Go

## Some Go Commands

See [How to Write Go Code](https://golang.org/doc/code.html).

### Compile and run (no binaray produced)

* Good for quickly testing single-file programs  
`go run file_name.go`

### Build the package and keep it in the current working directory

`go build file_name.go`

### Build the package and install the binary in $GOPATH/pkg

* The real deal: compile and keep the executable  
`go install file_name.go`

### Execute the binaray

* Typical file execution  
`./file_name`

