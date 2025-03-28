# go-lessons

Welcome to my personal Golang playground.

Check [INDEX.md](INDEX.md) for a full list of all the go code examples. You can rebuild this index
by simply running `make INDEX.md`, descriptions would be taken from `go doc <name>`.

## 1. How to install or upgrade go

1. Follow [official instructions](https://go.dev/doc/install) to remove old installation, then
   download and unpack the new one.

2. Make sure your `PATH` contains:

   - `/usr/local/go/bin` -> Go binaries themselves i.e. `go` and `gofmt`.

   - `$(go env GOPATH)/bin` -> Binaries you might compile and install via `go install`.

## 2. Visual Studio Code support

Just install the official extension:

- [Marketplace](https://marketplace.visualstudio.com/items?itemName=golang.Go)

- [Repository](https://github.com/golang/vscode-go)

## 3. How to manage a go project

### 3.1. Create a new project

1. Create a folder in your repository, like `mkdir <name>`.

2. Once inside that directory, run `go mod init github.com/lantero/go-lessons/<name>`.

### 3.2. Compile your code

1. From the directory that contains your `go.mod`, run `go install`.

2. The binary will be in `$(go env GOPATH)/bin` with the name of your project, ready to be run.

3. You can check the documentation with `go doc <name>`.
