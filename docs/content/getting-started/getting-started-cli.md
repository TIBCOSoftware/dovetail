---
title: Dovetail CLI
weight: 2030
pre: "<i class=\"fas fa-terminal\" aria-hidden=\"true\"></i> "
---
### Introduction
The dovetail cli is a tool to mainly generate smart contracts for a given model built using ui tool [Dovetail Studio](../getting-started-webui), so we recommend you to learn how to do that first.

### Before you get started
Before you can get started with the cli tools you need to make sure you have the [Go programming language](https://golang.org/doc/install) and [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) installed. Go v1.11.x is required.

{{% notice info %}}
Don't forget to set your `GOPATH` variable and make sure that `$GOPATH/bin` is part of your path. (see [here](https://golang.org/doc/code.html#GOPATH) or [here](https://github.com/golang/go/wiki/SettingGOPATH) for more details)
{{% /notice %}}

### Installing the cli tools

Copy and paste the following commands to install Project Dovetail™ commandline tool.

```
mkdir dovetail-cli && cd dovetail-cli && export GOPATH=${PWD} && go get -u github.com/TIBCOSoftware/dovetail-cli/... && cd $GOPATH/src/github.com/TIBCOSoftware && cd dovetail-cli && go1MODULE=on go install ./... && cd .. && rm -rf flogo-lib/ && rm -rf flogo-contrib/ && git clone https://github.com/TIBCOSoftware/flogo-contrib.git && git clone https://github.com/TIBCOSoftware/flogo-lib.git && cd flogo-contrib && git checkout tags/v0.5.5 && cd .. && cd flogo-lib && git checkout tags/v0.5.5 && cd .. && git clone https://github.com/TIBCOSoftware/dovetail-contrib.git && cd $GOPATH/bin && go get -u github.com/jteeuwen/go-bindata/... && go get -u github.com/kardianos/govendor && go get -u github.com/Sirupsen/logrus && go get -u github.com/julienschmidt/httprouter 

```

**The binary dovetail will be in the dovetail-cli/bin directory, please prepend the /path/to/dovetail-cli to your GOPATH environment variable, and prepend /path/to/dovetail-cli/bin to your PATH environment variable in your user profile**
