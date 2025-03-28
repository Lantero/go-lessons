# go-lessons

Welcome to my personal Golang playground.

## 1. How to install or upgrade go

1. Follow [official instructions](https://go.dev/doc/install) to remove old installation, then
   download and unpack the new one.

2. Make sure your `PATH` contains:

   - `/usr/local/go/bin` -> Go binaries themselves i.e. `go` and `gofmt`.

   - `/home/$USER/go` -> CLI tools or packages you might install via `go install` from your local
     user i.e. `bin` and `pkg` directories.

## 2. Visual Studio Code support

Just install the official extension:

- [Marketplace](https://marketplace.visualstudio.com/items?itemName=golang.Go)

- [Repository](https://github.com/golang/vscode-go)

## 3. How to manage a go project

### 3.1. Create a new project

1. Create a folder in your repository, like `mkdir project`.

2. Once inside that directory, run `go mod init github.com/lantero/go-lessons/project`.
