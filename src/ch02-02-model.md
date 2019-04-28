## 1. Model IOU Smart Contract

### 1.1 Introduction to smart contract modelling

In this section we will show step by step instructions to create the iou.bna that we will use in the next section of the tutorial.

> You can skip the instructions and download the iou.bna file [here](tutorials/iou/iou.bna).

Project Dovetail™ supports [Hyperledger Composer modeling language](https://hyperledger.github.io/composer/v0.19/reference/cto_language.html) to model smart contract assets and transactions, please refer to [Composer Connector](https://github.com/TIBCOSoftware/dovetail-contrib/tree/master/SmartContract/connector/composer) for for more detail.

### 1.2 IOU Smart Contract Data Model
We will use Visual Studio Code to create IOU smart contract model.

>*Create the top level folder called "iou"

```bash
mkdir iou
```

>*Copy following package metadata and save it as package.json in your iou folder

```json
{
  "engines": {
    "composer": "^0.19.0"
  },
  "name": "iou",
  "version": "0.0.1",
  "description": "IOU network"
}
```

>*Create a models folder called "models" inside the top level "iou" folder

```bash
cd iou
mkdir models
```

>*Copy following resource definitions and save it as iou.cto in your models folder

```json
namespace com.example.iou

import  com.tibco.dovetail.system.*

asset  IOU identified by linearId extends LinearState {
  --> Party issuer
  --> Party owner
  o Amount amt
}

/*
@InitiatedBy
  - arg0: comma delimited list of authorized participants in the format of $tx.path.to.participant, or * for any participant
  - arg1: optional, comma delimited list of required attributes that must exist in the initiator's certificate in the format of name=value.
*/
@InitiatedBy("$tx.iou.issuer")
transaction IssueIOU {
  o IOU iou
}

@InitiatedBy("$tx.iou.owner")
transaction TransferIOU {
  --> IOU iou //by reference because the asset is already on the ledger
  --> Party newOwner
}

/* 
@Query : mark transaction as query only
*/
@Query()
transaction getIOU {
  o String linearId
}

  @Query()
transaction getIOUIssuedBy {
  o String issuerPartyId
}
```

>*Copy following resource definitions and save it as dovetail.system.cto in your models folder

```json
namespace com.tibco.dovetail.system

/*
for Hyperledger Fabric, id is the member's mspId
for Corda, id is the member's cert name
*/
participant Party identified by id {
  o String id
}

/*
All non-fungible asset should inherit from LinearState
*/
@CordaClass("net.corda.core.contracts.LinearState")
abstract asset LinearState {
  o String linearId
}

@CordaClass("net.corda.core.contracts.OwnableState")
abstract asset OwnableState {
  o String linearId
  -->Party owner
}

@CordaClass("net.corda.core.contracts.FungibleAsset")
abstract asset FungibleAsset {
  -->Party owner
  o Amount amt
}

@CordaClass("net.corda.finance.contracts.asset.Cash.State")
asset Cash identified by assetId extends FungibleAsset {  
  o String assetId
}

@CordaClass("net.corda.core.contracts.Amount<Currency>")
concept Amount {
  o String currency
  o Long quantity default = 0
}

@CordaClass("net.corda.core.contracts.Amount<Issue<Currency>>")
concept IssueAmount {
  -->Party issuer
  o String currency
  o Long quantity default = 0
}
```

>*Run zip command from iou directory to package the iou project, it will create a iou.bna file in the tutorial folder

```bash
zip -r ../iou.bna *
```
