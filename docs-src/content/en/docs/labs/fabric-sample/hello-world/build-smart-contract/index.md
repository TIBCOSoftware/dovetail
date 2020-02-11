---
title: "Build smart contract"
linkTitle: "Build smart contract"
weight: 8
description: >
  Instructions on how to build a hyperledger smart contract built in TCI (TIBCO Cloud Integration)
---

Before anything download the fabric_hw.json and fabric_hw_client.json from the [create smart contract](../create-smart-contract) section.

Also make sure the [environment prerequisites](../../../../getting-started/environment-prerequisites) for hyperledger fabric have been installed.


#### Clone dovetail-contrib repository

```bash
git clone https://github.com/TIBCOSoftware/dovetail-contrib.git
```

#### Build using fabric-scripts

```bash
cd dovetail-contrib/hyperledger-fabric/fabric-scripts/
```

Replace fabric_hw.json with full path to your json file

```bash
make APP_FILE=/path/to/fabric_hw.json APP_NAME=fabric_hw create
```

```bash
make APP_NAME=fabric_hw build
```

After you follow these steps, your smart contract chaincode can be found under your current folder in a fabric_hw folder.






