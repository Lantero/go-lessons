# go-lessons

My own golang playground

## How to install/upgrade go

1. Uninstall existing version: `sudo rm -rf /usr/local/go`.

2. Download TGZ from https://go.dev/dl/.

3. Extract new version: `sudo tar -C /usr/local -xzf /home/carlos/Downloads/go1.22.1.linux-amd64.tar.gz`.

4. Make sure your `PATH` contains `/usr/local/go/bin` and `/home/carlos/go`:
   - `echo $PATH | grep "/usr/local/go/bin"`
   - `echo $PATH | grep "$(go env GOPATH)/bin"`

## Visual Studio Code support

![image](https://github.com/Lantero/go-lessons/assets/3510723/e5cfdd2d-5aa5-40c5-be05-2e5d2d30b08a)
