## Corda IOU Smart Contract Tutorial

In this tutorial, we will walk you through the steps to model, implement and test smart contracts on R3 Corda.

The example is a simple "I owe you" use case, the issuer of an IOU is obligated to pay the owner of the IOU amount issued, the ownership of the IOU can be transferred by current owner, and all transactions will be recorded on the ledger.

Before getting started, you should have [Project Dovetail™ Studio](ch01-01-installation.md) and [Project Dovetail™ CLI](ch01-02-dovetail-cli.md) installed, and your [development environment setup](ch01-03-environment.md).


Follow these steps to create the empty structure of the tutorial:

> * Create the tutorial initial structure

 * iou_tutorials
    * artifacts
    * iou
    * network
        * corda

```
mkdir -p iou_tutorial/artifacts
mkdir -p iou_tutorial/iou
mkdir -p iou_tutorial/network/corda
```

