---
title: "CorDapp for alice"
linkTitle: "Flows for alice"
weight: 2
description: >
  Step by step instructions to create distributed application "alice" for party alice.
---

<p><video width="480" height="320" controls="controls">
    <source src="videos/alice.mp4" type="video/mp4">
</video></p>

### 1. IssueIOU receiver flow
* Create a flow 
   * Create a flow 
   * flow name: IssueIOUResponder
   * select trigger : Dovetail CorDApp Flow Receiver
      * select receiver from flow type dropdown
      * select false for "Use confidential identities for this transaction?"
      * initiator flow name: com.charlie.iou.flows.IssueIOUInitiator
      * click "Continue" button
      * select "Copy Schema"
      * select the trigger, and map flow input

### 2. TransferIOU initiator flow

 * flow name: TransferIOUInitiator
   * add a trigger : select "Dovetail CorDApp Flow Initiator" from the list
      * select false for "Use confidential identities for this transaction?"
      * select false for "Send transactions to observers?"
      * click "Next" button
      * add following flow input parameters
         * iouId :    Type = LinearId
         * newHolder: Type = Party, PartyRole = Participant
      * click "Continue" button
      * select "Copy Schema"
      * select the trigger, and map flow input
   * implement TransferIOUInitiator flow
      * add SimpleVaultQuery activity from Dovetail-CorDApp category
        * Configuration screen
            * select IOU from asset dropdow
            * Map input
      * if IOU is found
        * add BuildTransactoinProposal activity from Dovetail-CorDApp category
            * select "IOUContract:" from contract dropdown
            * select com.example.iou.TransferIOU from transaction dropdown
            * map the activity input
      * otherwise throw error

### 3. Export the "alice" applicatio
export the application to artifiacts/alice.json


