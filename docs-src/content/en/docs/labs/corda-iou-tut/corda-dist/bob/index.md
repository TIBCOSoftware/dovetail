---
title: "CorDapp for bob"
linkTitle: "Flows for bob"
weight: 3
description: >
  Step by step instructions to create distributed application "bob" for party bob.
---

<p><video width="480" height="320" controls="controls">
    <source src="videos/bob.mp4" type="video/mp4">
</video></p>

### 1. TransferIOU receiver flow
* Create a flow 
   * Create a flow 
   * flow name: TransferIOUResponder
   * select trigger : Dovetail CorDApp Flow Receiver
      * select receiver from flow type dropdown
      * select false for "Use confidential identities for this transaction?"
      * initiator flow name: com.alice.iou.flows.TransferIOUInitiator
      * click "Continue" button
      * select "Copy Schema"
      * select the trigger, and map flow input

### 2. SettleIOU receiver flow
* Create a flow 
   * Create a flow 
   * flow name: SettleIOUResponder
   * select trigger : Dovetail CorDApp Flow Receiver
      * select receiver from flow type dropdown
      * select false for "Use confidential identities for this transaction?"
      * initiator flow name: com.charlie.iou.flows.SettleIOUInitiator
      * click "Continue" button
      * select "Copy Schema"
      * select the trigger, and map flow input

### 3. Export the "bob" applicatio
export the application to artifiacts/bob.json