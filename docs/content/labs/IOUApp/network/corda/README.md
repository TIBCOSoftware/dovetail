# Test IOU Smart Contract in R3 Cora Network

We will use [Corda Demo Bench](https://docs.corda.net/demobench.html), install it if you don't have it installed already.

## 1. Start up Corda Demo Bench
   * First add all jars from corda directory
   * Start up 3 nodes: Notary, charlie, alice, and bob
   * Make sure they are all running befor contine. Follow the video for step by step instructions.

<p><a target="_blank" rel="noopener noreferrer" href="../../recordings/corddemo.gif"><img src="../../recordings/corddemo.gif" alt="Corda Demo Bench" style="max-width:75%;"></a></p>

## 2. Issue an IOU

At charlie terminal, issue an IOU

> flow start com.example.iou.IOUIssueInitiatorFlow iouValue: $100, owner: "O=alice,L=New York,C=US", externalId: charlie100

Now run following command from charlie, alice and bob's terminals, you should see the IOU is now on both charlie and alice's ledgers, but not on bob's

> run vaultQuery contractStateType: com.example.iou.IOU

## 3. Transfer the IOU

At alice termial, transfer IOU to bob

> flow start com.example.iou.IOUTransferInitiatorFlow iouId: charlie100, newOwner: "C=FR,L=Paris,O=bob"

Now run following command from charlie, alice and bob's terminals, you should see the IOU is now on both charlie and bob's ledgers, but no longer on charlie's

> run vaultQuery contractStateType: com.example.iou.IOU
