# goinit

`goinit` is a simple command line tool to generate new golang project for GoLand.

## Usage

Installation:

```sh
$ go install github.com/ajinasokan/goinit
```

Usage:

```sh
$ goinit github.com/ajinasokan/helloworld
creating helloworld
$ cd helloworld
```

It does the following things:

* Creates go.mod with given repo URL
* Creates a hello world example
* Creates IntelliJ project
* Creates new build configuration
* Enables `Single instance only` mode
* Sets GOROOT in the project using `runtime.GOROOT()`
* Enables `goimports` in the project using `File Watchers`