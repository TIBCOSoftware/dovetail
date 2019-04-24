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
cd $GOPATH
curl https://github.com/TIBCOSoftware/dovetail-cli/releases/download/v0.1.3/dovetail-cli-install.sh -sSf | sh
```

**The binary dovetail will be in the dovetail-cli/bin directory, please prepend the /path/to/dovetail to your GOPATH environment variable, and prepend /path/to/dovetail/bin to your PATH environment variable in your user profile**


### Installation Tutorial Video

<video width="480" height="320" controls="controls">
    <source src="../../labs/videos/dovetail_cli_install.mp4" type="video/mp4">
</video>
