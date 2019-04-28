

## 3. Implement IOU Smart Contract
In this step, we will implement the 4 transactions defined in the smart contract model in Project Dovetail™ Studio. A completed iou smart contract is also available to be imported into the Studio, see step [3.5. Import an existing application] for details.

### 3.1 IssueIOU Transaction*

> * Go to Apps Tab > Create to create a new Application "IOU"

> * Create a flow > flow name "IssueIOU" > Select "SmartContract TXN Trigger" > Select "IOU" from model dropdown > Select "com.example.iou.IssueIOU" from transaction dropdown > Click "Done" button

> * We will add a logger activity, ledger activity and response activity. We will also validate input before writing to ledger. 

<p><video width="480" height="320" controls="controls">
    <source src="videos/issueiou.mp4" type="video/mp4">
</video></p>

### 3.2 TransferIOU Transaction*

> * Create a new flow to implement TransferIOU transaction
<p><video width="480" height="320" controls="controls">
    <source src="videos/transferiou.mp4" type="video/mp4">
</video></p>

### 3.3 getIOU Transaction*

> * Create a new flow to implement getIOU transaction
<p><video width="480" height="320" controls="controls">
    <source src="videos/getiou.mp4" type="video/mp4">
</video></p>

### 3.4 getIOUIssuedBy Transaction*

> * Create a new flow to implement getIOUIssuedBy transaction
> * We use "Custom Query" activity to query the ledger. This activity is **NOT blockchain agnostic**
   > - define a input parameter "issuerId"
   > - define blochchain specific query string, use _$paramName for input substitution, in this example, we use [Hyperledger Fabric CouchDB query syntax](https://hyperledger-fabric.readthedocs.io/en/release-1.3/couchdb_tutorial.html)
   *** During the mapping of transaction response, you will notice tht "array." is underlined with a red line, it is a UI validation error because Project Dovetail™ Studio UI does not support function yet, the error has no impact to runtime. Function will be supported in the very near future.

<p><video width="480" height="320" controls="controls">
    <source src="videos/getiouissuedby.mp4" type="video/mp4">
</video></p>

### 3.5. Import an existing application

> * You can also import an application into Project Dovetail™ Studio
    > * Go to Project Dovetail™ Studio
    > * If you haven't import iou.bna file from Connections tab, you can do so now following [this step](ch02-03-import.md)
    > * Create a new Application "IOUImport", click on "Import" button, use this [IOU.json](tutorials/IOU.json)

<p><video width="480" height="320" controls="controls">
    <source src="videos/importiou.mp4" type="video/mp4">
</video></p>