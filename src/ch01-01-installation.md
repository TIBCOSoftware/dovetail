## Installing the Project Dovetail Studio
Installing the Project Dovetail Studio is quite simple just follow these steps:

1.- Download the latest version for your os architecture.

> You can find the latest release [here](https://github.com/TIBCOSoftware/dovetail/releases/tag/v0.1.3)

Or you can just execute this commands

```
mkdir dovetail_installation
cd dovetail_installation
curl -OL https://github.com/TIBCOSoftware/dovetail/releases/download/v0.1.3/TIB_dovetail_0.1.3_macosx_x86_64.zip
```

> Change the version or architecture according to your environment


2.- Unzip the downloaded release (for example on mac) (from dovetail_installation folder).


```
unzip TIB_dovetail_0.1.3_macosx_x86_64.zip
```

## Starting Dovetail Studio
To get started with your downloaded version of the Dovetail Studio in the previous step just do the following:

1.- Run studio

```
dovetail/0.1/bin/run-studio.sh eula-accept
```

## Launching Dovetail Studio
To launch Dovetail Studio simply open your favorite web browser, and navigate to http://localhost:8090. You'll see the initial page to create your first smart contract!


## Installation Tutorial Video

<video width="480" height="320" controls="controls">
    <source src="videos/dovetail_studio_install.mp4" type="video/mp4">
</video>