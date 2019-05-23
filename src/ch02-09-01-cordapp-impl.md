## 6. Test IOU Smart Contract on Corda

## 6.1. CorDapp for charlie
> * Create an App "charlie"

### 6.1.1. IssueIOU initiator flow
> * Create a flow 
>   * flow name: IssueIOUInitiator
>   * select trigger : CorDApp Flow Initiator
>      * select true for "Use confidential identities for this transaction?"
>      * select true for "Send transactions to observers?"
>      * select true for "Send signed transactions to observers in a separate flow?"
>      * click "Next" button
>      * add following flow input parameters
>         * owner :    Type = Party, PartyType = Participant
>         * amt:       Type = Amount<Currency> 
>         * extId:     Type = String
>         * regulator: Type = Party, PartyType = Observer
>      * click "Finish" button
> * Open IssueIOUInitiator flow
>   * select BuildTransactoinProposal activity
>      * select IOU from contract model file dropdown
>      * contract class : com.example.iou.IOUContract 
>      * select IssueIOU from transaction dropdown
>      * map the activity input
>         * use cordapp.createLinearIdFromExternalId($flow.transactionInput.extId) to map iou.linearId
<p><video width="480" height="320" controls="controls">
    <source src="videos/iouissueinitiator.mp4" type="video/mp4">
</video></p>

### 6.1.2. TransferIOU receiver flow

> * Create a flow 
>   * flow name: TransferIOUResponder
>   * select trigger : CorDApp Flow Receiver
>      * select receiver from flow type dropdown
>      * select true for "Use confidential identities for this transaction?"
>      * initiator flow name: com.alice.iou.flows.TransferIOUInitiator
>      * click "Finish" button

<p><video width="480" height="320" controls="controls">
    <source src="videos/ioutransferresponder.mp4" type="video/mp4">
</video></p>

### 6.1.3. IssueIOU initiator flow
> * Create a flow 
>   * flow name: SettleIOUInitiator
>   * select trigger : CorDApp Flow Initiator
>      * select true for "Use confidential identities for this transaction?"
>      * select true for "Send transactions to observers?"
>      * select true for "Send signed transactions to observers in a separate flow?"
>      * click "Next" button
>      * add following flow input parameters
>         * linearId : Type = LinearId
>         * amt:       Type = Amount<Currency> 
>         * regulator: Type = Party, PartyType = Observer
>      * click "Finish" button
> * Open SettleIOUInitiator flow
>   * add SmartContract-Corda/SimpleVaultQuery activity
>      * Configuration screen
>         * select IOU contract
>         * select com.example.iou.IOU asset
>         * leave default for other properties
>      * Map input
>   * if IOU is found
>      * add SmartContract-Corda/CashWallet activity to retrieve cash
>      * add General/Mapper activity to convert wallet output to primitive array of refs
>      * move BuildTransactoinProposal activity from main branch to here
>      * select IOU from contract model file dropdown
>      * contract class : com.example.iou.IOUContract 
>      * select SettleIOU from transaction dropdown
>      * map the activity input
>   * otherwise throw error

<p><video width="480" height="320" controls="controls">
    <source src="videos/iousettleinitiator.mp4" type="video/mp4">
</video></p>

## 6.2. CorDapp for alice
> * Create an App "alice"

### 6.2.1. IssueIOU receiver flow

Create IssueIOUResponder flow, instructions are same as #6.1.2

> * initiator flow name: com.charlie.iou.flows.IssueIOUInitiator

### 6.2.1. TransferIOU initiator flow

<p><video width="480" height="320" controls="controls">
    <source src="videos/ioutransferinitiator.mp4" type="video/mp4">
</video></p>

## 6.3. CorDapp for bob
> * Create an App "bob"

### 6.3.1. TransferIOU receiver flow

Create TransferIOUResponder flow, instructions are same as #6.1.2
> * initiator flow name: com.alice.iou.flows.TransferIOUInitiator

### 6.3.2. SettleIOU receiver flow

Create SettleIOUResponder flow, instructions are same as #6.1.2
> * initiator flow name: com.charlie.iou.flows.SettleIOUInitiator

