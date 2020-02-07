---
title: "Corda IOU Smart Contract Tutorial"
linkTitle: "Corda IOU Tutorial"
weight: 5
description: >
  model and implement smart contracts and dapps for R3 Corda using Dovetail Studio
---

## Corda IOU Smart Contract Tutorial

In this tutorial, we will walk you through the steps to model and implement smart contracts and dapps for R3 Corda using Dovetail Studio, then we will use Dovetail CLI to generate runtime artifacts, and also generate RPC client with REST API interfaces for external integration, finally we will use Swagger UI to run some test cases.

The example is a simple "I owe you" use case, the issuer of an IOU is obligated to pay the owner of the IOU amount issued, the ownership of the IOU can be transferred by current owner, and all transactions will be recorded on the ledger.

Before getting started, you should have [Project Dovetail™ Studio](../../getting-started/installation/) and [Project Dovetail™ CLI](../../getting-started/dovetail-cli/) installed, and your [development environment setup](../../getting-started/environment-prerequisites/).

If you have Tibco Cloud subscription or Enterprise Flogo studio, you can also upload Dovetail extentions to start smart contract and dapp development. the extensions are included in the corda.zip (see below) under artifacts/studio folder.

Follow these steps to create the empty structure of the tutorial:

> * Create the tutorial initial structure

 * iou_tutorials
    * artifacts
    * cli
    * network
        * corda

```
mkdir -p iou_tutorial/artifacts
mkdir -p iou_tutorial/cli
mkdir -p iou_tutorial/network/corda
```

## Copy IOU network nodes and scripts
corda.zip has pre-implemented and generated artifacts, nodes and scripts that will help you to get the corda network up and running.

```
curl -OL https://TIBCOSoftware.github.io/dovetail/tutorials/iou/corda.zip && \
unzip corda.zip && \
rm corda.zip
```

* copy extracted network/corda/* to your network/corda folder
* copy extracted cli/* to your cli folder
