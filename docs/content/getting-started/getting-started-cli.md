---
title: Flogo CLI
weight: 2030
pre: "<i class=\"fas fa-terminal\" aria-hidden=\"true\"></i> "
---

### Before you get started
Before you can get started with the cli tools you need to make sure you the [Go programming language](https://golang.org/doc/install) installed. 

{{% notice info %}}
Don't forget to set your `GOPATH` variable and make sure that `$GOPATH/bin` is part of your path. (see [here](https://golang.org/doc/code.html#GOPATH) or [here](https://github.com/golang/go/wiki/SettingGOPATH) for more details)
{{% /notice %}}

Also make sure that you have installed [dep](https://github.com/golang/dep#setup) dependency management tool.

### Installing the cli tools
Now that you've installed the Go programming language and dep there are a few commands you can run to install the cli and make developing with the cli tools even easier

* First you'll need to **go get** flogo by running `go get -u github.com/TIBCOSoftware/dovetail-cli/...`. This will get you both the CLI tools.
* Second you can build the binary by running `go install ./...`.

{{% notice note %}}
If you want to update the CLI tools, you can run `go get -u github.com/TIBCOSoftware/flogo-cli/...` to get the latest version and repeat dep ensure step. 
{{% /notice %}}
