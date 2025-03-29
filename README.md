# go-lessons

Welcome to my personal Golang playground.

This `README.md` file contains some basic tips to manage go projects. `examples/` has a range of
practical use cases that cover the features of the language in a progressive way.

Check [INDEX.md](INDEX.md) for a full list of all the go code examples. You can rebuild this index
by simply running `make INDEX.md`, and descriptions would be taken from the output of
`go doc <name>`. You can install all the examples with `make install`, and revert with
`make uninstall`.

Kudos to
[Johanna Ratliff's learning path](https://gist.github.com/johannaratliff/364f638a1ca7a9dbe3871f876205d676)
for helping with practical tips on efficient Golang learning. I essentially followed the path laid
down there.

## 1. How to install or upgrade go

1. Follow [official instructions](https://go.dev/doc/install) to remove old installation, then
   download and unpack the new one.

2. Make sure your `PATH` contains:

   - `/usr/local/go/bin` -> Go binaries themselves i.e. `go` and `gofmt`.

   - `$(go env GOPATH)/bin` -> Binaries you might compile and install via `go install`.

### 1.1. Visual Studio Code support

Just install the official extension:

- [Marketplace](https://marketplace.visualstudio.com/items?itemName=golang.Go)

- [Repository](https://github.com/golang/vscode-go)

## 2. How to manage a go project

### 2.1. Create a new project

1. Create a folder in your repository. Taking this repo as example:
   `mkdir examples/go-lessons-<number>`.

2. Once inside that directory, run
   `go mod init github.com/lantero/go-lessons/examples/go-lessons-<number>`.

### 2.2. Compile your code

1. From the directory that contains your `go.mod`, run `go install`.

2. The binary will be in `$(go env GOPATH)/bin/go-lessons-<number>`.

3. You can check the self-documentation of your project with `go doc go-lessons-<number>`.
