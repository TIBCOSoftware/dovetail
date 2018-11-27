---
title:  A Simple IOU Smart Contract Tutorial
hidden: true
---

In this tutorial, we will walk you through the steps to model, implement and test smart contracts on distributed ledger technology platform of your choice. The example is a simple "I owe you" use case, the issuer of an IOU is obligated to pay the owner of the IOU amount issued, the ownership of the IOU can be transferred by current owner, and all transactions will be recorded on the ledger.

Before getting started, you should have [Project Dovetail™ Studio](../../getting-started/getting-started-webui) and [Project Dovetail™ CLI](../../getting-started/getting-started-cli) installed, and your [development environment setup](../../development/devenv).

******

## 1. Model IOU Smart Contract
### *1.1 Introduction to smart contract modeling*
Project Dovetail™ supports [Hyperledger Composer modeling language](https://hyperledger.github.io/composer/v0.19/reference/cto_language.html) to model smart contract assets and transactions, please refer to [Composer Connector](https://github.com/TIBCOSoftware/dovetail-contrib/tree/master/SmartContract/connector/composer) for support detail

### *1.2 IOU Smart Contract*
We will use Visio Studio Code to create IOU smart contract model.

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

``` 
> cd /path/to/iou
> zip -r ../iou.bna *
```


## 2. Import IOU Smart Contract Model

* Start up Project Dovetail™ Studio, it will display web URL to use
* Go to a browser and open the Web Studio
* Go to Connections Tab > Add Connection > Select Composer Connector
* Give the connection a name, then "Browse" to find the iou.bna file, click "Save Model" button. 

<p><a target="_blank" rel="noopener noreferrer" href="../images/recordings/importbna.gif"><img src="../images/recordings/importbna.gif" alt="Import Model" style="max-width:100%;"></a></p>

## 3. Implement IOU Smart Contract
### *3.1 IssueIOU Transaction*

* Go to Apps Tab > Create to create a new Application "IOU"
* Create a flow > flow name "IssueIOU" > Select "SmartContract TXN Trigger" > Select "IOU" from model dropdown > Select "com.example.iou.IssueIOU" from transaction dropdown > Click "Done" button
* We will add a logger activity, ledger activity and response activity. We will also validate input before writing to ledger. 

<p><a target="_blank" rel="noopener noreferrer" href="../images/recordings/issueiou.gif"><img src="../images/recordings/issueiou.gif" alt="Issue IOU" style="max-width:100%;"></a></p>

### *3.2 TransferIOU Transaction*

* Create a new flow to implement TransferIOU transaction
<p><a target="_blank" rel="noopener noreferrer" href="../images/recordings/transferiou.gif"><img src="../images/recordings/transferiou.gif" alt="Transfer IOU" style="max-width:100%;"></a></p>

### *3.3 getIOU Transaction*

* Create a new flow to implement getIOU transaction
<p><a target="_blank" rel="noopener noreferrer" href="../images/recordings/getiou.gif"><img src="../images/recordings/getiou.gif" alt="Get IOU" style="max-width:100%;"></a></p>

### *3.4 getIOUIssuedBy Transaction*

* Create a new flow to implement getIOUIssuedBy transaction
* We use "Custom Query" activity to query the ledger. This activity is **NOT blockchain agnostic**
   - define a input parameter "issuerId"
   - define blochchain specific query string, use _$paramName for input substitution, in this example, we use [Hyperledger Fabric CouchDB query syntax](https://hyperledger-fabric.readthedocs.io/en/release-1.3/couchdb_tutorial.html)
   *** During the mapping of transaction response, you will notice tht "array." is underlined with a red line, it is a UI validation error because Project Dovetail™ Studio UI does not support function yet, the error has no impact to runtime. Function will be supported in the very near future.

<p><a target="_blank" rel="noopener noreferrer" href="../images/recordings/getiouissuedby.gif"><img src="../images/recordings/getiouissuedby.gif" alt="Get IOU Issued by" style="max-width:100%;"></a></p>

## 4. Export IOU Smart Contract Application

* Go to IOU application
* Click on "Export app" button, save the file IOU.json to artifacts folder

## 5. Generate Blockchain Artifacts

### *5.1 Hyperledger Fabric*
Run following command from tutorial directory Hyperledger Fabric chaincode is written to tutorial/artifacts/hlf folder

For testing, transaction security support is not enabled.

> /path/to/dovetail-cli contract generate -b fabric -m artifacts/IOU.json -v 1.0.0 -t artifacts/

### *5.2 R3 Corda*
Run folloowing command from tutorial directory, R3 Corda contract is written to tutorial/artifacts/corda folder, put kotlin-IOU-1.0.0.jar on your classpath to develop your CorDapp. A sample IOUApp has been created and availble [here](https://github.com/TIBCOSoftware/dovetail/blob/master/docs/content/labs/network/corda)

> /path/to/dovetail-cli contract generate -b corda -m artifacts/IOU.json -v 1.0.0 -t artifacts/ --namespace com.example.iou

## 6. Test IOU Smart Contract
### *6.1 Hyperledger Fabric*

* Copy [sample Hyperledger Fabric Network](https://github.com/TIBCOSoftware/dovetail/blob/master/docs/content/labs/network/fabric) to your network/fabric directory
* Follow [instructions](https://github.com/TIBCOSoftware/dovetail/blob/master/docs/content/labs/network/fabric/fabric_network.md), assuing you are running commands from fabric directory.

### *6.2 R3 Corda*

* Copy [sample R3 CorDapp](https://github.com/TIBCOSoftware/dovetail/blob/master/docs/content/labs/network/corda) to network/corda directory
* Follow [instructions](https://github.com/TIBCOSoftware/dovetail/blob/master/docs/content/labs/network/corda/corda_network.md)

## 7. Import an existing application

* You can import an application into Project Dovetail™ Studio
   * Go to Project Dovetail™ Studio
   * If you haven't import iou.bna file from Connections tab, you can do so now following step 2 in this document
   * Create a new Application "IOUImport", click on "Import" button, use this [IOU.json](https://github.com/TIBCOSoftware/dovetail/blob/master/docs/content/labs/artifacts/IOU.json)

<p><a target="_blank" rel="noopener noreferrer" href="../images/recordings/importiou.gif"><img src="../images/recordings/importiou.gif" alt="Get IOU Issued by" style="max-width:100%;"></a></p>


