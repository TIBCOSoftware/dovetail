---
title: "Installation"
linkTitle: "Installation"
weight: 1
description: >
  Installing the Project Dovetail Studio is quite simple
---

## Installing the Project Dovetail Studio
just follow these steps:

1.- Download the latest version for your os architecture and your blockchain target.

> You can find the latest release [here](https://github.com/TIBCOSoftware/dovetail/releases/latest)

Or you can just execute this commands


> For Hyperledger fabric development:

```
mkdir dovetail_installation
cd dovetail_installation
curl -OL https://github.com/TIBCOSoftware/dovetail/releases/download/v0.2.0/TIB_dovetail-fabric_0.2.0_macosx_x86_64.zip
```

> For R3 Corda development:

```
mkdir dovetail_installation
cd dovetail_installation
curl -OL https://github.com/TIBCOSoftware/dovetail/releases/download/v0.2.0/TIB_dovetail-corda_0.2.0_macosx_x86_64.zip
```

> For Multitarget development:

```
mkdir dovetail_installation
cd dovetail_installation
curl -OL https://github.com/TIBCOSoftware/dovetail/releases/download/v0.2.0/TIB_dovetail-multitarget_0.2.0_macosx_x86_64.zip
```



2.- Unzip the downloaded release (for example on mac) (from dovetail_installation folder).

```
unzip TIB_dovetail-fabric_0.2.0_macosx_x86_64.zip
```

## Starting Dovetail Studio
To get started with your downloaded version of the Dovetail Studio in the previous step just do the following:

1.- Run studio

```
dovetail/0.2/bin/run-studio.sh eula-accept
```

## Launching Dovetail Studio
To launch Dovetail Studio simply open your favorite web browser, and navigate to http://localhost:8090. You'll see the initial page to create your first smart contract!


## Installation Tutorial Video

<video width="480" height="320" controls="controls">
    <source src="https://github.com/TIBCOSoftware/dovetail/blob/master/src/videos/dovetail_studio_install.mp4?raw=true" type="video/mp4">
</video>