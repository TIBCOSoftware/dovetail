## 1. Model IOU Smart Contract

### *1.1 Introduction to smart contract modeling*
Project Dovetail™ supports [Hyperledger Composer modeling language](https://hyperledger.github.io/composer/v0.19/reference/cto_language.html) to model smart contract assets and transactions, please refer to [Composer Connector](https://github.com/TIBCOSoftware/dovetail-contrib/tree/master/SmartContract/connector/composer) for support detail

### *1.2 IOU Smart Contract*
We will use Visual Studio Code to create IOU smart contract model.

* create a workspace folder, e.g. tutorial
* create subfolders under tutorial
     * artifacts
     * network
        * fabric
        * corda
* copy [template project](https://github.com/TIBCOSoftware/dovetail/tree/master/docs/content/labs/artifacts/composer-project-template) to the workspace, and rename the project as iou
* under folder iou/model, create a file iou.cto
* copy following resource definitions into iou.cto file

```
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
* Run zip command from iou directory to package the iou project, it will create a iou.bna file in the tutorial folder

> cd /path/to/iou

> zip -r ../iou.bna *
