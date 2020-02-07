---
title: "Corda Distributed Application"
linkTitle: "Corda Application"
weight: 2
description: >
  Development of a Corda Distributed Application
---

### 1 Parties and Roles for the network
For this tutorial, the story line is party "charlie" issues an IOU to party "alice", party "alice" then transfers the IOU to party "bob", finally party "charlie" settles the IOU with party "bob". party "bank" is the cash issuer.

Below the are the dapp flows each party should implement:

* charlie is the IOU issuer
    * Implement IssueIOU initiating flow
    * Implememt TransferIOU receiving flow
    * Implement SettleIOU initiating flow
* alice is the original IOU holder, and will transfer the IOU to bob
    * Implement IssueIOU receiving flow
    * Implement TransferIOU initiating flow
* bob is the new IOU holder
    * Implement TransferIOU receiving flow
    * Implement SettleIOU receiving flow
* bank is the cash issuer
    * Will use Corda CashIssueAndPayment flow to issue cash to charlie

### 2. Import Smart Contract
Before we start implementing the flows, Dovetail needs to know the smart contract transactions and the input parameters of each transaction.

* Go to Connection tab 
    * Add Connection 
    * Select "Import Dovetail Contract"
    * Enter name "IOUContract"
    * Browse and select artifacts/iou.json file
    * Click "Done"
