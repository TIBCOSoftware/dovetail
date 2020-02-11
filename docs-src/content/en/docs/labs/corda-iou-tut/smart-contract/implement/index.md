---
title: "Implement IOU Smart Contract"
linkTitle: "Implement S-Contract"
weight: 7
description: >
    In this step, we will walk you through step by step the implementation of IOU smart contract transactions: IssueIOU, TransferIOU and SettleIOU.
---

<p><video width="480" height="320" controls="controls">
    <source src="https://github.com/TIBCOSoftware/dovetail/blob/master/src/videos/iou_smart_contract_impl.mp4?raw=true" type="video/mp4">
</video></p>

### 1 IssueIOU Transaction

* Go to Apps Tab > Create to create a new Application "IOU"

* Create a flow 
    * flow name "IssueIOU" >
    * Add a trigger, select "SmartContract Action Trigger" from the list
        * Select "IOU" from the dropdown menu
        * Click "Next" and "Next" to get to "Define Transaction Input" section, security and time window are not supported for Corda smart contract
        * Add input fields
            * linearId, select "LinearId" from type dropdown menu
            * issuer, select "Party" from type dropdown menu
            * holder, select "Party" from type dropdown menu
            * amt, select "Amount<Currency>" from type dropdown menu
            * click "Next"
        * Add output fields, leave this blank, click "Continue"
        * Select "Copy Schema"
    * Select the trigger
        * Map to flow inputs
    * Now we will implement the flow
        * Add "Ledger Operation" activity from Dovetail-Ledger category
        * Map input
 

### 2 TransferIOU Transaction

* Create a new flow to implement TransferIOU transaction
    * Add a trigger, select "SmartContract Action Trigger" from the list
        * Select "IOU" from the dropdown menu
        * Click "Next" and "Next" to get to "Define Transaction Input" section, security and time window are not supported for Corda smart contract
        * Add input fields
            * iou, select "AssetRef" from Type dropdown menu, enter "com.example.iou.IOU" in the RecordType field, select "True" from Consuming dropdown menu
            * newHolder, select "Party" from type dropdown menu
            * click "Next"
        * Add output fields, leave this blank, click "Continue"
        * Select "Copy Schema"
    * Select the trigger
        * Map to flow inputs
    * Now we will implement the flow
        * Add "Ledger Operation" activity from Dovetail-Ledger category
        * Map input


### 3 SettleIOU Transaction

* Create a new flow to implement SettleIOU transaction
    * Add a trigger, select "SmartContract Action Trigger" from the list
            * Select "IOU" from the dropdown menu
            * Click "Next" and "Next" to get to "Define Transaction Input" section, security and time window are not supported for Corda smart contract
            * Add input fields
                * iou, select "AssetRef" from Type dropdown menu, enter "com.example.iou.IOU" in the RecordType field, select "True" from Consuming dropdown menu
                * funds, select "AssetRef<Cash>" from Type dropdown menu, select "True" from Consuming dropdown menu, select "True" from Repeating dropdown menu
                * click "Next"
            * Add output fields, leave this blank, click "Continue"
            * Select "Copy Schema"
        * Select the trigger
            * Map to flow inputs
        * Now we will implement the flow
            * Add "Payment Processor" activity from Dovetail-Ledger category
            * Map input
