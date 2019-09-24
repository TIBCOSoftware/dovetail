## 6. Test IOU Smart Contract on Corda

We will need to develop distributed applications(CorDapp) for each interested party to invoke the IOU smart contract. Let's assume following organizations are interested parties:

> * Notary: to validate and notarize transactions
>      * Corda platform implementation
> * Bank : cash issuer
>      * We will use Corda finance cash flows to issue and transfer cash
> * Regulator : an observer that receives all IOU transactions
>      * We will use com.tibco.dovetail.container.cordapp.flows.ObserverFlowReceiver to receive and record transactions
> * charlie: issuer of an IOU
>      * We will implement IssueIOU initiator flow to issue an IOU to alice
>      * We will implement TransferIOU receiver flow to sign and record TransferIOU transaction
>      * We will implment SettleIOU initiator flow to pay IOU with Cash to bob
>      * We will use com.tibco.dovetail.container.cordapp.flows.ObserverFlowInitiator to send IssueIOU and SettleIOU transactions to Regulator
> * alice: original owner of charlie's IOU
>      * We will implement IssueIOU receiver flow to sign and record IssueIOU transaction
>      * We will implement TransferIOU initiator flow to transfer the IOU bob
>      * We will use com.tibco.dovetail.container.cordapp.flows.ObserverFlowInitiator to send TransferIOU transaction to Regulator
> * bob : new owner of charlie's IOU
>      * We will implement TransferIOU receiver flow to sign and record TransferIOU transaction
>      * We will implement SettleIOU receiver flow to sign and record SettleIOU transaction

We will also use confidential identity in all transactions to protect IOU smart contract participants (charlie, alice and bob).

OK, let's get started.