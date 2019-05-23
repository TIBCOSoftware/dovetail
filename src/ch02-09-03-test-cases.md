## 6. Test IOU Smart Contract on Corda

### 6.5 Run test cases

#### 6.5.1 Issue an IOU

At charlie's terminal, issue an IOU

```
flow start com.charlie.iou.flows.IssueIOUInitiatorImpl owner: "O=alice,L=New York,C=US", amt: $100, extId: charlie100, regulator: "O=regulator,L=New York,C=US"
```

Now run following command from charlie, alice and bob's terminals, you should see the IOU is now on both charlie and alice's ledgers, but not on bob's

```
run vaultQuery contractStateType: com.example.iou.IOU
```
select and copy the uuid of the IOU from the output

#### 6.5.2 Transfer the IOU

At alice's terminal, transfer IOU to bob. 

Before running the command, replace [YOUR_IOU_UUID] with uuid you just copied
```
flow start com.alice.iou.flows.TransferIOUInitiatorImpl linearId: charlie100_[YOUR_IOU_UUID], newOwner: "O=bob,L=New York,C=US", regulator: "O=regulator,L=New York,C=US"
```

Now run following command from charlie, alice and bob's terminals, you should see the IOU is now on both charlie and bob's ledgers, but no longer on charlie's

```
run vaultQuery contractStateType: com.example.iou.IOU
```

#### 6.5.3 Issue cash

At bank's terminal, issue and transfer cash to charlie. 

```
flow start net.corda.finance.flows.CashIssueFlow amount: $200, issuerBankPartyRef: 00, notary: "C=GB,L=London,O=Notary"

flow start net.corda.finance.flows.CashPaymentFlow amount: $200, recipient: "O=charlie,L=New York,C=US", anonymous: true
```

#### 6.5.4 Settle the IOU

At charlie's terminal, settle the IOU with cash

Before running the command, replace [YOUR_IOU_UUID] with uuid you copied

```
flow start com.charlie.iou.flows.SettleIOUInitiatorImpl linearId: charlie100_[YOUR_IOU_UUID], amt: $50, regulator: "O=regulator,L=New York,C=US"
```

#### 6.5.5 Check IOU and Cash status

At charlie's terminal, run following commands, you should see now IOU's paid amount is $50, and there is $150 left on charlie's account.

```
run vaultQuery contractStateType: net.corda.finance.contracts.asset.Cash$State
run vaultQuery contractStateType: com.example.iou.IOU
```

