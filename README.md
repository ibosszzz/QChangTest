# QChangTest

## Install go(lang)

with [homebrew](http://mxcl.github.io/homebrew/):

```Shell
sudo brew install go
```

with [apt](http://packages.qa.debian.org/a/apt.html)-get:

```Shell
sudo apt-get install golang
```

[install Golang manually](https://golang.org/doc/install)
or
[compile it yourself](https://golang.org/doc/install/source)

## Use

### Run Service
1.Go to dir code/QChangTest
2.Run command go run main.go (listen in port 8900)

```Shell
go run main.go
```

### Run Test
Run command go test -v .\controller\

```Shell
go test -v .\controller\
```
