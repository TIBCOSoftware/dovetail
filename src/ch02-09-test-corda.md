## 6. Test IOU Smart Contract on Corda

We will use [Corda Demo Bench](https://docs.corda.net/demobench.html), install it if you don't have it installed already.

## 1. Copy Sample CorDapp

Copy [sample R3 CorDapp](https://github.com/TIBCOSoftware/dovetail/blob/master/docs/content/labs/network/corda) to network/corda directory

## 2. Start up Corda Demo Bench
   * Click on "+Add CorDapp" button then browse to corda directory, select a jar to add(you can only add one jar at a time). Repeat, until all jars are added. You can replace kotlin-IOU-1.0.0.jar with your own.
   * Start up 3 nodes: Notary, charlie, alice, and bob
   * Make sure they are all running before continuing. 
   * You can watch the video for step by step instructions.

<p><a target="_blank" rel="noopener noreferrer" href="videos/corddemo.gif"><img src="videos/corddemo.gif" alt="Corda Demo Bench" style="max-width:75%;"></a></p>

## 2. Issue an IOU

At charlie's terminal, issue an IOU

```
flow start com.example.iou.IOUIssueInitiatorFlow iouValue: $100, owner: "O=alice,L=New York,C=US", externalId: charlie100
```

Now run following command from charlie, alice and bob's terminals, you should see the IOU is now on both charlie and alice's ledgers, but not on bob's

```
run vaultQuery contractStateType: com.example.iou.IOU
```

## 3. Transfer the IOU

At alice's terminal, transfer IOU to bob

```
flow start com.example.iou.IOUTransferInitiatorFlow iouId: charlie100, newOwner: "C=FR,L=Paris,O=bob"
```

Now run following command from charlie, alice and bob's terminals, you should see the IOU is now on both charlie and bob's ledgers, but no longer on charlie's

```
run vaultQuery contractStateType: com.example.iou.IOU
```
