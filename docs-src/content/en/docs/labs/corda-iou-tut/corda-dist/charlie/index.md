---
title: "CorDapp for charlie"
linkTitle: "Flows for charlie"
weight: 1
description: >
  Step by step instructions to create distributed application "charlie" for party charlie.
---

<p><video width="480" height="320" controls="controls">
    <source src="videos/charlie.mp4" type="video/mp4">
</video></p>

### 1. IssueIOU initiator flow
* Create a flow 
   * flow name: IssueIOUInitiator
   * add a trigger : select "Dovetail CorDApp Flow Initiator" from the list
      * select false for "Use confidential identities for this transaction?"
      * select false for "Send transactions to observers?"
      * click "Next" button
      * add following flow input parameters
         * holder :    Type = Party, PartyRole = Participant
         * amt:        Type = Amount<Currency> 
         * extId:      Type = String
      * click "Continue" button
      * select "Copy Schema"
      * select the trigger, and map flow input
   * implement IssueIOUInitiator flow
      * select BuildTransactoinProposal activity from Dovetail-CorDApp category
      * select "IOUContract:" from contract dropdown
      * select com.example.iou.IssueIOU from transaction dropdown
      * map the activity input
         * use cordapp.createLinearIdFromExternalId($flow.transactionInput.extId) to map iou.linearId
         > * You will see an error at design time for this function, ignore the error for now, it is supported at runtime


### 2. TransferIOU receiver flow

 * Create a flow 
   * flow name: TransferIOUResponder
   * select trigger : Dovetail CorDApp Flow Receiver
      * select receiver from flow type dropdown
      * select false for "Use confidential identities for this transaction?"
      * initiator flow name: com.alice.iou.flows.TransferIOUInitiator
      * click "Continue" button
      * select "Copy Schema"
      * select the trigger, and map flow input


### 3. SettleIOU initiator flow
 * Create a flow 
   * flow name: SettleIOUInitiator
   * select trigger : Dovetail CorDApp Flow Initiator
      * select false for "Use confidential identities for this transaction?"
      * select false for "Send transactions to observers?"
      * click "Next" button
      * add following flow input parameters
         * iouId : Type = LinearId
      * click "Continue" button
      * select "Copy Schema"
      * select the trigger, and map flow input
   * Implement SettleIOUInitiator flow
      * add SimpleVaultQuery activity from Dovetail-CorDApp category
        * Configuration screen
            * select IOU from asset dropdow
            * Map input
      * if IOU is found
            * add CashWallet activity from Dovetail-CorDApp category
                * Select "Retrieve Funds" from dropdown
                * Map input
            * add BuildTransactoinProposal Dovetail-CorDApp category
                * select IOU from contract dropdown
                * select com.example.iou.SettleIOU from transaction dropdown
                * map the activity input
      * otherwise throw error

### 4. Export the "charlie" applicatio
export the application to artifiacts/charlie.json



