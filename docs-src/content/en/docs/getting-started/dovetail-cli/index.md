---
title: "Dovetail CLI"
linkTitle: "Dovetail CLI"
weight: 2
description: >
  generate smart contracts for a given model
---

## Introduction
The dovetail cli is a tool to mainly generate smart contracts for a given model built using ui tool [Dovetail Studio](ch01-01-installation.md), so we recommend you to learn how to do that first.

## Before you get started
Before you can get started with the cli tools you need to make sure you have the [Go programming language](https://golang.org/doc/install) and [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) installed. Go v1.13.x is required.

> Don't forget to set your `GOPATH` variable and make sure that `$GOPATH/bin` is part of your `PATH`
>(see [here](https://golang.org/doc/code.html#GOPATH) or [here](https://github.com/golang/go/wiki/SettingGOPATH) for more details)

## Installing the cli tools

Copy and paste the following commands to install Project Dovetail™ commandline tool.

```
curl https://github.com/TIBCOSoftware/dovetail-cli/releases/download/v0.2.0/dovetail-cli-install.sh -sSfL v0.2.0 | sh
```

>The binary dovetail will be in your dovetail-cli/bin directory.
>You can add /path/to/dovetail-cli/bin/dovetail to your PATH env variable for easier access. 


## Installation Tutorial Video

<video width="480" height="320" controls="controls">
    <source src="/videos/dovetail_cli_install.mp4" type="video/mp4">
</video>
